// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-problem-api/internal/repo (interfaces: RepoRemover)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	utils "github.com/ozoncp/ocp-problem-api/internal/utils"
)

// MockRepoRemover is a mock of RepoRemover interface.
type MockRepoRemover struct {
	ctrl     *gomock.Controller
	recorder *MockRepoRemoverMockRecorder
}

// MockRepoRemoverMockRecorder is the mock recorder for MockRepoRemover.
type MockRepoRemoverMockRecorder struct {
	mock *MockRepoRemover
}

// NewMockRepoRemover creates a new mock instance.
func NewMockRepoRemover(ctrl *gomock.Controller) *MockRepoRemover {
	mock := &MockRepoRemover{ctrl: ctrl}
	mock.recorder = &MockRepoRemoverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoRemover) EXPECT() *MockRepoRemoverMockRecorder {
	return m.recorder
}

// AddEntities mocks base method.
func (m *MockRepoRemover) AddEntities(arg0 context.Context, arg1 []utils.Problem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEntities", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEntities indicates an expected call of AddEntities.
func (mr *MockRepoRemoverMockRecorder) AddEntities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntities", reflect.TypeOf((*MockRepoRemover)(nil).AddEntities), arg0, arg1)
}

// DescribeEntity mocks base method.
func (m *MockRepoRemover) DescribeEntity(arg0 context.Context, arg1 uint64) (*utils.Problem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEntity", arg0, arg1)
	ret0, _ := ret[0].(*utils.Problem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntity indicates an expected call of DescribeEntity.
func (mr *MockRepoRemoverMockRecorder) DescribeEntity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntity", reflect.TypeOf((*MockRepoRemover)(nil).DescribeEntity), arg0, arg1)
}

// ListEntities mocks base method.
func (m *MockRepoRemover) ListEntities(arg0 context.Context, arg1, arg2 uint64) ([]utils.Problem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntities", arg0, arg1, arg2)
	ret0, _ := ret[0].([]utils.Problem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntities indicates an expected call of ListEntities.
func (mr *MockRepoRemoverMockRecorder) ListEntities(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntities", reflect.TypeOf((*MockRepoRemover)(nil).ListEntities), arg0, arg1, arg2)
}

// RemoveEntity mocks base method.
func (m *MockRepoRemover) RemoveEntity(arg0 context.Context, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEntity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEntity indicates an expected call of RemoveEntity.
func (mr *MockRepoRemoverMockRecorder) RemoveEntity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEntity", reflect.TypeOf((*MockRepoRemover)(nil).RemoveEntity), arg0, arg1)
}
