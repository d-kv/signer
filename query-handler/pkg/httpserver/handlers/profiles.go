package handlers

import (
	repo "d-kv/signer/db-common/entity"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapProfiles(profiles []repo.Profile) []entities.Profile {
	var result []entities.Profile
	for _, profile := range profiles {
		result = append(result, entities.Profile{
			Identifier:    profile.ID,
			Name:          profile.Name,
			BundleId:      profile.BundleId.ID,
			IntegrationId: profile.Integration.ID,
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
	id := c.Param("id")

	repoProfile, err := h.QueryProcessor.ProfileRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	bundleId := repoProfile.BundleId.ID
	integrationId := repoProfile.Integration.ID

	profile := entities.Profile{
		Identifier:    id,
		Name:          repoProfile.Name,
		BundleId:      bundleId,
		IntegrationId: integrationId,
	}

	c.JSON(http.StatusOK, profile)
}
