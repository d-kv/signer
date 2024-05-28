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

type ProfileType string

const (
	IosAppDevelopment         ProfileType = "IOS_APP_DEVELOPMENT"
	IosAppStore               ProfileType = "IOS_APP_STORE"
	IosAppAdHoc               ProfileType = "IOS_APP_ADHOC"
	IosAppInHouse             ProfileType = "IOS_APP_INHOUSE"
	MacAppDevelopment         ProfileType = "MAC_APP_DEVELOPMENT"
	MacAppStore               ProfileType = "MAC_APP_STORE"
	MacAppDirect              ProfileType = "MAC_APP_DIRECT"
	TvosAppDevelopment        ProfileType = "TVOS_APP_DEVELOPMENT"
	TvosAppStore              ProfileType = "TVOS_APP_STORE"
	TvosAppAdHoc              ProfileType = "TVOS_APP_ADHOC"
	TvosAppInHouse            ProfileType = "TVOS_APP_INHOUSE"
	MacCatalystAppDevelopment ProfileType = "MAC_CATALYST_APP_DEVELOPMENT"
	MacCatalystAppStore       ProfileType = "MAC_CATALYST_APP_STORE"
	MacCatalystAppDirect      ProfileType = "MAC_CATALYST_APP_DIRECT"
)

var ProfileTypeValues = []string{
	"IOS_APP_DEVELOPMENT",
	"IOS_APP_STORE",
	"IOS_APP_ADHOC",
	"IOS_APP_INHOUSE",
	"MAC_APP_DEVELOPMENT",
	"MAC_APP_STORE",
	"MAC_APP_DIRECT",
	"TVOS_APP_DEVELOPMENT",
	"TVOS_APP_STORE",
	"TVOS_APP_ADHOC",
	"TVOS_APP_INHOUSE",
	"MAC_CATALYST_APP_DEVELOPMENT",
	"MAC_CATALYST_APP_STORE",
	"MAC_CATALYST_APP_DIRECT",
}

type CreateProfile struct {
	ID             uint
	TenantId       string
	IntegrationId  string
	Name           string
	ProfileType    ProfileType
	BundleId       string
	CertificateIds []string
	DeviceIds      []string
	Status         Status
}

type CertificateType string

const (
	IosDevelopment                   CertificateType = "IOS_DEVELOPMENT"
	IosDistribution                  CertificateType = "IOS_DISTRIBUTION"
	MacAppDistribution               CertificateType = "MAC_APP_DISTRIBUTION"
	MacInstallerDistribution         CertificateType = "MAC_INSTALLER_DISTRIBUTION"
	MacAppDevelopmentCertificateType CertificateType = "MAC_APP_DEVELOPMENT"
	DeveloperIdKext                  CertificateType = "DEVELOPER_ID_KEXT"
	DeveloperIdApplication           CertificateType = "DEVELOPER_ID_APPLICATION"
	Development                      CertificateType = "DEVELOPMENT"
	Distribution                     CertificateType = "DISTRIBUTION"
	PassTypeId                       CertificateType = "PASS_TYPE_ID"
	PassTypeIdWithNfc                CertificateType = "PASS_TYPE_ID_WITH_NFC"
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
	ID            uint
	TenantId      string
	IntegrationId string
	CsrContent    string
	Type          CertificateType
	Status        Status
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
