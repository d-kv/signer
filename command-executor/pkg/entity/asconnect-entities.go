package entity

import "time"

type ApiEntity interface {
	GetURL() URL
}

//--------------------------------------

type BundleIdResponse struct {
	Id string `json:"id"`
}

// ApiCreateBundleId AppstoreConnectAPI body
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

func (b *ApiCreateBundleId) GetURL() URL {
	return BundleIdsURL
}

//--------------------------------------

// ApiCreateDevice AppstoreConnectAPI body
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

func (b *ApiCreateDevice) GetURL() URL {
	return DevicesURL
}

//--------------------------------------

// ApiEnableCapability AppstoreConnectAPI body
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

func (b *ApiEnableCapability) GetURL() URL {
	return CapabilitiesURL
}

//--------------------------------------

// ProfileResponse Response
type ProfileResponse struct {
	Data ProfileResponseData `json:"data"`
}

type ProfileResponseAttributes struct {
	Name           string    `json:"name"`
	Platform       string    `json:"platform"`
	ProfileType    string    `json:"profileType"`
	ProfileState   string    `json:"profileState"`
	ProfileContent string    `json:"profileContent"`
	UUID           string    `json:"uuid"`
	CreatedDate    time.Time `json:"createdDate"`
	ExpirationDate time.Time `json:"expirationDate"`
}

type ProfileResponseData struct {
	ID         string                    `json:"id"`
	Type       string                    `json:"type"`
	Attributes ProfileResponseAttributes `json:"attributes"`
}

// ApiCreateProfile AppstoreConnectAPI body
type ApiCreateProfile struct {
	Data ProfileData `json:"data"`
}
type ProfileData struct {
	Type          string               `json:"type"`
	Attributes    ProfileAttributes    `json:"attributes"`
	Relationships ProfileRelationships `json:"relationships"`
}

type ProfileAttributes struct {
	Name        string `json:"name"`
	ProfileType string `json:"profileType"`
}

type ProfileRelationships struct {
	BundleID     ProfileArgument `json:"bundleId"`
	Certificates ProfileArgument `json:"certificates"`
	Devices      ProfileArgument `json:"devices"`
}

type ProfileArgument struct {
	Data []ProfileArgumentData `json:"data"`
}
type ProfileArgumentData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (b *ApiCreateProfile) GetURL() URL {
	return ProfilesURL
}

//--------------------------------------

// CertificateResponse Response
type CertificateResponse struct {
	Data CertificateResponseData `json:"data"`
}

type CertificateResponseAttributes struct {
	Name               string    `json:"name"`
	CertificateType    string    `json:"certificateType"`
	DisplayName        string    `json:"displayName"`
	SerialNumber       string    `json:"serialNumber"`
	Platform           string    `json:"platform"`
	ExpirationDate     time.Time `json:"expirationDate"`
	CertificateContent string    `json:"certificateContent"`
}

type CertificateResponseData struct {
	ID         string                        `json:"id"`
	Type       string                        `json:"type"`
	Attributes CertificateResponseAttributes `json:"attributes"`
}

// ApiAddCertificate AppstoreConnectAPI body
type ApiAddCertificate struct {
	Data CertificateData `json:"data"`
}

type CertificateData struct {
	Attributes CertificateAttributes `json:"attributes"`
	Type       string                `json:"type"`
}

type CertificateAttributes struct {
	CsrContent      string `json:"csrContent"`
	CertificateType string `json:"certificateType"`
}

func (b *ApiAddCertificate) GetURL() URL {
	return CertificatesURL
}
