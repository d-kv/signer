package handlers

import (
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getCertificates(c *gin.Context) {
	integrationId := c.Param("integrationId")
	certificate, err := h.DomainRepos.CertificateRepo.FindByIntegrationId(c.Request.Context(), integrationId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, entities.Certificate{
		ID:                 certificate.ID,
		CertificateContent: certificate.CertificateContent,
	})
}

func (h *Handler) getCertificateByID(c *gin.Context) {
	id := c.Param("id")
	certificate, err := h.DomainRepos.CertificateRepo.FindById(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, entities.Certificate{
		ID:                 certificate.ID,
		CertificateContent: certificate.CertificateContent,
	})
}
