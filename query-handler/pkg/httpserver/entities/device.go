package entities

import "d-kv/signer/db-common/entity"

type Device struct {
	UDID       string          `json:"udid"`
	Name       string          `json:"name"`
	UserId     string          `json:"userId"`
	Identifier string          `json:"id"`
	Platform   entity.Platform `json:"platform"`
}
