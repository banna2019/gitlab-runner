// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package ca_chain

import (
	x509 "crypto/x509"

	mock "github.com/stretchr/testify/mock"
)

// mockResolver is an autogenerated mock type for the resolver type
type mockResolver struct {
	mock.Mock
}

// Resolve provides a mock function with given fields: certs
func (_m *mockResolver) Resolve(certs []*x509.Certificate) ([]*x509.Certificate, error) {
	ret := _m.Called(certs)

	var r0 []*x509.Certificate
	if rf, ok := ret.Get(0).(func([]*x509.Certificate) []*x509.Certificate); ok {
		r0 = rf(certs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*x509.Certificate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*x509.Certificate) error); ok {
		r1 = rf(certs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}