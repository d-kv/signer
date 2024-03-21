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
	id := c.Param("id")

	profile := entities.Profile{Identifier: id}

	c.JSON(http.StatusOK, profile)
}
