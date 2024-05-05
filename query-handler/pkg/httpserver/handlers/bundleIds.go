package handlers

import (
	"d-kv/signer/db-common/entity"
	"d-kv/signer/query-handler/pkg/httpserver/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func (h *Handler) getBundleIds(c *gin.Context) {
	integrationId := c.Param("integrationId")
	bundleId, err := h.DomainRepos.BundleIdRepo.FindByIntegrationId(c.Request.Context(), integrationId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	capabilities, err := h.DomainRepos.CapabilityRepo.FindByBundleIdId(c.Request.Context(), bundleId.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entities.BundleId{
		Identifier:   bundleId.Identifier,
		Name:         bundleId.Name,
		Capabilities: mapCapabilitiesToTypes(capabilities),
	})
}

func mapCapabilitiesToTypes(capabilities []entity.Capability) []entity.CapabilityType {
	var types []entity.CapabilityType
	for _, capability := range capabilities {
		types = append(types, mapStringCapabilityType(capability.Type))
	}
	return types
}

func mapStringCapabilityType(stringType string) entity.CapabilityType {
	if !slices.Contains(entity.CapabilityTypeValues, stringType) {
		return entity.CapabilityType("Undefined")
	}
	return entity.CapabilityType(stringType)
}

func (h *Handler) getBundleIdByID(c *gin.Context) {
	id := c.Param("id")
	bundleId, err := h.DomainRepos.BundleIdRepo.FindById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	capabilities, err := h.DomainRepos.CapabilityRepo.FindByBundleIdId(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entities.BundleId{
		Identifier:   bundleId.Identifier,
		Name:         bundleId.Name,
		Capabilities: mapCapabilitiesToTypes(capabilities),
	})
}
