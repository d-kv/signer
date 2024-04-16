package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	repo "query-handler/internal/entity"
	"query-handler/pkg/httpserver/entities"
)

func mapBundleIds(bundleIds []repo.BundleId) []entities.BundleId {
	result := make([]entities.BundleId, len(bundleIds))
	for _, bundleId := range bundleIds {
		result = append(result, entities.BundleId{
			Identifier: parseStr(bundleId.ID),
			Name:       bundleId.Name,
		})
	}
	return result
}

func (h *Handler) getBundleIds(c *gin.Context) {
	repoBundleIds, err := h.QueryProcessor.BundleIdRepo.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bundleIds := mapBundleIds(repoBundleIds)
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
