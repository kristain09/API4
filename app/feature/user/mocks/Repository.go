// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	user "github.com/kristain09/API4/app/feature/user"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: userID
func (_m *Repository) GetUserByID(userID string) (user.Core, error) {
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

// Insert provides a mock function with given fields: newUser
func (_m *Repository) Insert(newUser user.Core) (user.Core, error) {
	ret := _m.Called(newUser)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(user.Core) (user.Core, error)); ok {
		return rf(newUser)
	}
	if rf, ok := ret.Get(0).(func(user.Core) user.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(user.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: hp, password
func (_m *Repository) Login(hp string, password string) (user.Core, error) {
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

// UpdateProfile provides a mock function with given fields: newData, userID
func (_m *Repository) UpdateProfile(newData user.Core, userID string) error {
	ret := _m.Called(newData, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core, string) error); ok {
		r0 = rf(newData, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
