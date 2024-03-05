package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	identification := router.Group("/:tenantId/:integrationId")
	{
		identification.POST("/bundleid", h.postBundleId)
		identification.GET("/bundleid", h.getBundleIds)
		identification.GET("/bundleid/:id", h.getBundleIdByID)
		identification.DELETE("/bundleid/:id", h.deleteBundleIdByID)

		identification.POST("/capability", h.postCapability)
		identification.DELETE("/capability/:id", h.deleteCapabilityByID)

		identification.POST("/devices", h.postDevice)
		identification.GET("/devices", h.getDevices)
		identification.GET("/devices/:id", h.getDevice)

		identification.POST("/certificates", h.postCertificate)
		identification.GET("/certificates", h.getCertificates)
		identification.GET("/certificates/:id", h.postCertificate)
		identification.DELETE("/certificates/:id", h.deleteCertificate)

		identification.POST("/profiles", h.postProfile)
		identification.GET("/profiles", h.getProfiles)
		identification.GET("/profiles/:id", h.getProfile)
		identification.DELETE("/profiles/:id", h.deleteProfile)
	}
	return router
}
