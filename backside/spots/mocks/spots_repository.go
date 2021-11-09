// Code generated by MockGen. DO NOT EDIT.
// Source: sk8.town/backside/spots/domain (interfaces: SpotsRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errs "sk8.town/backside/errs"
	domain "sk8.town/backside/spots/domain"
)

// MockSpotsRepository is a mock of SpotsRepository interface.
type MockSpotsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSpotsRepositoryMockRecorder
}

// MockSpotsRepositoryMockRecorder is the mock recorder for MockSpotsRepository.
type MockSpotsRepositoryMockRecorder struct {
	mock *MockSpotsRepository
}

// NewMockSpotsRepository creates a new mock instance.
func NewMockSpotsRepository(ctrl *gomock.Controller) *MockSpotsRepository {
	mock := &MockSpotsRepository{ctrl: ctrl}
	mock.recorder = &MockSpotsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpotsRepository) EXPECT() *MockSpotsRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockSpotsRepository) Add(arg0 *domain.Spot) (*domain.Spot, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(*domain.Spot)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockSpotsRepositoryMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSpotsRepository)(nil).Add), arg0)
}

// Delete mocks base method.
func (m *MockSpotsRepository) Delete(arg0 int) *errs.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*errs.AppError)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSpotsRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSpotsRepository)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockSpotsRepository) Get(arg0 int) (*domain.Spot, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*domain.Spot)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSpotsRepositoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSpotsRepository)(nil).Get), arg0)
}

// GetAll mocks base method.
func (m *MockSpotsRepository) GetAll() ([]*domain.Spot, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*domain.Spot)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockSpotsRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockSpotsRepository)(nil).GetAll))
}

// Update mocks base method.
func (m *MockSpotsRepository) Update(arg0 int, arg1 *domain.Spot) (*domain.Spot, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*domain.Spot)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSpotsRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSpotsRepository)(nil).Update), arg0, arg1)
}
