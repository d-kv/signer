package entities

import "d-kv/signer/db-common/entity"

type BundleId struct {
	ID           string                  `json:"id"`
	Identifier   string                  `json:"identifier"`
	Name         string                  `json:"name"`
	Capabilities []entity.CapabilityType `json:"capabilities"`
}
