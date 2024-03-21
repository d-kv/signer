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
	id := c.Param("id")
	bundleId := entities.BundleId{Identifier: id}

	c.JSON(http.StatusOK, bundleId)
}
