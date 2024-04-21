package handlers

import (
	"d-kv/signer/db-common/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	QueryProcessor *services.QueryProcessor
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	identification := router.Group("/v1/query")
	{
		bundleId := identification.Group("/bundleIds")
		{
			bundleId.GET("/", h.getBundleIds)
			bundleId.GET("/:id", h.getBundleIdByID)
		}
		device := identification.Group("/devices")
		{
			device.GET("/", h.getDevices)
			device.GET("/:id", h.getDeviceByID)
		}
		certificate := identification.Group("/certificates")
		{
			certificate.GET("/", h.getCertificates)
			certificate.GET("/:id", h.getCertificateByID)
		}
		profile := identification.Group("/profiles")
		{
			profile.GET("/", h.getProfiles)
			profile.GET("/:id", h.getProfileByID)
		}

	}
	return router
}
