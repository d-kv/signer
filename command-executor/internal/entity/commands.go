package entity

type EnableCapabilityType struct {
	ID            uint
	TenantId      uint
	IntegrationId uint

	BundleId       string
	CapabilityType CapabilityType

	Status Status
}

type CreateBundleId struct {
	ID            uint
	TenantId      uint
	IntegrationId uint

	BundleIdentifier string
	BundleName       string
	BundlePlatform   Platform
	SeedId           string

	Status Status
}

type CreateDevice struct {
	ID            uint
	TenantId      uint
	IntegrationId uint

	DeviceName     string
	DevicePlatform Platform
	DeviceUdid     string

	Status Status
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
