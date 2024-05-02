package entities

import "d-kv/signer/db-common/entity"

type Device struct {
	Identifier string          `json:"identifier"`
	Name       string          `json:"name"`
	UserId     string          `json:"userId"`
	UDID       string          `json:"udid"`
	Platform   entity.Platform `json:"platform"`
}
