package handlers

import (
	"d-kv/signer/api/pkg/httpserver/entities"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
	var _ entities.InputCreateCertificate
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
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /certificates/status/{id} [post]
func (h *Handler) getCertificateStatusByID(c *gin.Context) {
	status, err := h.services.CommandExecutorService.GetDeviceStatusByID(c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErrorResponse(c, http.StatusNotFound, "No commands with this status")
		} else {
			newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to CE")
		}
		return
	}
	c.JSON(http.StatusOK, status)
}

// @Summary Get certificates list
// @Description Get all certificates
// @Tags certificate
// @ID get-certificates-list
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Produce  json
// @Success 200 {array} []entities.Certificate
// @Failure 400,404 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /certificates [get]
func (h *Handler) getCertificates(c *gin.Context) {
	response, err := h.services.QueryHandlerService.GetCertificates(c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErrorResponse(c, http.StatusNotFound, "No records with this status")
		} else {
			newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to QH")
		}
		return
	}
	c.JSON(http.StatusOK, response)
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
// @Failure 400,404 {object} errorResponse
// @Failure 503 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /certificates/{id} [get]
func (h *Handler) getCertificateByID(c *gin.Context) {
	response, err := h.services.QueryHandlerService.GetCertificateById(c)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErrorResponse(c, http.StatusNotFound, "No records with this status")
		} else {
			newErrorResponse(c, http.StatusServiceUnavailable, "error while sending to QH")
		}
		return
	}
	c.JSON(http.StatusOK, response)
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
