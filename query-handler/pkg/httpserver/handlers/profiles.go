package handlers

import (
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getProfiles(c *gin.Context) {
	integrationId := c.Param("integrationId")
	profile, err := h.DomainRepos.ProfileRepo.FindByIntegrationId(c.Request.Context(), integrationId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, entities.Profile{
		ID:             profile.ID,
		ProfileContent: profile.ProfileContent,
	})
}

func (h *Handler) getProfileByID(c *gin.Context) {
	id := c.Param("id")
	profile, err := h.DomainRepos.ProfileRepo.FindById(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, entities.Profile{
		ID:             profile.ID,
		ProfileContent: profile.ProfileContent,
	})
}
