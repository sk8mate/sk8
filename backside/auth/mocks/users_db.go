// Code generated by MockGen. DO NOT EDIT.
// Source: sk8.town/backside/auth/domain (interfaces: UsersRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	domain "sk8.town/backside/auth/domain"
	errs "sk8.town/backside/errs"
)

// MockUsersRepository is a mock of UsersRepository interface
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockUsersRepository) Add(arg0 *domain.User) *errs.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(*errs.AppError)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockUsersRepositoryMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUsersRepository)(nil).Add), arg0)
}

// Get mocks base method
func (m *MockUsersRepository) Get(arg0 string) (*domain.User, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUsersRepositoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsersRepository)(nil).Get), arg0)
}
