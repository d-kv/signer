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
	id := c.Param("id")

	device := entities.Device{Identifier: id, Name: "TestName"}

	c.JSON(http.StatusOK, device)
}
