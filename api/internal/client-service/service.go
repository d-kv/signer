package client_service

import (
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
	PostBundleId(c *gin.Context)
	DelBundleIdById(c *gin.Context)
	PostCapability(c *gin.Context)
	DelCapability(c *gin.Context)
	PostDevice(c *gin.Context)
}

type Service struct {
	QueryHandlerService
	CommandExecutorService
}

func NewService() *Service {
	return &Service{
		QueryHandlerService: NewQueryService(),
	}
}
