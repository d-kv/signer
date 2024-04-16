package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"query-handler/pkg/httpserver/entities"
)

func (h *Handler) getProfiles(c *gin.Context) {
	profiles := []entities.Profile{}

	c.JSON(http.StatusOK, profiles)
}

func (h *Handler) getProfileByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := parseUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	repoProfile, err := h.QueryProcessor.ProfileRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	bundleId := parseStr(repoProfile.BundleId.ID)
	integrationId := parseStr(repoProfile.Integration.ID)

	profile := entities.Profile{
		Identifier:    idStr,
		Name:          repoProfile.Name,
		BundleId:      bundleId,
		IntegrationId: integrationId,
	}

	c.JSON(http.StatusOK, profile)
}
