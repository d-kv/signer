package entity

type EnableCapabilityType struct {
	TenantId       int
	IntegrationId  int
	BundleId       string
	CapabilityType CapabilityType
	Status         Status
}

type CreateBundleId struct {
	TenantId         int
	IntegrationId    int
	BundleIdentifier int
	BundleName       string
	BundlePlatform   Platform
	SeedId           string
	Status           Status
}

type CreateDevice struct {
	TenantId       int
	IntegrationId  int
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

type CapabilityType string

const (
	InAppPurchase     CapabilityType = "IN_APP_PURCHASE"
	PushNotifications CapabilityType = "PUSH_NOTIFICATIONS"
	ApplePay          CapabilityType = "APPLE_PAY"
	NfcTagReading     CapabilityType = "NFC_TAG_READING"
	AppleIdAuth       CapabilityType = "APPLE_ID_AUTH"
)

type Status string

const (
	Created    Status = "CREATED"
	Processing Status = "PROCESSING"
	Completed  Status = "COMPLETED"
)
