package handlers

import (
	"github.com/gin-gonic/gin"
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
	//TODO implement me

}

// @Summary Get profiles list
// @Description Get all profiles
// @Tags profile
// @ID get-profiles-list
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Produce  json
// @Success 200 {array} []entities.Profile
// @Router /profiles [get]
func (h *Handler) getProfiles(c *gin.Context) {
	h.services.QueryHandlerService.GetProfiles(c)
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
// @Router /profiles/{id} [get]
func (h *Handler) getProfileByID(c *gin.Context) {
	h.services.QueryHandlerService.GetProfileById(c)
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
