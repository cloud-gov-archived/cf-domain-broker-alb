package mocks

import (
	models "github.com/cloud-gov/cf-domain-broker-alb/models"
	mock "github.com/stretchr/testify/mock"
)

// RouteManagerIface is an autogenerated mock type for the RouteManagerIface type
type RouteManagerIface struct {
	mock.Mock
}

// Create provides a mock function with given fields: instanceId, domains
func (_m *RouteManagerIface) Create(instanceId string, domains []string) (*models.Route, error) {
	ret := _m.Called(instanceId, domains)

	var r0 *models.Route
	if rf, ok := ret.Get(0).(func(string, []string) *models.Route); ok {
		r0 = rf(instanceId, domains)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(instanceId, domains)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOrphanedCerts provides a mock function with given fields:
func (_m *RouteManagerIface) DeleteOrphanedCerts() {
	_m.Called()
}

// Destroy provides a mock function with given fields: instanceId
func (_m *RouteManagerIface) Destroy(instanceId string) error {
	ret := _m.Called(instanceId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(instanceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: instanceId
func (_m *RouteManagerIface) Get(instanceId string) (*models.Route, error) {
	ret := _m.Called(instanceId)

	var r0 *models.Route
	if rf, ok := ret.Get(0).(func(string) *models.Route); ok {
		r0 = rf(instanceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(instanceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDNSInstructions provides a mock function with given fields: route
func (_m *RouteManagerIface) GetDNSInstructions(route *models.Route) (string, error) {
	ret := _m.Called(route)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.Route) string); ok {
		r0 = rf(route)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Route) error); ok {
		r1 = rf(route)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Poll provides a mock function with given fields: route
func (_m *RouteManagerIface) Poll(route *models.Route) error {
	ret := _m.Called(route)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Route) error); ok {
		r0 = rf(route)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Populate provides a mock function with given fields:
func (_m *RouteManagerIface) Populate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Renew provides a mock function with given fields: route
func (_m *RouteManagerIface) Renew(route *models.Route) error {
	ret := _m.Called(route)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Route) error); ok {
		r0 = rf(route)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenewAll provides a mock function with given fields:
func (_m *RouteManagerIface) RenewAll() {
	_m.Called()
}

// Update provides a mock function with given fields: instanceId, domains
func (_m *RouteManagerIface) Update(instanceId string, domains []string) error {
	ret := _m.Called(instanceId, domains)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(instanceId, domains)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ models.RouteManagerIface = (*RouteManagerIface)(nil)
