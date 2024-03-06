package handlers

import (
	_ "api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	identification := router.Group("/v1/:tenantId/:integrationId")
	{
		bundleId := identification.Group("/bundleIds")
		{
			bundleId.POST("/", h.postBundleId)
			bundleId.GET("/", h.getBundleIds)
			bundleId.GET("/:id", h.getBundleIdByID)
			bundleId.DELETE("/:id", h.deleteBundleIdByID)
		}
		capability := identification.Group("/capabilities")
		{
			capability.POST("/", h.postCapability)
			capability.DELETE("/:id", h.deleteCapabilityByID)
		}
		device := identification.Group("/devices")
		{
			device.POST("/", h.postDevice)
			device.GET("/", h.getDevices)
			device.GET("/:id", h.getDeviceByID)
		}
		certificate := identification.Group("/certificates")
		{
			certificate.POST("/", h.postCertificate)
			certificate.GET("/", h.getCertificates)
			certificate.GET("/:id", h.getCertificateByID)
			certificate.DELETE("/:id", h.deleteCertificateByID)
		}
		profile := identification.Group("/profiles")
		{
			profile.POST("/", h.postProfile)
			profile.GET("/", h.getProfiles)
			profile.GET("/:id", h.getProfileByID)
			profile.DELETE("/:id", h.deleteProfileByID)
		}

	}
	return router
}
