package handlers

import (
	"d-kv/signer/api/pkg/httpserver/entities"
	"d-kv/signer/db-common/entity"
	_ "d-kv/signer/db-common/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add new BundleId
// @Tags bundleId
// @Description Create new bundleId instance
// @ID add-bundleId
// @Accept json
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param input body entities.InputCreateBundleId true "bundleId params"
// @Success 200 {string} string "commandID"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /bundleIds [post]
func (h *Handler) postBundleId(c *gin.Context) {
	var input entities.InputCreateBundleId
	tenantId := c.Param("tenantId")
	integrationId := c.Param("integrationId")
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	bundleInput, err := ConvertBundleInput(&input, tenantId, integrationId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	err = h.services.CommandExecutorService.PostBundleId(c, bundleInput)
	if err != nil {
		return
	}
}

func ConvertBundleInput(input *entities.InputCreateBundleId, tenantId string, integrationId string) (*entity.CreateBundleId, error) {
	newCommand := &entity.CreateBundleId{}
	return newCommand, nil
}

// @Summary Check status
// @Tags bundleId
// @Description Get command status by command id
// @ID get-status-bundleId
// @Produce json
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "command identifier"
// @Success 200 {string} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /bundleIds/status/{id} [post]
func (h *Handler) getBundleIdStatusByID(c *gin.Context) {

}

// @Summary Get bundleIds list
// @Description Get all bundleIds
// @Tags bundleId
// @ID get-bundleIds-list
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Produce  json
// @Success 200 {array} []entities.BundleId
// @Router /bundleIds [get]
func (h *Handler) getBundleIds(c *gin.Context) {
	h.services.QueryHandlerService.GetBundleIds(c)
}

// @Summary Get bundleId
// @Description Get bundleId instance by id
// @Tags bundleId
// @ID get-bundleId
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "bundleId identifier"
// @Produce  json
// @Success 200 {object} entities.BundleId
// @Router /bundleIds/{id} [get]
func (h *Handler) getBundleIdByID(c *gin.Context) {
	h.services.QueryHandlerService.GetBundleIdById(c)
}

// @Summary Delete bundleId
// @Description BundleId deletion by id
// @Tags bundleId
// @ID delete-bundleId
// @Param tenantId path string true "tenantId"
// @Param integrationId path string true "integrationId"
// @Param id path string true "bundleId identifier"
// @Produce  json
// @Success 204 {string} string "Deletion success"
// @Failure 404 {string} string "BundleId not found"
// @Router /bundleId/{id} [delete]
func (h *Handler) deleteBundleIdByID(c *gin.Context) {
	err := h.services.CommandExecutorService.DelBundleIdById(c)
	if err != nil {
		return
	}
}
