// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package kubernetes

import mock "github.com/stretchr/testify/mock"

// mockFeatureChecker is an autogenerated mock type for the featureChecker type
type mockFeatureChecker struct {
	mock.Mock
}

// IsHostAliasSupported provides a mock function with given fields:
func (_m *mockFeatureChecker) IsHostAliasSupported() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
