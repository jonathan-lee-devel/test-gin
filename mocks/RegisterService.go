// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import enums "test-gin/users/enums"
import mock "github.com/stretchr/testify/mock"

import userDtos "test-gin/users/dtos"

// RegisterService is an autogenerated mock type for the RegisterService type
type RegisterService struct {
	mock.Mock
}

// RegisterUser provides a mock function with given fields: input
func (_m *RegisterService) RegisterUser(input *userDtos.RegisterInput) (enums.RegistrationStatus, error) {
	ret := _m.Called(input)

	var r0 enums.RegistrationStatus
	if rf, ok := ret.Get(0).(func(*userDtos.RegisterInput) enums.RegistrationStatus); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(enums.RegistrationStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*userDtos.RegisterInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
