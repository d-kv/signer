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
		return nil, errors.InvalidArgumentError(input.BundlePlatform)
	}
	newCommand := &entity.CreateBundleId{
		TenantId:         tenantId,
		IntegrationId:    integrationId,
		BundleIdentifier: input.BundleIdentifier,
		BundleName:       input.BundleName,
		BundlePlatform:   platform,
		SeedId:           input.SeedId,
		Status:           entity.Created,
	}
	return newCommand, nil
}

func ConvertCapability(input *InputEnableCapability, tenantId string, integrationId string) (*entity.EnableCapabilityType, error) {
	var capabilityType entity.CapabilityType
	if slices.Contains(entity.CapabilityTypeValues, input.CapabilityType) {
		capabilityType = entity.CapabilityType(input.CapabilityType)
	} else {
		return nil, errors.InvalidArgumentError(input.CapabilityType)
	}
	newCommand := &entity.EnableCapabilityType{
		TenantId:       tenantId,
		IntegrationId:  integrationId,
		BundleId:       input.BundleId,
		CapabilityType: capabilityType,
		Status:         entity.Created,
	}
	return newCommand, nil
}

func ConvertDevice(input *InputCreateDevice, tenantId string, integrationId string) (*entity.CreateDevice, error) {
	var platform entity.Platform
	if slices.Contains(entity.PlatformValues, input.DevicePlatform) {
		platform = entity.Platform(input.DevicePlatform)
	} else {
		return nil, errors.InvalidArgumentError(input.DevicePlatform)
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

func ConvertProfile(input *InputCreateProfile, tenantId string, integrationId string) (*entity.CreateProfile, error) {
	var profileType entity.ProfileType
	if slices.Contains(entity.ProfileTypeValues, input.ProfileType) {
		profileType = entity.ProfileType(input.ProfileType)
	} else {
		return nil, errors.InvalidArgumentError(input.ProfileType)
	}
	newCommand := &entity.CreateProfile{
		TenantId:       tenantId,
		IntegrationId:  integrationId,
		Name:           input.Name,
		ProfileType:    profileType,
		BundleId:       input.BundleId,
		CertificateIds: input.CertificateIds,
		DeviceIds:      input.DeviceIds,
		Status:         entity.Created,
	}
	return newCommand, nil
}

func ConvertCertificate(input *InputCreateCertificate, tenantId string, integrationId string) (*entity.CreateCertificate, error) {
	var certType entity.CertificateType
	if slices.Contains(entity.CertificateTypeValues, input.CertificateType) {
		certType = entity.CertificateType(input.CertificateType)
	} else {
		return nil, errors.InvalidArgumentError(input.CertificateType)
	}
	newCommand := &entity.CreateCertificate{
		TenantId:      tenantId,
		IntegrationId: integrationId,
		CsrContent:    input.CsrContent,
		Type:          certType,
		Status:        entity.Created,
	}
	return newCommand, nil
}
