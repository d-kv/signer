package entity

import "d-kv/signer/db-common/entity"

func ConvertBundleId(integration *entity.Integration, id *entity.CreateBundleId) *entity.BundleId {
	newBid := &entity.BundleId{
		ID:          id.BundleIdentifier,
		Name:        id.BundleName,
		Integration: *integration,
	}
	return newBid
}

func ConvertCapability(id *entity.BundleId, capability *entity.EnableCapabilityType) *entity.Capability {
	newCap := &entity.Capability{
		BundleId: *id,
		Type:     string(capability.CapabilityType),
	}
	return newCap
}

func ConvertDevice(device *entity.CreateDevice, user *entity.User, profile []entity.Profile, integration []entity.Integration) *entity.Device {
	newDevice := &entity.Device{
		UDID:         device.DeviceUdid,
		Name:         device.DeviceName,
		User:         *user,
		Profiles:     profile,
		Integrations: integration,
	}
	return newDevice
}
