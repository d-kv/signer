package entity

type EnableCapabilityType struct {
	ID             uint
	TenantId       string
	IntegrationId  string
	BundleId       string
	CapabilityType CapabilityType
	Status         Status
}

type CreateBundleId struct {
	ID               uint
	TenantId         string
	IntegrationId    string
	BundleIdentifier string
	BundleName       string
	BundlePlatform   Platform
	SeedId           string
	Status           Status
}

type CreateDevice struct {
	ID             uint
	TenantId       string
	IntegrationId  string
	DeviceName     string
	UserID         string
	UserName       string
	DevicePlatform Platform
	DeviceUdid     string
	Status         Status
}

type CreateProfile struct {
}

type CertificateType string

const (
	IosDevelopment           CertificateType = "IOS_DEVELOPMENT"
	IosDistribution          CertificateType = "IOS_DISTRIBUTION"
	MacAppDistribution       CertificateType = "MAC_APP_DISTRIBUTION"
	MacInstallerDistribution CertificateType = "MAC_INSTALLER_DISTRIBUTION"
	MacAppDevelopment        CertificateType = "MAC_APP_DEVELOPMENT"
	DeveloperIdKext          CertificateType = "DEVELOPER_ID_KEXT"
	DeveloperIdApplication   CertificateType = "DEVELOPER_ID_APPLICATION"
	Development              CertificateType = "DEVELOPMENT"
	Distribution             CertificateType = "DISTRIBUTION"
	PassTypeId               CertificateType = "PASS_TYPE_ID"
	PassTypeIdWithNfc        CertificateType = "PASS_TYPE_ID_WITH_NFC"
)

var CertificateTypeValues = []string{
	"IOS_DEVELOPMENT",
	"IOS_DISTRIBUTION",
	"MAC_APP_DISTRIBUTION",
	"MAC_INSTALLER_DISTRIBUTION",
	"MAC_APP_DEVELOPMENT",
	"DEVELOPER_ID_KEXT",
	"DEVELOPER_ID_APPLICATION",
	"DEVELOPMENT",
	"DISTRIBUTION",
	"PASS_TYPE_ID",
	"PASS_TYPE_ID_WITH_NFC",
}

type CreateCertificate struct {
	CsrContent string
	Type       CertificateType
}

type Platform string

const (
	Ios   Platform = "IOS"
	MacOs Platform = "MAC_OS"
)

var PlatformValues = []string{"IOS", "MAC_OS"}

type CapabilityType string

const (
	InAppPurchase     CapabilityType = "IN_APP_PURCHASE"
	PushNotifications CapabilityType = "PUSH_NOTIFICATIONS"
	ApplePay          CapabilityType = "APPLE_PAY"
	NfcTagReading     CapabilityType = "NFC_TAG_READING"
	AppleIdAuth       CapabilityType = "APPLE_ID_AUTH"
)

var CapabilityTypeValues = []string{"IN_APP_PURCHASE", "PUSH_NOTIFICATIONS", "APPLE_PAY", "NFC_TAG_READING", "APPLE_ID_AUTH"}

type Status string

const (
	Error      Status = "ERROR"
	Created    Status = "CREATED"
	Processing Status = "PROCESSING"
	Completed  Status = "COMPLETED"
)

var StatusValues = []string{"CREATED", "PROCESSING", "COMPLETED", "ERROR"}
