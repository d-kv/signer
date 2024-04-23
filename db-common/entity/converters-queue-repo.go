package entity

func ConvertBundleId(integration *Integration, id *CreateBundleId) *BundleId {
	newBid := &BundleId{
		ID:          id.BundleIdentifier,
		Name:        id.BundleName,
		Integration: *integration,
	}
	return newBid
}

func ConvertCapability(id *BundleId, capability *EnableCapabilityType) *Capability {
	newCap := &Capability{
		BundleId: *id,
		Type:     string(capability.CapabilityType),
	}
	return newCap
}

func ConvertDevice(device *CreateDevice, user *User, profile []Profile, integration []Integration) *Device {
	newDevice := &Device{
		ID:           device.DeviceUdid,
		Name:         device.DeviceName,
		User:         *user,
		Profiles:     profile,
		Integrations: integration,
	}
	return newDevice
}
