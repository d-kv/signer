package handlers

import (
	repo "d-kv/signer/query-handler/internal/entity"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapProfiles(profiles []repo.Profile) []entities.Profile {
	var result []entities.Profile
	for _, profile := range profiles {
		result = append(result, entities.Profile{
			Identifier:    parseStr(profile.ID),
			Name:          profile.Name,
			BundleId:      parseStr(profile.BundleId.ID),
			IntegrationId: parseStr(profile.Integration.ID),
		})
	}
	return result
}

func (h *Handler) getProfiles(c *gin.Context) {
	repoProfiles, err := h.QueryProcessor.ProfileRepo.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	profiles := mapProfiles(repoProfiles)
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
