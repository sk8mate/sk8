// Code generated by MockGen. DO NOT EDIT.
// Source: sk8.town/backside/auth (interfaces: TokenService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	auth "sk8.town/backside/auth"
	errs "sk8.town/backside/errs"
)

// MockTokenService is a mock of TokenService interface
type MockTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenServiceMockRecorder
}

// MockTokenServiceMockRecorder is the mock recorder for MockTokenService
type MockTokenServiceMockRecorder struct {
	mock *MockTokenService
}

// NewMockTokenService creates a new mock instance
func NewMockTokenService(ctrl *gomock.Controller) *MockTokenService {
	mock := &MockTokenService{ctrl: ctrl}
	mock.recorder = &MockTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenService) EXPECT() *MockTokenServiceMockRecorder {
	return m.recorder
}

// CreateToken mocks base method
func (m *MockTokenService) CreateToken(arg0 string) (string, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken
func (mr *MockTokenServiceMockRecorder) CreateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockTokenService)(nil).CreateToken), arg0)
}

// ParseToken mocks base method
func (m *MockTokenService) ParseToken(arg0 string) (*auth.UserClaims, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", arg0)
	ret0, _ := ret[0].(*auth.UserClaims)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken
func (mr *MockTokenServiceMockRecorder) ParseToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockTokenService)(nil).ParseToken), arg0)
}
