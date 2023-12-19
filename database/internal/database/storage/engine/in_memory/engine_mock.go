// Code generated by MockGen. DO NOT EDIT.
// Source: engine.go

// Package in_memory is a generated GoMock package.
package in_memory

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockhashTable is a mock of hashTable interface.
type MockhashTable struct {
	ctrl     *gomock.Controller
	recorder *MockhashTableMockRecorder
}

// MockhashTableMockRecorder is the mock recorder for MockhashTable.
type MockhashTableMockRecorder struct {
	mock *MockhashTable
}

// NewMockhashTable creates a new mock instance.
func NewMockhashTable(ctrl *gomock.Controller) *MockhashTable {
	mock := &MockhashTable{ctrl: ctrl}
	mock.recorder = &MockhashTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockhashTable) EXPECT() *MockhashTableMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockhashTable) Del(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Del", arg0)
}

// Del indicates an expected call of Del.
func (mr *MockhashTableMockRecorder) Del(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockhashTable)(nil).Del), arg0)
}

// Get mocks base method.
func (m *MockhashTable) Get(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockhashTableMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockhashTable)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockhashTable) Set(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockhashTableMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockhashTable)(nil).Set), arg0, arg1)
}
