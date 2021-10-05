// Code generated by MockGen. DO NOT EDIT.
// Source: sk8.town/backside/spots (interfaces: SpotService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errs "sk8.town/backside/errs"
	dto "sk8.town/backside/spots/dto"
)

// MockSpotService is a mock of SpotService interface.
type MockSpotService struct {
	ctrl     *gomock.Controller
	recorder *MockSpotServiceMockRecorder
}

// MockSpotServiceMockRecorder is the mock recorder for MockSpotService.
type MockSpotServiceMockRecorder struct {
	mock *MockSpotService
}

// NewMockSpotService creates a new mock instance.
func NewMockSpotService(ctrl *gomock.Controller) *MockSpotService {
	mock := &MockSpotService{ctrl: ctrl}
	mock.recorder = &MockSpotServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpotService) EXPECT() *MockSpotServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockSpotService) Add(arg0 *dto.SpotsAddRequest) (*dto.SpotsAddData, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(*dto.SpotsAddData)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockSpotServiceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSpotService)(nil).Add), arg0)
}
