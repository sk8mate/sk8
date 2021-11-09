// Code generated by MockGen. DO NOT EDIT.
// Source: sk8.town/backside/spots (interfaces: FiltersService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errs "sk8.town/backside/errs"
	dto "sk8.town/backside/spots/dto"
)

// MockFiltersService is a mock of FiltersService interface.
type MockFiltersService struct {
	ctrl     *gomock.Controller
	recorder *MockFiltersServiceMockRecorder
}

// MockFiltersServiceMockRecorder is the mock recorder for MockFiltersService.
type MockFiltersServiceMockRecorder struct {
	mock *MockFiltersService
}

// NewMockFiltersService creates a new mock instance.
func NewMockFiltersService(ctrl *gomock.Controller) *MockFiltersService {
	mock := &MockFiltersService{ctrl: ctrl}
	mock.recorder = &MockFiltersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFiltersService) EXPECT() *MockFiltersServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockFiltersService) GetAll() ([]*dto.FilterData, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*dto.FilterData)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockFiltersServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockFiltersService)(nil).GetAll))
}
