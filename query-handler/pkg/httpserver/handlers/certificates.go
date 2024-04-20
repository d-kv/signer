package handlers

import (
	repo "d-kv/signer/query-handler/internal/entity"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapCertificates(certificates []repo.Certificate) []entities.Certificate {
	var result []entities.Certificate
	for _, cert := range certificates {
		result = append(result, entities.Certificate{
			Identifier: parseStr(cert.ID),
			Name:       cert.Name,
			Type:       cert.Type,
		})
	}
	return result
}

func (h *Handler) getCertificates(c *gin.Context) {
	repoCertificates, err := h.QueryProcessor.CertificateRepo.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	certificates := mapCertificates(repoCertificates)
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
