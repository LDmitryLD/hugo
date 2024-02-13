// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	service "projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"

	mock "github.com/stretchr/testify/mock"
)

// Georer is an autogenerated mock type for the Georer type
type Georer struct {
	mock.Mock
}

// GeoCode provides a mock function with given fields: in
func (_m *Georer) GeoCode(in service.GeoCodeIn) service.GeoCodeOut {
	ret := _m.Called(in)

	var r0 service.GeoCodeOut
	if rf, ok := ret.Get(0).(func(service.GeoCodeIn) service.GeoCodeOut); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(service.GeoCodeOut)
	}

	return r0
}

// SearchAddresses provides a mock function with given fields: in
func (_m *Georer) SearchAddresses(in service.SearchAddressesIn) service.SearchAddressesOut {
	ret := _m.Called(in)

	var r0 service.SearchAddressesOut
	if rf, ok := ret.Get(0).(func(service.SearchAddressesIn) service.SearchAddressesOut); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(service.SearchAddressesOut)
	}

	return r0
}

// NewGeorer creates a new instance of Georer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGeorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Georer {
	mock := &Georer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
