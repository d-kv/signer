package handlers

import (
	"d-kv/signer/api/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add new capability
// @Tags capability
// @Description Create new capability instance
// @ID add-capability
// @Accept json
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param input body entities.InputEnableCapability true "capability params"
// @Success 200 {object} string "commandID"
// @Failure 400 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /capabilities [post]
func (h *Handler) postCapability(c *gin.Context) {
	var input entities.InputEnableCapability
	tenantId := c.Param("tenantId")
	integrationId := c.Param("integrationId")
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	bundleInput, err := entities.ConvertCapability(&input, tenantId, integrationId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	opId, err := h.services.CommandExecutorService.PostCapability(c, bundleInput)
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to QH")
		return
	}
	c.JSON(http.StatusOK, opId)
}

// @Summary Check status
// @Tags capability
// @Description Get command status by command id
// @ID get-status-capability
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "command identifier"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /capabilities/status/{id} [post]
func (h *Handler) getCapabilityStatusByID(c *gin.Context) {
	status, err := h.services.CommandExecutorService.GetCapabilityStatusByID(c)
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to QH")
		return
	}
	c.JSON(http.StatusOK, status)
}

// @Summary Delete capability
// @Description Capability deletion by id
// @Tags capability
// @ID delete-capability
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "capability identifier"
// @Produce  json
// @Success 204 {string} string "Deletion success"
// @Failure 404 {string} string "Capability not found"
// @Router /capabilities/{id} [delete]
func (h *Handler) deleteCapabilityByID(c *gin.Context) {
}
