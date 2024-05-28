package client_service

import (
	"d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
)

type QueryHandlerService interface {
	GetBundleIds(c *gin.Context) ([]entities.BundleId, error)
	GetBundleIdById(c *gin.Context) (entities.BundleId, error)

	GetDevices(c *gin.Context) ([]entities.Device, error)
	GetDeviceById(c *gin.Context) (entities.Device, error)

	GetCertificates(c *gin.Context) ([]entities.Certificate, error)
	GetCertificateById(c *gin.Context) (entities.Certificate, error)

	GetProfiles(c *gin.Context) ([]entities.Profile, error)
	GetProfileById(c *gin.Context) (entities.Profile, error)
}

type CommandExecutorService interface {
	PostBundleId(c *gin.Context, ent *entity.CreateBundleId) (uint, error)
	GetBundleIdStatusByID(c *gin.Context) (entity.Status, error)
	DelBundleIdById(c *gin.Context) (uint, error)

	PostCapability(c *gin.Context, ent *entity.EnableCapabilityType) (uint, error)
	GetCapabilityStatusByID(c *gin.Context) (entity.Status, error)
	DelCapability(c *gin.Context) (uint, error)

	PostDevice(c *gin.Context, ent *entity.CreateDevice) (uint, error)
	GetDeviceStatusByID(c *gin.Context) (entity.Status, error)

	PostProfile(c *gin.Context, ent *entity.CreateProfile) (uint, error)
	GetProfileStatusByID(c *gin.Context) (entity.Status, error)
	DelProfileById(c *gin.Context) (uint, error)

	PostCertificate(c *gin.Context, ent *entity.CreateCertificate) (uint, error)
	GetCertificateStatusByID(c *gin.Context) (entity.Status, error)
	DelCertificateById(c *gin.Context) (uint, error)
}

type Service struct {
	QueryHandlerService
	CommandExecutorService
}

func NewService(queue *command.Repo, queryURL string) *Service {
	return &Service{
		QueryHandlerService:    NewQueryService(queryURL),
		CommandExecutorService: NewCommandService(queue),
	}
}
