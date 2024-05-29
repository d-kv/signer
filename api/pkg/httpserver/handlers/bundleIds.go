package handlers

import (
	"d-kv/signer/api/pkg/httpserver/entities"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// @Summary Add new BundleId
// @Tags bundleId
// @Description Create new bundleId instance
// @ID add-bundleId
// @Accept json
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param input body entities.InputCreateBundleId true "bundleId params"
// @Success 200 {string} string "commandID"
// @Failure 400 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /bundleIds [post]
func (h *Handler) postBundleId(c *gin.Context) {
	var input entities.InputCreateBundleId
	tenantId := c.Param("tenantId")
	integrationId := c.Param("integrationId")
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	converted, err := entities.ConvertBundleInput(&input, tenantId, integrationId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	opId, err := h.services.CommandExecutorService.PostBundleId(c, converted)
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to CE")
		return
	}
	c.JSON(http.StatusOK, opId)
}

// @Summary Check status
// @Tags bundleId
// @Description Get command status by command id
// @ID get-status-bundleId
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "command identifier"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /bundleIds/status/{id} [post]
func (h *Handler) getBundleIdStatusByID(c *gin.Context) {
	status, err := h.services.CommandExecutorService.GetBundleIdStatusByID(c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErrorResponse(c, http.StatusNotFound, "No commands with this status")
		} else {
			newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to CE")
		}
		return
	}
	c.JSON(http.StatusOK, status)
}

// @Summary Get bundleIds list
// @Description Get all bundleIds
// @Tags bundleId
// @ID get-bundleIds-list
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Produce  json
// @Success 200 {array} []entities.BundleId
// @Failure 400,404 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /bundleIds [get]
func (h *Handler) getBundleIds(c *gin.Context) {
	response, err := h.services.QueryHandlerService.GetBundleIds(c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErrorResponse(c, http.StatusNotFound, "No records with this status")
		} else {
			newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to QH")
		}
		return
	}
	c.JSON(http.StatusOK, response)
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
// @Failure 400,404 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /bundleIds/{id} [get]
func (h *Handler) getBundleIdByID(c *gin.Context) {
	response, err := h.services.QueryHandlerService.GetBundleIdById(c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErrorResponse(c, http.StatusNotFound, "No records with this status")
		} else {
			newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to QH")
		}
		return
	}
	c.JSON(http.StatusOK, response)
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
}
