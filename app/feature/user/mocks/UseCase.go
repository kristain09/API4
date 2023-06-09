// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	user "github.com/kristain09/API4/app/feature/user"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: userID
func (_m *UseCase) GetUserByID(userID string) (user.Core, error) {
	ret := _m.Called(userID)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.Core, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) user.Core); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: hp, password
func (_m *UseCase) Login(hp string, password string) (user.Core, error) {
	ret := _m.Called(hp, password)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (user.Core, error)); ok {
		return rf(hp, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) user.Core); ok {
		r0 = rf(hp, password)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(hp, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newUser
func (_m *UseCase) Register(newUser user.Core) error {
	ret := _m.Called(newUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core) error); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfile provides a mock function with given fields: newData, userID
func (_m *UseCase) UpdateProfile(newData user.Core, userID string) error {
	ret := _m.Called(newData, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core, string) error); ok {
		r0 = rf(newData, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
