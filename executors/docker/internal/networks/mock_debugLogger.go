// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package networks

import mock "github.com/stretchr/testify/mock"

// mockDebugLogger is an autogenerated mock type for the debugLogger type
type mockDebugLogger struct {
	mock.Mock
}

// Debugln provides a mock function with given fields: args
func (_m *mockDebugLogger) Debugln(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}
