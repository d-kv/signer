// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "d-kv/signer/db-common/entity"

	mock "github.com/stretchr/testify/mock"

	pkgentity "d-kv/signer/command-executor/pkg/entity"
)

// DataBaseService is an autogenerated mock type for the DataBaseService type
type DataBaseService struct {
	mock.Mock
}

// UpdateArrayCertificates provides a mock function with given fields: ctx, certificates, converted, err
func (_m *DataBaseService) UpdateArrayCertificates(ctx context.Context, certificates []entity.Certificate, converted *entity.Profile, err error) error {
	ret := _m.Called(ctx, certificates, converted, err)

	if len(ret) == 0 {
		panic("no return value specified for UpdateArrayCertificates")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []entity.Certificate, *entity.Profile, error) error); ok {
		r0 = rf(ctx, certificates, converted, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateArrayDevices provides a mock function with given fields: ctx, devices, converted, err
func (_m *DataBaseService) UpdateArrayDevices(ctx context.Context, devices []entity.Device, converted *entity.Profile, err error) error {
	ret := _m.Called(ctx, devices, converted, err)

	if len(ret) == 0 {
		panic("no return value specified for UpdateArrayDevices")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []entity.Device, *entity.Profile, error) error); ok {
		r0 = rf(ctx, devices, converted, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteBundleId provides a mock function with given fields: ctx, operation, response
func (_m *DataBaseService) WriteBundleId(ctx context.Context, operation entity.CreateBundleId, response *pkgentity.BundleIdResponse) error {
	ret := _m.Called(ctx, operation, response)

	if len(ret) == 0 {
		panic("no return value specified for WriteBundleId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateBundleId, *pkgentity.BundleIdResponse) error); ok {
		r0 = rf(ctx, operation, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteCapability provides a mock function with given fields: ctx, err, operation
func (_m *DataBaseService) WriteCapability(ctx context.Context, err error, operation entity.EnableCapabilityType) error {
	ret := _m.Called(ctx, err, operation)

	if len(ret) == 0 {
		panic("no return value specified for WriteCapability")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, error, entity.EnableCapabilityType) error); ok {
		r0 = rf(ctx, err, operation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteCertificate provides a mock function with given fields: ctx, operation, resp
func (_m *DataBaseService) WriteCertificate(ctx context.Context, operation entity.CreateCertificate, resp *pkgentity.CertificateResponse) error {
	ret := _m.Called(ctx, operation, resp)

	if len(ret) == 0 {
		panic("no return value specified for WriteCertificate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateCertificate, *pkgentity.CertificateResponse) error); ok {
		r0 = rf(ctx, operation, resp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteDevice provides a mock function with given fields: ctx, operation
func (_m *DataBaseService) WriteDevice(ctx context.Context, operation entity.CreateDevice) error {
	ret := _m.Called(ctx, operation)

	if len(ret) == 0 {
		panic("no return value specified for WriteDevice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateDevice) error); ok {
		r0 = rf(ctx, operation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteProfile provides a mock function with given fields: ctx, operation, resp
func (_m *DataBaseService) WriteProfile(ctx context.Context, operation entity.CreateProfile, resp *pkgentity.ProfileResponse) error {
	ret := _m.Called(ctx, operation, resp)

	if len(ret) == 0 {
		panic("no return value specified for WriteProfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateProfile, *pkgentity.ProfileResponse) error); ok {
		r0 = rf(ctx, operation, resp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDataBaseService creates a new instance of DataBaseService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataBaseService(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataBaseService {
	mock := &DataBaseService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
