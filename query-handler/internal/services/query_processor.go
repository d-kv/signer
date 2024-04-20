package services

import "d-kv/signer/query-handler/internal/usecase"

type QueryProcessor struct {
	TenantRepo      usecase.TenantRepo
	DeviceRepo      usecase.DeviceRepo
	IntegrationRepo usecase.IntegrationRepo
	ProfileRepo     usecase.ProfileRepo
	BundleIdRepo    usecase.BundleIdRepo
	CertificateRepo usecase.CertificateRepo
	UserRepo        usecase.UserRepo
	CapabilityRepo  usecase.CapabilityRepo
}
