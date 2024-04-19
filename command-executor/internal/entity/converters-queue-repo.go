package entity

func ConvertBundleId(integration *Integration, id *CreateBundleId) *BundleId {
	newBid := &BundleId{
		ID:          id.ID,
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

func ConvertDevice(device *Device, user *User, profile *Profile, integration *Integration) *Device {
	newDevice := &Device{
		ID:           device.ID,
		Name:         device.Name,
		User:         *user,
		Profiles:     []Profile{*profile},
		Integrations: []Integration{*integration},
	}
	return newDevice
}
