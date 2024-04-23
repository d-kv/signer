package handlers

import (
	_ "d-kv/signer/db-common/entity"
	"github.com/gin-gonic/gin"
)

// @Summary Add new BundleId
// @Tags bundleId
// @Description Create new bundleId instance
// @ID add-bundleId
// @Accept json
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param input body entities.BundleId true "bundleId params"
// @Success 200 {object} entities.BundleId
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /bundleIds [post]
func (h *Handler) postBundleId(c *gin.Context) {
	h.services.CommandExecutorService.PostBundleId(c)
}

// @Summary Get bundleIds list
// @Description Get all bundleIds
// @Tags bundleId
// @ID get-bundleIds-list
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
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
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
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
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "bundleId identifier"
// @Produce  json
// @Success 204 {string} string "Deletion success"
// @Failure 404 {string} string "BundleId not found"
// @Router /bundleId/{id} [delete]
func (h *Handler) deleteBundleIdByID(c *gin.Context) {
	h.services.CommandExecutorService.DelBundleIdById(c)
}
