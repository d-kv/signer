package handlers

import (
	repo "d-kv/signer/db-common/entity"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapBundleIds(bundleIds []repo.BundleId) []entities.BundleId {
	result := make([]entities.BundleId, len(bundleIds))
	for _, bundleId := range bundleIds {
		result = append(result, entities.BundleId{
			Identifier: bundleId.ID,
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
	id := c.Param("id")

	repoBundleId, err := h.QueryProcessor.BundleIdRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bundle id not found"})
		return
	}

	bundleId := entities.BundleId{
		Identifier: id,
		Name:       repoBundleId.Name,
	}

	c.JSON(http.StatusOK, bundleId)
}
