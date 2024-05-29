package entity

type URL string

const (
	CapabilitiesURL URL = "https://api.appstoreconnect.apple.com/v1/bundleIdCapabilities"
	DevicesURL      URL = "https://api.appstoreconnect.apple.com/v1/devices"
	BundleIdsURL    URL = "https://api.appstoreconnect.apple.com/v1/bundleIds"
	CertificatesURL URL = "https://api.appstoreconnect.apple.com/v1/certificates"
	ProfilesURL     URL = "https://api.appstoreconnect.apple.com/v1/profiles"
)
