package entity

type IntegrationToken struct {
	IntegrationId string `json:"integration_id"`
	Token         string `json:"token"`
	KeyId         string `json:"key_id"`
}
