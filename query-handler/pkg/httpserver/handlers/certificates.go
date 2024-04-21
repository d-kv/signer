package handlers

import (
	repo "d-kv/signer/db-common/entity"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapCertificates(certificates []repo.Certificate) []entities.Certificate {
	var result []entities.Certificate
	for _, cert := range certificates {
		result = append(result, entities.Certificate{
			Identifier: cert.ID,
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
	id := c.Param("id")

	repoCertificate, err := h.QueryProcessor.CertificateRepo.FindById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "certificate not found"})
		return
	}

	certificate := entities.Certificate{
		Identifier: id,
		Name:       repoCertificate.Name,
		Type:       repoCertificate.Type,
	}

	c.JSON(http.StatusOK, certificate)
}
