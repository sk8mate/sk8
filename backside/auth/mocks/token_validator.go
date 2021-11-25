// Code generated by MockGen. DO NOT EDIT.
// Source: sk8.town/backside/auth (interfaces: TokenValidator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	auth "sk8.town/backside/auth"
	errs "sk8.town/backside/errs"
)

// MockTokenValidator is a mock of TokenValidator interface
type MockTokenValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenValidatorMockRecorder
}

// MockTokenValidatorMockRecorder is the mock recorder for MockTokenValidator
type MockTokenValidatorMockRecorder struct {
	mock *MockTokenValidator
}

// NewMockTokenValidator creates a new mock instance
func NewMockTokenValidator(ctrl *gomock.Controller) *MockTokenValidator {
	mock := &MockTokenValidator{ctrl: ctrl}
	mock.recorder = &MockTokenValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenValidator) EXPECT() *MockTokenValidatorMockRecorder {
	return m.recorder
}

// Verify mocks base method
func (m *MockTokenValidator) Verify(arg0 string) (*auth.UserClaims, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0)
	ret0, _ := ret[0].(*auth.UserClaims)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockTokenValidatorMockRecorder) Verify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockTokenValidator)(nil).Verify), arg0)
}
