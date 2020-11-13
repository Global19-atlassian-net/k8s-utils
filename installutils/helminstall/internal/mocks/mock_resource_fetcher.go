// Code generated by MockGen. DO NOT EDIT.
// Source: ./get_resource.go

// Package mock_internal is a generated GoMock package.
package mock_internal

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResourceFetcher is a mock of resourceFetcher interface
type MockResourceFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockResourceFetcherMockRecorder
}

// MockResourceFetcherMockRecorder is the mock recorder for MockResourceFetcher
type MockResourceFetcherMockRecorder struct {
	mock *MockResourceFetcher
}

// NewMockResourceFetcher creates a new mock instance
func NewMockResourceFetcher(ctrl *gomock.Controller) *MockResourceFetcher {
	mock := &MockResourceFetcher{ctrl: ctrl}
	mock.recorder = &MockResourceFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResourceFetcher) EXPECT() *MockResourceFetcherMockRecorder {
	return m.recorder
}

// GetResource mocks base method
func (m *MockResourceFetcher) GetResource(uri string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", uri)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource
func (mr *MockResourceFetcherMockRecorder) GetResource(uri interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockResourceFetcher)(nil).GetResource), uri)
}
