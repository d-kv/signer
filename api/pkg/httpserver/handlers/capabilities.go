package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary Add new capability
// @Tags capability
// @Description Create new capability instance
// @ID add-capability
// @Accept json
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param input body entities.Capability true "capability params"
// @Success 200 {object} entities.Capability
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /capabilities [post]
func (h *Handler) postCapability(c *gin.Context) {
	h.services.CommandExecutorService.PostCapability(c)
}

func (h *Handler) getCapabilityStatusByID(c *gin.Context) {

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
	h.services.CommandExecutorService.DelCapability(c)
}
