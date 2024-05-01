package entities

import (
	"d-kv/signer/db-common/entity"
	"golang.org/x/crypto/openpgp/errors"
	"slices"
)

func ConvertBundleInput(input *InputCreateBundleId, tenantId string, integrationId string) (*entity.CreateBundleId, error) {
	var platform entity.Platform
	if slices.Contains(entity.PlatformValues, input.BundlePlatform) {
		platform = entity.Platform(input.BundlePlatform)
	} else {
		return nil, errors.InvalidArgumentError("")
	}
	newCommand := &entity.CreateBundleId{
		TenantId:         tenantId,
		IntegrationId:    integrationId,
		BundleIdentifier: input.BundleIdentifier,
		BundleName:       input.BundleName,
		BundlePlatform:   platform,
		SeedId:           input.SeedId,
		Status:           entity.Completed,
	}
	return newCommand, nil
}
func ConvertCapability(input *InputEnableCapability, tenantId string, integrationId string) (*entity.EnableCapabilityType, error) {
	var c entity.CapabilityType
	if slices.Contains(entity.CapabilityTypeValues, input.CapabilityType) {
		c = entity.CapabilityType(input.CapabilityType)
	} else {
		return nil, errors.InvalidArgumentError("")
	}
	newCommand := &entity.EnableCapabilityType{
		TenantId:       tenantId,
		IntegrationId:  integrationId,
		BundleId:       input.BundleId,
		CapabilityType: c,
		Status:         entity.Created,
	}
	return newCommand, nil
}
func ConvertDevice(input *InputCreateDevice, tenantId string, integrationId string) (*entity.CreateDevice, error) {
	var platform entity.Platform
	if slices.Contains(entity.PlatformValues, input.DevicePlatform) {
		platform = entity.Platform(input.DevicePlatform)
	} else {
		return nil, errors.InvalidArgumentError("")
	}
	newCommand := &entity.CreateDevice{
		TenantId:       tenantId,
		IntegrationId:  integrationId,
		DeviceName:     input.DeviceName,
		DevicePlatform: platform,
		DeviceUdid:     input.DeviceUdid,
		Status:         entity.Created,
	}
	return newCommand, nil
}
