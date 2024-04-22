package entity

type DataBaseCommand interface {
	Convert() ApiEntity
	GetId() uint
}

type URL string

const (
	CapabilitiesURL URL = "https://api.appstoreconnect.apple.com/v1/bundleIdCapabilities"
	DevicesURL      URL = "https://api.appstoreconnect.apple.com/v1/devices"
	BundleIdURL     URL = "https://api.appstoreconnect.apple.com/v1/bundleIds"
)
