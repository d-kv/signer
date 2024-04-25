package entities

type InputEnableCapabilityInput struct {
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
