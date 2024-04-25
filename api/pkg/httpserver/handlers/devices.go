package handlers

import (
	"d-kv/signer/api/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add new device
// @Tags device
// @Description Create new device instance
// @ID add-device
// @Accept json
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param input body entities.Device true "device params"
// @Success 200 {string} string "commandID"
// @Failure 400 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /devices [post]
func (h *Handler) postDevice(c *gin.Context) {
	var input entities.InputCreateDevice
	tenantId := c.Param("tenantId")
	integrationId := c.Param("integrationId")
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	bundleInput, err := entities.ConvertDevice(&input, tenantId, integrationId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	opId, err := h.services.CommandExecutorService.PostDevice(c, bundleInput)
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to QH")
		return
	}
	c.JSON(http.StatusOK, opId)
}

// @Summary Check status
// @Tags device
// @Description Get command status by command id
// @ID get-status-device
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "command identifier"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /devices/status/{id} [post]
func (h *Handler) getDeviceStatusByID(c *gin.Context) {

}

// @Summary Get devices list
// @Description Get all devices
// @Tags device
// @ID get-devices-list
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
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
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "device identifier"
// @Produce  json
// @Success 200 {object} entities.Device
// @Router /devices/{id} [get]
func (h *Handler) getDeviceByID(c *gin.Context) {
	h.services.QueryHandlerService.GetDeviceById(c)
}
