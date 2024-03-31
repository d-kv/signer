package entities

type Profile struct {
	Identifier    string `json:"identifier"`
	Name          string `json:"name"`
	BundleId      string `json:"bundleId"`
	IntegrationId string `json:"integrationId"`
}
