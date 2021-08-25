// Code generated by MockGen. DO NOT EDIT.
// Source: sk8.town/backside/places (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errors "sk8.town/backside/errs"
	domain "sk8.town/backside/places/domain"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetPlaces mocks base method.
func (m *MockRepository) GetPlaces(arg0, arg1 string) (*domain.GetPlacesResponse, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaces", arg0, arg1)
	ret0, _ := ret[0].(*domain.GetPlacesResponse)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetPlaces indicates an expected call of GetPlaces.
func (mr *MockRepositoryMockRecorder) GetPlaces(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaces", reflect.TypeOf((*MockRepository)(nil).GetPlaces), arg0, arg1)
}
