package handlers

import (
	"d-kv/signer/api/pkg/httpserver/entities"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// @Summary (NOT IMPLEMENTED) Add new profile
// @Tags profile
// @Description Create new profile instance
// @ID add-profile
// @Accept json
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param input body entities.Profile true "profile params"
// @Success 200 {object} entities.Profile
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /profiles [post]
func (h *Handler) postProfile(c *gin.Context) {
	var input entities.InputCreateProfile
	tenantId := c.Param("tenantId")
	integrationId := c.Param("integrationId")
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	converted, err := entities.ConvertProfile(&input, tenantId, integrationId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	opId, err := h.services.CommandExecutorService.PostProfile(c, converted)
	if err != nil {
		newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to CE")
		return
	}
	c.JSON(http.StatusOK, opId)
}

// @Summary Check status
// @Tags profile
// @Description Get command status by command id
// @ID get-status-profile
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "command identifier"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /profiles/status/{id} [post]
func (h *Handler) getProfileStatusByID(c *gin.Context) {
	status, err := h.services.CommandExecutorService.GetProfileStatusByID(c)
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

// @Summary Get profiles list
// @Description Get all profiles
// @Tags profile
// @ID get-profiles-list
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Produce  json
// @Success 200 {array} []entities.Profile
// @Failure 400,404 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /profiles [get]
func (h *Handler) getProfiles(c *gin.Context) {
	response, err := h.services.QueryHandlerService.GetProfiles(c)
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

// @Summary Get profile
// @Description Get profile instance by id
// @Tags profile
// @ID get-profile
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "profile identifier"
// @Produce  json
// @Success 200 {object} entities.Profile
// @Failure 400,404 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /profiles/{id} [get]
func (h *Handler) getProfileByID(c *gin.Context) {
	response, err := h.services.QueryHandlerService.GetProfileById(c)
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

// @Summary (NOT IMPLEMENTED) Delete profile
// @Description Profile deletion by id
// @Tags profile
// @ID delete-profile
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "profile identifier"
// @Produce  json
// @Success 204 {string} string "Deletion success"
// @Failure 404 {string} string "Profile not found"
// @Router /profiles/{id} [delete]
func (h *Handler) deleteProfileByID(c *gin.Context) {
	//TODO implement me
}
