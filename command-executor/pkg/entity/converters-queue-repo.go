package entity

import dbEntity "d-kv/signer/db-common/entity"

func ConvertBundleId(integration *dbEntity.Integration, id *dbEntity.CreateBundleId, response *BundleIdResponse) *dbEntity.BundleId {
	newBid := &dbEntity.BundleId{
		ID: response.Id,
		/*Identifier: id.BundleIdentifier,*/
		Name:        id.BundleName,
		Integration: *integration,
	}
	return newBid
}

func ConvertCapability(id *dbEntity.BundleId, capability *dbEntity.EnableCapabilityType) *dbEntity.Capability {
	newCap := &dbEntity.Capability{
		BundleId: *id,
		Type:     string(capability.CapabilityType),
	}
	return newCap
}

func ConvertDevice(device *dbEntity.CreateDevice, user *dbEntity.User, profile []dbEntity.Profile, integration []dbEntity.Integration) *dbEntity.Device {
	newDevice := &dbEntity.Device{
		UDID:         device.DeviceUdid,
		Name:         device.DeviceName,
		User:         *user,
		Profiles:     profile,
		Integrations: integration,
	}
	return newDevice
}

func ConvertProfile(profile *dbEntity.CreateProfile, devices []dbEntity.Device, certificates []dbEntity.Certificate,
	response *ProfileResponse, bundleId *dbEntity.BundleId, integration *dbEntity.Integration) *dbEntity.Profile {
	newProfile := &dbEntity.Profile{
		ID:             response.Data.ID,
		Name:           profile.Name,
		ProfileContent: response.Data.Attributes.ProfileContent,
		Devices:        devices,
		Certificates:   certificates,
		BundleIdId:     bundleId.ID,
		IntegrationId:  integration.ID,
	}
	return newProfile
}

func ConvertCertificate(certificate *dbEntity.CreateCertificate, profiles []dbEntity.Profile,
	integration *dbEntity.Integration, response *CertificateResponse) *dbEntity.Certificate {
	newCertificate := &dbEntity.Certificate{
		ID:                 response.Data.ID,
		Name:               response.Data.Attributes.Name,
		Type:               string(certificate.Type),
		CertificateContent: response.Data.Attributes.CertificateContent,
		Profiles:           profiles,
		IntegrationId:      integration.ID,
	}
	return newCertificate
}
