package handlers

import (
	"d-kv/signer/api/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add new BundleId
// @Tags bundleId
// @Description Create new bundleId instance
// @ID add-bundleId
// @Accept json
// @Produce json
// @Param input body entities.BundleId true "bundleId params"
// @Success 200 {object} entities.BundleId
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /bundleIds [post]
func (h *Handler) postBundleId(c *gin.Context) {
	var inp entities.BundleId

	err := c.BindJSON(&inp)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, inp)
}

// @Summary Get bundleIds list
// @Description Get all bundleIds
// @Tags bundleId
// @ID get-bundleIds-list
// @Produce  json
// @Success 200 {array} []entities.BundleId
// @Router /bundleIds [get]
func (h *Handler) getBundleIds(c *gin.Context) {
	h.services.QueryHandlerService.GetBundleIds(c)
}

// @Summary Get bundleId
// @Description Get bundleId instance by id
// @Tags bundleId
// @ID get-bundleId
// @Param id path string true "bundleId identifier"
// @Produce  json
// @Success 200 {object} entities.BundleId
// @Router /bundleIds/{id} [get]
func (h *Handler) getBundleIdByID(c *gin.Context) {
	h.services.QueryHandlerService.GetBundleIdById(c)
}

// @Summary Delete bundleId
// @Description BundleId deletion by id
// @Tags bundleId
// @ID delete-bundleId
// @Param id path string true "bundleId identifier"
// @Produce  json
// @Success 204 {string} string "Deletion success"
// @Failure 404 {string} string "BundleId not found"
// @Router /bundleId/{id} [delete]
func (h *Handler) deleteBundleIdByID(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
