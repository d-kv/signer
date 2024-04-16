package entity

func convertCreateDevice(input *CreateDevice) *ApiCreateDevice {
	newDeviceData := &ApiCreateDevice{
		Data: DeviceData{
			Type: "",
			Attributes: DeviceAttributes{
				Name:     input.DeviceName,
				UDID:     input.DeviceUdid,
				Platform: string(input.DevicePlatform),
			},
		},
	}
	return newDeviceData
}

func convertCreateBundleId(input *CreateBundleId) *ApiCreateBundleId {
	newBid := &ApiCreateBundleId{
		Data: BundleIdData{
			Type: "bundleIds",
			Attributes: BundleIdAttributes{
				Name:       input.BundleName,
				Identifier: input.BundleIdentifier,
				SeedID:     input.SeedId,
				Platform:   string(input.BundlePlatform),
			},
		},
	}
	return newBid
}

func convertEnableCapability(enableCap *EnableCapabilityType) *ApiEnableCapability {
	settingData := SettingData{}

	optionData := OptionData{}

	settingData.Options = append(settingData.Options, optionData)

	attributesData := AttributesData{
		CapabilityType: string(enableCap.CapabilityType),
		Settings:       []SettingData{settingData},
	}

	bundleIdForCapabilityData := BundleIdForCapabilityData{
		Id:   enableCap.BundleId,
		Type: "bundleIds",
	}

	bundleIdRelationship := BundleIdRelationship{
		Data: bundleIdForCapabilityData,
	}

	relationshipsData := RelationshipsData{
		BundleId: bundleIdRelationship,
	}

	dataForCapability := DataForCapability{
		Relationships: relationshipsData,
		Attributes:    attributesData,
		Type:          "bundleIdCapabilities",
	}

	apiEnableCapability := &ApiEnableCapability{
		Data: dataForCapability,
	}

	return apiEnableCapability
}
