// Code generated by MockGen. DO NOT EDIT.
// Source: sk8.town/backside/places (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errors "sk8.town/backside/errs"
	dto "sk8.town/backside/places/dto"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetPlaces mocks base method.
func (m *MockService) GetPlaces(arg0 dto.AutocompleteRequest) ([]dto.AutocompleteEntryResponse, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaces", arg0)
	ret0, _ := ret[0].([]dto.AutocompleteEntryResponse)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetPlaces indicates an expected call of GetPlaces.
func (mr *MockServiceMockRecorder) GetPlaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaces", reflect.TypeOf((*MockService)(nil).GetPlaces), arg0)
}