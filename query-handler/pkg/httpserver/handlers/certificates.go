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
	idStr := c.Param("id")

	id, err := parseUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be integer"})
		return
	}

	repoCertificate, err := h.QueryProcessor.CertificateRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "certificate not found"})
		return
	}

	certificate := entities.Certificate{
		Identifier: idStr,
		Name:       repoCertificate.Name,
		Type:       repoCertificate.Type,
	}

	c.JSON(http.StatusOK, certificate)
}
