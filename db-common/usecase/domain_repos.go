package usecase

type DomainRepos struct {
	TenantRepo      TenantRepo
	DeviceRepo      DeviceRepo
	IntegrationRepo IntegrationRepo
	ProfileRepo     ProfileRepo
	BundleIdRepo    BundleIdRepo
	CertificateRepo CertificateRepo
	UserRepo        UserRepo
	CapabilityRepo  CapabilityRepo
}
