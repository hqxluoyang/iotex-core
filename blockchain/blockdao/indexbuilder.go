// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package blockdao

import (
	"strconv"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/iotexproject/iotex-core/blockchain/block"
	"github.com/iotexproject/iotex-core/blockindex"
	"github.com/iotexproject/iotex-core/db"
	"github.com/iotexproject/iotex-core/pkg/log"
	"github.com/iotexproject/iotex-core/pkg/prometheustimer"
)

var batchSizeMtc = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "iotex_indexer_batch_size",
		Help: "Indexer batch size",
	},
	[]string{},
)

func init() {
	prometheus.MustRegister(batchSizeMtc)
}

type addrIndex map[hash.Hash160]db.CountingIndex

// IndexBuilder defines the index builder
type IndexBuilder struct {
	pendingBlks  chan *block.Block
	cancelChan   chan interface{}
	timerFactory *prometheustimer.TimerFactory
	dao          BlockDAO
	indexer      blockindex.Indexer
}

// NewIndexBuilder instantiates an index builder
func NewIndexBuilder(chainID uint32, dao BlockDAO, indexer blockindex.Indexer) (*IndexBuilder, error) {
	timerFactory, err := prometheustimer.New(
		"iotex_indexer_batch_time",
		"Indexer batch time",
		[]string{"topic", "chainID"},
		[]string{"default", strconv.FormatUint(uint64(chainID), 10)},
	)
	if err != nil {
		return nil, err
	}
	return &IndexBuilder{
		pendingBlks:  make(chan *block.Block, 8),
		cancelChan:   make(chan interface{}),
		timerFactory: timerFactory,
		dao:          dao,
		indexer:      indexer,
	}, nil
}

// Start starts the index builder
func (ib *IndexBuilder) Start(_ context.Context) error {
	if err := ib.init(); err != nil {
		return err
	}
	// start handler to index incoming new block
	go ib.handler()
	return nil
}

// Stop stops the index builder
func (ib *IndexBuilder) Stop(_ context.Context) error {
	close(ib.cancelChan)
	return nil
}

// HandleBlock handles the block and create the indices for the actions and receipts in it
func (ib *IndexBuilder) HandleBlock(blk *block.Block) error {
	ib.pendingBlks <- blk
	return nil
}

func (ib *IndexBuilder) handler() {
	for {
		select {
		case <-ib.cancelChan:
			return
		case blk := <-ib.pendingBlks:
			timer := ib.timerFactory.NewTimer("indexBlock")
			if err := ib.indexer.IndexBlock(blk, false); err != nil {
				log.L().Error(
					"Error when indexing the block",
					zap.Uint64("height", blk.Height()),
					zap.Error(err),
				)
			}
			if err := ib.indexer.IndexAction(blk); err != nil {
				log.L().Error(
					"Error when indexing the block",
					zap.Uint64("height", blk.Height()),
					zap.Error(err),
				)
			}
			if err := ib.indexer.Commit(); err != nil {
				log.L().Error(
					"Error when committing the block index",
					zap.Uint64("height", blk.Height()),
					zap.Error(err),
				)
			}
			timer.End()
		}
	}
}

func (ib *IndexBuilder) init() error {
	startHeight, err := ib.indexer.GetBlockchainHeight()
	if err != nil {
		return err
	}
	tipHeight, err := ib.dao.GetTipHeight()
	if err != nil {
		return err
	}
	if startHeight == tipHeight {
		// indexer height consistent with dao height
		zap.L().Info("Consistent DB", zap.Uint64("height", startHeight))
		return nil
	}
	if startHeight > tipHeight {
		// indexer height > dao height
		// this shouldn't happen unless blocks are deliberately removed from dao w/o removing index
		// in this case we revert the extra block index, but nothing we can do to revert action index
		zap.L().Error("Inconsistent DB: indexer height > blockDAO height",
			zap.Uint64("indexer", startHeight), zap.Uint64("blockDAO", tipHeight))
		return ib.indexer.RevertBlocks(startHeight - tipHeight)
	}
	// update index to latest block
	for startHeight++; startHeight <= tipHeight; startHeight++ {
		blk, err := ib.dao.GetBlockByHeight(startHeight)
		if err != nil {
			return err
		}
		if err := ib.indexer.IndexBlock(blk, true); err != nil {
			return err
		}
		if err := ib.indexer.IndexAction(blk); err != nil {
			return err
		}
		// commit once every 10000 blocks
		if startHeight%10000 == 0 || startHeight == tipHeight {
			if err := ib.indexer.Commit(); err != nil {
				return err
			}
			zap.L().Info("Finished indexing blocks up to", zap.Uint64("height", startHeight))
		}
	}
	if startHeight == tipHeight {
		// successfully migrated to latest block
		zap.L().Info("Finished migrating DB", zap.Uint64("height", startHeight))
		return ib.purgeObsoleteIndex()
	}
	return nil
}

func (ib *IndexBuilder) purgeObsoleteIndex() error {
	store := ib.dao.KVStore()
	if err := store.Delete(blockAddressActionMappingNS, nil); err != nil {
		return err
	}
	if err := store.Delete(blockAddressActionCountMappingNS, nil); err != nil {
		return err
	}
	if err := store.Delete(blockActionBlockMappingNS, nil); err != nil {
		return err
	}
	if err := store.Delete(blockActionReceiptMappingNS, nil); err != nil {
		return err
	}
	if err := store.Delete(numActionsNS, nil); err != nil {
		return err
	}
	if err := store.Delete(transferAmountNS, nil); err != nil {
		return err
	}
	return nil
}
