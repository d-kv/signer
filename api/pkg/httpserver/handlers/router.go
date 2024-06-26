package handlers

import (
	_ "d-kv/signer/api/docs"
	"d-kv/signer/api/internal/client-service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *client_service.Service
}

func NewHandler(services *client_service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	identification := router.Group("/v1/:tenantId/:integrationId")
	{
		bundleId := identification.Group("/bundleIds")
		{
			bundleId.POST("/", h.postBundleId)
			bundleId.GET("/status/:id", h.getBundleIdStatusByID)
			bundleId.GET("/", h.getBundleIds)
			bundleId.GET("/:id", h.getBundleIdByID)
			//bundleId.DELETE("/:id", h.deleteBundleIdByID)
		}
		capability := identification.Group("/capabilities")
		{
			capability.POST("/", h.postCapability)
			capability.GET("/status/:id", h.getCapabilityStatusByID)
			//capability.DELETE("/:id", h.deleteCapabilityByID)
		}
		device := identification.Group("/devices")
		{
			device.POST("/", h.postDevice)
			device.GET("/status/:id", h.getDeviceStatusByID)
			device.GET("/", h.getDevices)
			device.GET("/:id", h.getDeviceByID)
		}
		certificate := identification.Group("/certificates")
		{
			certificate.POST("/", h.postCertificate)
			certificate.GET("/status/:id", h.getCertificateStatusByID)
			certificate.GET("/", h.getCertificates)
			certificate.GET("/:id", h.getCertificateByID)
			//certificate.DELETE("/:id", h.deleteCertificateByID)
		}
		profile := identification.Group("/profiles")
		{
			profile.POST("/", h.postProfile)
			profile.GET("/status/:id", h.getProfileStatusByID)
			profile.GET("/", h.getProfiles)
			profile.GET("/:id", h.getProfileByID)
			//profile.DELETE("/:id", h.deleteProfileByID)
		}

	}
	return router
}
