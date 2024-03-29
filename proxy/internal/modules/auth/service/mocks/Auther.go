// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"
	service "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/service"

	mock "github.com/stretchr/testify/mock"
)

// Auther is an autogenerated mock type for the Auther type
type Auther struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, in
func (_m *Auther) Login(ctx context.Context, in service.AuthorizeIn) service.AuthorizeOut {
	ret := _m.Called(ctx, in)

	var r0 service.AuthorizeOut
	if rf, ok := ret.Get(0).(func(context.Context, service.AuthorizeIn) service.AuthorizeOut); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(service.AuthorizeOut)
	}

	return r0
}

// Register provides a mock function with given fields: in
func (_m *Auther) Register(in service.RegisterIn) service.RegisterOut {
	ret := _m.Called(in)

	var r0 service.RegisterOut
	if rf, ok := ret.Get(0).(func(service.RegisterIn) service.RegisterOut); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(service.RegisterOut)
	}

	return r0
}

// NewAuther creates a new instance of Auther. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuther(t interface {
	mock.TestingT
	Cleanup(func())
}) *Auther {
	mock := &Auther{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
