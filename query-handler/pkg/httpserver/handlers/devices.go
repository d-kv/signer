package handlers

import (
	repo "d-kv/signer/db-common/entity"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapDevices(devices []repo.Device) []entities.Device {
	var result []entities.Device
	for _, device := range devices {
		result = append(result, entities.Device{
			Identifier: device.ID,
			Name:       device.Name,
			UserId:     device.User.ID,
		})
	}
	return result
}

func (h *Handler) getDevices(c *gin.Context) {
	repoDevices, err := h.QueryProcessor.DeviceRepo.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	devices := mapDevices(repoDevices)
	c.JSON(http.StatusOK, devices)
}

func (h *Handler) getDeviceByID(c *gin.Context) {
	id := c.Param("id")

	repoDevice, err := h.QueryProcessor.DeviceRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
		return
	}

	userId := repoDevice.User.ID

	device := entities.Device{
		Identifier: id,
		Name:       repoDevice.Name,
		UserId:     userId,
	}

	c.JSON(http.StatusOK, device)
}
