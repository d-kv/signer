package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	repo "query-handler/internal/entity"
	"query-handler/pkg/httpserver/entities"
)

func mapDevices(devices []repo.Device) []entities.Device {
	var result []entities.Device
	for _, device := range devices {
		result = append(result, entities.Device{
			Identifier: parseStr(device.ID),
			Name:       device.Name,
			UserId:     parseStr(device.User.ID),
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
	idStr := c.Param("id")

	id, err := parseUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	repoDevice, err := h.QueryProcessor.DeviceRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
		return
	}

	userId := parseStr(repoDevice.User.ID)

	device := entities.Device{
		Identifier: idStr,
		Name:       repoDevice.Name,
		UserId:     userId,
	}

	c.JSON(http.StatusOK, device)
}
