// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// Add provides a mock function with given fields: name, configuration
func (_m *Config) Add(name string, configuration map[string]interface{}) {
	_m.Called(name, configuration)
}

// Env provides a mock function with given fields: envName, defaultValue
func (_m *Config) Env(envName string, defaultValue ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, envName)
	_ca = append(_ca, defaultValue...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(envName, defaultValue...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Get provides a mock function with given fields: path, defaultValue
func (_m *Config) Get(path string, defaultValue ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, path)
	_ca = append(_ca, defaultValue...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(path, defaultValue...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetBool provides a mock function with given fields: path, defaultValue
func (_m *Config) GetBool(path string, defaultValue ...interface{}) bool {
	var _ca []interface{}
	_ca = append(_ca, path)
	_ca = append(_ca, defaultValue...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...interface{}) bool); ok {
		r0 = rf(path, defaultValue...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetInt provides a mock function with given fields: path, defaultValue
func (_m *Config) GetInt(path string, defaultValue ...interface{}) int {
	var _ca []interface{}
	_ca = append(_ca, path)
	_ca = append(_ca, defaultValue...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...interface{}) int); ok {
		r0 = rf(path, defaultValue...)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetString provides a mock function with given fields: path, defaultValue
func (_m *Config) GetString(path string, defaultValue ...interface{}) string {
	var _ca []interface{}
	_ca = append(_ca, path)
	_ca = append(_ca, defaultValue...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(path, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfig(t mockConstructorTestingTNewConfig) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}