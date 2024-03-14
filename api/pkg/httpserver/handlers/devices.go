package handlers

import (
	"api/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add new device
// @Tags device
// @Description Create new device instance
// @ID add-device
// @Accept json
// @Produce json
// @Param input body entities.Device true "device params"
// @Success 200 {object} entities.Device
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /devices [post]
func (h *Handler) postDevice(c *gin.Context) {
	var inp entities.Device

	err := c.BindJSON(&inp)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, inp)
}

// @Summary Get devices list
// @Description Get all devices
// @Tags device
// @ID get-devices-list
// @Produce  json
// @Success 200 {array} []entities.Device
// @Router /devices [get]
func (h *Handler) getDevices(c *gin.Context) {
	h.services.QueryHandlerService.GetDevices(c)
}

// @Summary Get device
// @Description Get device instance by id
// @Tags device
// @ID get-device
// @Param id path string true "device identifier"
// @Produce  json
// @Success 200 {object} entities.Device
// @Router /devices/{id} [get]
func (h *Handler) getDeviceByID(c *gin.Context) {
	h.services.QueryHandlerService.GetDeviceById(c)
}
