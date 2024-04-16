package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"query-handler/pkg/httpserver/entities"
)

func (h *Handler) getDevices(c *gin.Context) {
	devices := []entities.Device{}
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
