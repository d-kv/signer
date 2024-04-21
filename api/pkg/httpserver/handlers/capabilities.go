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
// @Param input body entities.Capability true "capability params"
// @Success 200 {object} entities.Capability
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /capabilities [post]
func (h *Handler) postCapability(c *gin.Context) {
	var inp entities.Capability

	err := c.BindJSON(&inp)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, inp)
}

// @Summary Delete capability
// @Description Capability deletion by id
// @Tags capability
// @ID delete-capability
// @Param id path string true "capability identifier"
// @Produce  json
// @Success 204 {string} string "Deletion success"
// @Failure 404 {string} string "Capability not found"
// @Router /capabilities/{id} [delete]
func (h *Handler) deleteCapabilityByID(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
