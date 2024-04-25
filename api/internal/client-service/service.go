package client_service

import (
	"d-kv/signer/db-common/entity"
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
	PostBundleId(c *gin.Context, ent *entity.CreateBundleId) error
	GetBundleIdStatusByID(c *gin.Context) (entity.Status, error)
	DelBundleIdById(c *gin.Context) error
	PostCapability(c *gin.Context) error
	GetCapabilityStatusByID(c *gin.Context) (entity.Status, error)
	DelCapability(c *gin.Context) error
	PostDevice(c *gin.Context) error
	GetDeviceStatusByID(c *gin.Context) (entity.Status, error)
}

type Service struct {
	QueryHandlerService
	CommandExecutorService
}

func NewService() *Service {
	return &Service{
		QueryHandlerService: NewQueryService(),
		//CommandExecutorService: NewCommandService(),
	}
}
