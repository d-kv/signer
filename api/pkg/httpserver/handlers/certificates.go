package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary (NOT IMPLEMENTED) Add new certificate
// @Tags certificate
// @Description Create new certificate instance
// @ID add-certificate
// @Accept json
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param input body entities.InputCreateCertificate true "certificate params"
// @Success 200 {string} string "commandID"
// @Failure 400 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /certificates [post]
func (h *Handler) postCertificate(c *gin.Context) {
	//TODO implement me
}

// @Summary Check status
// @Tags certificate
// @Description Get command status by command id
// @ID get-status-certificate
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "command identifier"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /certificates/status/{id} [post]
func (h *Handler) getCertificateStatusByID(c *gin.Context) {
	//TODO implement me
}

// @Summary Get certificates list
// @Description Get all certificates
// @Tags certificate
// @ID get-certificates-list
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Produce  json
// @Success 200 {array} []entities.Certificate
// @Router /certificates [get]
func (h *Handler) getCertificates(c *gin.Context) {
	h.services.QueryHandlerService.GetCertificates(c)
}

// @Summary Get certificate
// @Description Get certificate instance by id
// @Tags certificate
// @ID get-certificate
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "certificate identifier"
// @Produce  json
// @Success 200 {object} entities.Certificate
// @Router /certificates/{id} [get]
func (h *Handler) getCertificateByID(c *gin.Context) {
	h.services.QueryHandlerService.GetCertificateById(c)
}

// @Summary (NOT IMPLEMENTED) Delete certificate
// @Description Certificate deletion by id
// @Tags certificate
// @ID delete-certificate
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "certificate identifier"
// @Produce  json
// @Success 204 {string} string "Deletion success"
// @Failure 404 {string} string "Certificate not found"
// @Router /certificates/{id} [delete]
func (h *Handler) deleteCertificateByID(c *gin.Context) {
	//TODO implement me
}
