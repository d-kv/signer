package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"query-handler/pkg/httpserver/entities"
)

func (h *Handler) getBundleIds(c *gin.Context) {
	bundleIds := []entities.BundleId{}
	c.JSON(http.StatusOK, bundleIds)
}

func (h *Handler) getBundleIdByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := parseUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	repoBundleId, err := h.QueryProcessor.BundleIdRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bundle id not found"})
		return
	}

	bundleId := entities.BundleId{
		Identifier: idStr,
		Name:       repoBundleId.Name,
	}

	c.JSON(http.StatusOK, bundleId)
}
