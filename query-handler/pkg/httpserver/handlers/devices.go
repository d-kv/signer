package handlers

import (
	"d-kv/signer/db-common/entity"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func (h *Handler) getDevices(c *gin.Context) {
	integrationId := c.Param("integrationId")
	devices, err := h.DomainRepos.DeviceRepo.FindByIntegrationId(c.Request.Context(), integrationId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapDevices(devices))
}

func mapDevices(devices []entity.Device) []entities.Device {
	var results []entities.Device
	for _, device := range devices {
		results = append(results, entities.Device{
			UDID:     device.UDID,
			Name:     device.Name,
			UserId:   device.UserID,
			Platform: mapStringPlatform(device.Platform),
		})
	}
	return results
}

func mapStringPlatform(platform string) entity.Platform {
	if !slices.Contains(entity.PlatformValues, platform) {
		return "Undefined"
	}
	return entity.Platform(platform)
}

func (h *Handler) getDeviceByID(c *gin.Context) {
	id := c.Param("id")
	device, err := h.DomainRepos.DeviceRepo.FindById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entities.Device{
		UDID:     device.UDID,
		Name:     device.Name,
		UserId:   device.UserID,
		Platform: mapStringPlatform(device.Platform),
	})
}
