package entities

type InputEnableCapability struct {
	BundleId       string `json:"bundleId"`
	CapabilityType string `json:"capabilityType"`
}

type InputCreateBundleId struct {
	BundleIdentifier string `json:"bundleId"`
	BundleName       string `json:"bundleName"`
	BundlePlatform   string `json:"platform"`
	SeedId           string `json:"seedId"`
}

type InputCreateDevice struct {
	DeviceName     string `json:"deviceName"`
	DevicePlatform string `json:"devicePlatform"`
	DeviceUdid     string `json:"udid"`
}

type InputCreateProfile struct {
	CertificateIds []string `json:"certificateIds"`
	DeviceIds      []string `json:"deviceIds"`
	BundleId       string   `json:"bundleId-id"`
	ProfileType    string   `json:"type"`
	Name           string   `json:"name"`
}

type InputCreateCertificate struct {
	CsrContent      string `json:"csrContent"`
	CertificateType string `json:"certificateType"`
}
