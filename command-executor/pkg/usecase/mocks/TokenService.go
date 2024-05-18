// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	ecdsa "crypto/ecdsa"
	entity "d-kv/signer/db-common/entity"

	mock "github.com/stretchr/testify/mock"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// GenerateECDSAPrivateKey provides a mock function with given fields: base64PrivateKey
func (_m *TokenService) GenerateECDSAPrivateKey(base64PrivateKey string) (*ecdsa.PrivateKey, error) {
	ret := _m.Called(base64PrivateKey)

	if len(ret) == 0 {
		panic("no return value specified for GenerateECDSAPrivateKey")
	}

	var r0 *ecdsa.PrivateKey
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*ecdsa.PrivateKey, error)); ok {
		return rf(base64PrivateKey)
	}
	if rf, ok := ret.Get(0).(func(string) *ecdsa.PrivateKey); ok {
		r0 = rf(base64PrivateKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecdsa.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(base64PrivateKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJwtToken provides a mock function with given fields: tokenInfo
func (_m *TokenService) GetJwtToken(tokenInfo *entity.IntegrationToken) (string, error) {
	ret := _m.Called(tokenInfo)

	if len(ret) == 0 {
		panic("no return value specified for GetJwtToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.IntegrationToken) (string, error)); ok {
		return rf(tokenInfo)
	}
	if rf, ok := ret.Get(0).(func(*entity.IntegrationToken) string); ok {
		r0 = rf(tokenInfo)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*entity.IntegrationToken) error); ok {
		r1 = rf(tokenInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenService creates a new instance of TokenService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenService {
	mock := &TokenService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}