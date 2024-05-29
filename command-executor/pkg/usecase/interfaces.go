package usecase

import (
	"context"
	"crypto/ecdsa"
	"d-kv/signer/command-executor/pkg/entity"
	dbEntity "d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/usecase"
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --name=ProcessorService
type ProcessorService interface {
	SetStatusById(ctx context.Context, baseCommand *DataBaseCommand, status dbEntity.Status) error
	Processing(ctx context.Context, operation DataBaseCommand) (*http.Response, error)
	StartProcessor(ctx context.Context)
}

//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --name=ApiService
type ApiService interface {
	SendCreateCommand(ctx context.Context, e entity.ApiEntity, IntegrationId string) (*http.Response, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --name=TokenService
type TokenService interface {
	GenerateECDSAPrivateKey(base64PrivateKey string) (*ecdsa.PrivateKey, error)
	GetJwtToken(tokenInfo *dbEntity.IntegrationToken) (string, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --name=DataBaseCommand
type DataBaseCommand interface {
	GetId() uint
	Convert() entity.ApiEntity
	GetIntegrationId() string
}

//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --name=DataBaseService
type DataBaseService interface {
	WriteCapability(ctx context.Context, err error, operation dbEntity.EnableCapabilityType) error
	WriteDevice(ctx context.Context, operation dbEntity.CreateDevice) error
	WriteBundleId(ctx context.Context, operation dbEntity.CreateBundleId, response *entity.BundleIdResponse) error
	WriteProfile(ctx context.Context, operation dbEntity.CreateProfile, resp *entity.ProfileResponse) error
	WriteCertificate(ctx context.Context, operation dbEntity.CreateCertificate, resp *entity.CertificateResponse) error
	UpdateArrayDevices(ctx context.Context, devices []dbEntity.Device, converted *dbEntity.Profile, err error) error
	UpdateArrayCertificates(ctx context.Context, certificates []dbEntity.Certificate, converted *dbEntity.Profile, err error) error
}

type Service struct {
	Queue usecase.CommandRepo
	Vault usecase.VaultRepo
	Api   ApiService
	Token TokenService
	Db    DataBaseService
}

func NewService(queue usecase.CommandRepo, vault usecase.VaultRepo, api ApiService, token TokenService, db DataBaseService) *Service {
	return &Service{
		Queue: queue,
		Vault: vault,
		Api:   api,
		Token: token,
		Db:    db,
	}
}
