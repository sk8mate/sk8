// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sk8mate/sk8/backside/places-microservice/service (interfaces: PlacesService)

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/sk8mate/sk8/backside/places-microservice/dto"
	errs "github.com/sk8mate/sk8/backside/places-microservice/errs"
)

// MockPlacesService is a mock of PlacesService interface.
type MockPlacesService struct {
	ctrl     *gomock.Controller
	recorder *MockPlacesServiceMockRecorder
}

// MockPlacesServiceMockRecorder is the mock recorder for MockPlacesService.
type MockPlacesServiceMockRecorder struct {
	mock *MockPlacesService
}

// NewMockPlacesService creates a new mock instance.
func NewMockPlacesService(ctrl *gomock.Controller) *MockPlacesService {
	mock := &MockPlacesService{ctrl: ctrl}
	mock.recorder = &MockPlacesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlacesService) EXPECT() *MockPlacesServiceMockRecorder {
	return m.recorder
}

// GetPlaces mocks base method.
func (m *MockPlacesService) GetPlaces(arg0 dto.PlacesAutocompleteRequest) ([]dto.PlacesAutocompleteResponseEntry, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaces", arg0)
	ret0, _ := ret[0].([]dto.PlacesAutocompleteResponseEntry)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetPlaces indicates an expected call of GetPlaces.
func (mr *MockPlacesServiceMockRecorder) GetPlaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaces", reflect.TypeOf((*MockPlacesService)(nil).GetPlaces), arg0)
}