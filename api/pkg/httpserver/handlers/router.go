package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	bundleId := router.Group("/bundleid")
	{
		bundleId.POST("/", h.postBundleId)
		bundleId.GET("/", h.getBundleId)
	}
	devices := router.Group("/devices")
	{
		devices.POST("/", h.postDevice)
		devices.GET("/", h.getDevice)
	}
	certificates := router.Group("/certificates")
	{
		certificates.POST("/", h.postCertificate)
		certificates.GET("/", h.getCertificate)
	}
	profiles := router.Group("/profiles")
	{
		profiles.POST("/", h.postProfile)
		profiles.GET("/", h.getProfile)
	}
	return router
}
