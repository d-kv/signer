package entity

type DataBaseCommand interface {
	Convert() ApiEntity
	GetId() uint
}

type EnableCapabilityType struct {
	ID            uint
	TenantId      uint
	IntegrationId uint

	BundleId       string
	CapabilityType CapabilityType

	Status Status
}

func (input *EnableCapabilityType) GetId() uint {
	return input.ID
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

func (input *CreateBundleId) GetId() uint {
	return input.ID
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

func (input *CreateDevice) GetId() uint {
	return input.ID
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
	Error      Status = "ERROR"
)

type URL string

const (
	CapabilitiesURL URL = "https://api.appstoreconnect.apple.com/v1/bundleIdCapabilities"
	DevicesURL      URL = "https://api.appstoreconnect.apple.com/v1/devices"
	BundleIdURL     URL = "https://api.appstoreconnect.apple.com/v1/bundleIds"
)
