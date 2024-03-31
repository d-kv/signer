package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"query-handler/pkg/httpserver/entities"
)

func (h *Handler) getCertificates(c *gin.Context) {
	certificates := []entities.Certificate{}

	c.JSON(http.StatusOK, certificates)
}

func (h *Handler) getCertificateByID(c *gin.Context) {
	id := c.Param("id")

	certificate := entities.Certificate{Identifier: id}

	c.JSON(http.StatusOK, certificate)
}
