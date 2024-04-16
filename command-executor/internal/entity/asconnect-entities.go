package entity

type ApiCreateBundleId struct {
	Data BundleIdData `json:"data"`
}

type BundleIdAttributes struct {
	Name           string `json:"name"`
	Identifier     string `json:"identifier"`
	SeedID         string `json:"seedId"`
	Platform       string `json:"platform"`
	BundleIDSuffix string `json:"bundleIdSuffix,omitempty"`
}

type BundleIdData struct {
	Type       string             `json:"type"`
	Attributes BundleIdAttributes `json:"attributes"`
}

//BundleID

type ApiCreateDevice struct {
	Data DeviceData `json:"data"`
}

type DeviceAttributes struct {
	Name     string `json:"name"`
	UDID     string `json:"udid"`
	Platform string `json:"platform"`
}

type DeviceData struct {
	Attributes DeviceAttributes `json:"attributes"`
	Type       string           `json:"type"`
}

//Device

type ApiEnableCapability struct {
	Data DataForCapability `json:"data"`
}

type OptionData struct {
	Key              string `json:"key,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	EnabledByDefault string `json:"enabledByDefault,omitempty"`
	Enabled          string `json:"enabled,omitempty"`
	SupportsWildcard string `json:"supportsWildcard,omitempty"`
}

type SettingData struct {
	Key              string       `json:"key,omitempty"`
	Name             string       `json:"name,omitempty"`
	Description      string       `json:"description,omitempty"`
	EnabledByDefault string       `json:"enabledByDefault,omitempty"`
	Visible          string       `json:"visible,omitempty"`
	AllowedInstances string       `json:"allowedInstances,omitempty"`
	MinInstances     string       `json:"minInstances,omitempty"`
	Options          []OptionData `json:"options,omitempty"`
}

type AttributesData struct {
	CapabilityType string        `json:"capabilityType"`
	Settings       []SettingData `json:"settings,omitempty"`
}

type BundleIdForCapabilityData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type BundleIdRelationship struct {
	Data BundleIdForCapabilityData `json:"data"`
}

type RelationshipsData struct {
	BundleId BundleIdRelationship `json:"bundleId"`
}

type DataForCapability struct {
	Relationships RelationshipsData `json:"relationships"`
	Attributes    AttributesData    `json:"attributes"`
	Type          string            `json:"type"`
}

//Capability
