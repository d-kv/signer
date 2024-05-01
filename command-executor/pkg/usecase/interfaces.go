package usecase

import "d-kv/signer/command-executor/pkg/entity"

type DataBaseCommand interface {
	Convert() entity.ApiEntity
	GetId() uint
}
