package handlers

import (
	"d-kv/signer/api/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add new certificate
// @Tags certificate
// @Description Create new certificate instance
// @ID add-certificate
// @Accept json
// @Produce json
// @Param input body entities.Certificate true "certificate params"
// @Success 200 {object} entities.Certificate
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /certificates [post]
func (h *Handler) postCertificate(c *gin.Context) {
	var inp entities.Certificate

	err := c.BindJSON(&inp)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, inp)
}

// @Summary Get certificates list
// @Description Get all certificates
// @Tags certificate
// @ID get-certificates-list
// @Produce  json
// @Success 200 {array} []entities.Profile
// @Router /certificates [get]
func (h *Handler) getCertificates(c *gin.Context) {
	h.services.QueryHandlerService.GetCertificates(c)
}

// @Summary Get certificate
// @Description Get certificate instance by id
// @Tags certificate
// @ID get-certificate
// @Param id path string true "certificate identifier"
// @Produce  json
// @Success 200 {object} entities.Certificate
// @Router /certificates/{id} [get]
func (h *Handler) getCertificateByID(c *gin.Context) {
	h.services.QueryHandlerService.GetCertificateById(c)
}

// @Summary Delete certificate
// @Description Certificate deletion by id
// @Tags certificate
// @ID delete-certificate
// @Param id path string true "certificate identifier"
// @Produce  json
// @Success 204 {string} string "Deletion success"
// @Failure 404 {string} string "Certificate not found"
// @Router /certificates/{id} [delete]
func (h *Handler) deleteCertificateByID(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
