package client_service

import (
	"d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	"github.com/gin-gonic/gin"
)

type QueryHandlerService interface {
	GetBundleIds(c *gin.Context)
	GetBundleIdById(c *gin.Context)

	GetDevices(c *gin.Context)
	GetDeviceById(c *gin.Context)

	GetCertificates(c *gin.Context)
	GetCertificateById(c *gin.Context)

	GetProfiles(c *gin.Context)
	GetProfileById(c *gin.Context)
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
