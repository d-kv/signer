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
	DevicePlatform Platform
	DeviceUdid     string
	Status         Status
}

type Platform string

const (
	Ios   Platform = "IOS"
	MacOs Platform = "MAC_OS"
)

var ValidPlatforms = []string{"IOS", "MAC_OS"}

type CapabilityType string

const (
	InAppPurchase     CapabilityType = "IN_APP_PURCHASE"
	PushNotifications CapabilityType = "PUSH_NOTIFICATIONS"
	ApplePay          CapabilityType = "APPLE_PAY"
	NfcTagReading     CapabilityType = "NFC_TAG_READING"
	AppleIdAuth       CapabilityType = "APPLE_ID_AUTH"
)

var ValidCapabilities = []string{"IN_APP_PURCHASE", "PUSH_NOTIFICATIONS", "APPLE_PAY", "NFC_TAG_READING", "APPLE_ID_AUTH"}

type Status string

const (
	Created    Status = "CREATED"
	Processing Status = "PROCESSING"
	Completed  Status = "COMPLETED"
)

var ValidStatuses = []string{"CREATED", "PROCESSING", "COMPLETED", "ERROR"}
