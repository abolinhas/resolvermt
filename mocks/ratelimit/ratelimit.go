// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/ratelimit (interfaces: Limiter)

// Package mock_ratelimit is a generated GoMock package.
package mock_ratelimit

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockLimiter is a mock of Limiter interface.
type MockLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockLimiterMockRecorder
}

// MockLimiterMockRecorder is the mock recorder for MockLimiter.
type MockLimiterMockRecorder struct {
	mock *MockLimiter
}

// NewMockLimiter creates a new mock instance.
func NewMockLimiter(ctrl *gomock.Controller) *MockLimiter {
	mock := &MockLimiter{ctrl: ctrl}
	mock.recorder = &MockLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLimiter) EXPECT() *MockLimiterMockRecorder {
	return m.recorder
}

// Take mocks base method.
func (m *MockLimiter) Take() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Take")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Take indicates an expected call of Take.
func (mr *MockLimiterMockRecorder) Take() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*MockLimiter)(nil).Take))
}
