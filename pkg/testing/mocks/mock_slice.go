// Code generated by MockGen. DO NOT EDIT.
// Source: slice.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSlice is a mock of Slice interface.
type MockSlice struct {
	ctrl     *gomock.Controller
	recorder *MockSliceMockRecorder
}

// MockSliceMockRecorder is the mock recorder for MockSlice.
type MockSliceMockRecorder struct {
	mock *MockSlice
}

// NewMockSlice creates a new mock instance.
func NewMockSlice(ctrl *gomock.Controller) *MockSlice {
	mock := &MockSlice{ctrl: ctrl}
	mock.recorder = &MockSliceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlice) EXPECT() *MockSliceMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockSlice) Do(arg0 []int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Do", arg0)
}

// Do indicates an expected call of Do.
func (mr *MockSliceMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockSlice)(nil).Do), arg0)
}
