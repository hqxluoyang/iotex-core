// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package account

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-core/action/protocol"
	"github.com/iotexproject/iotex-core/config"
	"github.com/iotexproject/iotex-core/db"
	"github.com/iotexproject/iotex-core/pkg/util/byteutil"
	"github.com/iotexproject/iotex-core/state"
	"github.com/iotexproject/iotex-core/state/factory"
	"github.com/iotexproject/iotex-core/test/testaddress"
	"github.com/iotexproject/iotex-core/testutil"
)

func TestLoadOrCreateAccountState(t *testing.T) {
	require := require.New(t)

	cfg := config.Default
	sf, err := factory.NewFactory(cfg, factory.PrecreatedTrieDBOption(db.NewMemKVStore()))
	require.NoError(err)
	require.NoError(sf.Start(context.Background()))
	ws, err := sf.NewWorkingSet()
	require.NoError(err)
	addrv1 := testaddress.Addrinfo["producer"]
	s, err := LoadAccount(ws, byteutil.BytesTo20B(addrv1.Bytes()))
	require.NoError(err)
	require.Equal(s.Balance, state.EmptyAccount().Balance)
	require.Equal(s.VotingWeight, state.EmptyAccount().VotingWeight)
	s, err = LoadOrCreateAccount(ws, addrv1.String(), big.NewInt(5))
	require.NoError(err)
	s, err = LoadAccount(ws, byteutil.BytesTo20B(addrv1.Bytes()))
	require.NoError(err)
	require.Equal(uint64(0x0), s.Nonce)
	require.Equal("5", s.Balance.String())

	gasLimit := testutil.TestGasLimit
	ctx := protocol.WithRunActionsCtx(context.Background(),
		protocol.RunActionsCtx{
			Producer:        testaddress.Addrinfo["producer"],
			GasLimit:        &gasLimit,
			EnableGasCharge: testutil.EnableGasCharge,
		})
	_, _, err = ws.RunActions(ctx, 0, nil)
	require.NoError(err)
	require.NoError(sf.Commit(ws))
	ss, err := sf.AccountState(addrv1.String())
	require.Nil(err)
	require.Equal(uint64(0x0), ss.Nonce)
	require.Equal("5", ss.Balance.String())
}