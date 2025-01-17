// Code generated by MockGen. DO NOT EDIT.
// Source: ./state/factory/factory.go

// Package mock_factory is a generated GoMock package.
package mock_factory

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	hash "github.com/iotexproject/go-pkgs/hash"
	address "github.com/iotexproject/iotex-address/address"
	action "github.com/iotexproject/iotex-core/action"
	evm "github.com/iotexproject/iotex-core/action/protocol/execution/evm"
	state "github.com/iotexproject/iotex-core/state"
	factory "github.com/iotexproject/iotex-core/state/factory"
	reflect "reflect"
)

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockFactory) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockFactoryMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockFactory)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockFactory) Stop(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockFactoryMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockFactory)(nil).Stop), arg0)
}

// AccountState mocks base method
func (m *MockFactory) AccountState(arg0 string) (*state.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountState", arg0)
	ret0, _ := ret[0].(*state.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountState indicates an expected call of AccountState
func (mr *MockFactoryMockRecorder) AccountState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountState", reflect.TypeOf((*MockFactory)(nil).AccountState), arg0)
}

// RootHash mocks base method
func (m *MockFactory) RootHash() hash.Hash256 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootHash")
	ret0, _ := ret[0].(hash.Hash256)
	return ret0
}

// RootHash indicates an expected call of RootHash
func (mr *MockFactoryMockRecorder) RootHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHash", reflect.TypeOf((*MockFactory)(nil).RootHash))
}

// RootHashByHeight mocks base method
func (m *MockFactory) RootHashByHeight(arg0 uint64) (hash.Hash256, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootHashByHeight", arg0)
	ret0, _ := ret[0].(hash.Hash256)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootHashByHeight indicates an expected call of RootHashByHeight
func (mr *MockFactoryMockRecorder) RootHashByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHashByHeight", reflect.TypeOf((*MockFactory)(nil).RootHashByHeight), arg0)
}

// Height mocks base method
func (m *MockFactory) Height() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Height indicates an expected call of Height
func (mr *MockFactoryMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockFactory)(nil).Height))
}

// NewWorkingSet mocks base method
func (m *MockFactory) NewWorkingSet() (factory.WorkingSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWorkingSet")
	ret0, _ := ret[0].(factory.WorkingSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWorkingSet indicates an expected call of NewWorkingSet
func (mr *MockFactoryMockRecorder) NewWorkingSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWorkingSet", reflect.TypeOf((*MockFactory)(nil).NewWorkingSet))
}

// RunActions mocks base method
func (m *MockFactory) RunActions(arg0 context.Context, arg1 []action.SealedEnvelope) ([]*action.Receipt, factory.WorkingSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunActions", arg0, arg1)
	ret0, _ := ret[0].([]*action.Receipt)
	ret1, _ := ret[1].(factory.WorkingSet)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RunActions indicates an expected call of RunActions
func (mr *MockFactoryMockRecorder) RunActions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunActions", reflect.TypeOf((*MockFactory)(nil).RunActions), arg0, arg1)
}

// PickAndRunActions mocks base method
func (m *MockFactory) PickAndRunActions(arg0 context.Context, arg1 map[string][]action.SealedEnvelope, arg2 []action.SealedEnvelope) ([]*action.Receipt, []action.SealedEnvelope, factory.WorkingSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PickAndRunActions", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*action.Receipt)
	ret1, _ := ret[1].([]action.SealedEnvelope)
	ret2, _ := ret[2].(factory.WorkingSet)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// PickAndRunActions indicates an expected call of PickAndRunActions
func (mr *MockFactoryMockRecorder) PickAndRunActions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PickAndRunActions", reflect.TypeOf((*MockFactory)(nil).PickAndRunActions), arg0, arg1, arg2)
}

// SimulateExecution mocks base method
func (m *MockFactory) SimulateExecution(arg0 context.Context, arg1 address.Address, arg2 *action.Execution, arg3 evm.GetBlockHash) ([]byte, *action.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulateExecution", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*action.Receipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SimulateExecution indicates an expected call of SimulateExecution
func (mr *MockFactoryMockRecorder) SimulateExecution(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateExecution", reflect.TypeOf((*MockFactory)(nil).SimulateExecution), arg0, arg1, arg2, arg3)
}

// Commit mocks base method
func (m *MockFactory) Commit(arg0 factory.WorkingSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockFactoryMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockFactory)(nil).Commit), arg0)
}

// CandidatesByHeight mocks base method
func (m *MockFactory) CandidatesByHeight(arg0 uint64) ([]*state.Candidate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CandidatesByHeight", arg0)
	ret0, _ := ret[0].([]*state.Candidate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CandidatesByHeight indicates an expected call of CandidatesByHeight
func (mr *MockFactoryMockRecorder) CandidatesByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CandidatesByHeight", reflect.TypeOf((*MockFactory)(nil).CandidatesByHeight), arg0)
}

// State mocks base method
func (m *MockFactory) State(arg0 hash.Hash160, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// State indicates an expected call of State
func (mr *MockFactoryMockRecorder) State(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockFactory)(nil).State), arg0, arg1)
}
