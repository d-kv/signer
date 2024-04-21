package entity

import "d-kv/signer/db-common/entity"

func ConvertCreateDevice(input entity.CreateDevice) ApiEntity {
	newDeviceData := ApiCreateDevice{
		Data: DeviceData{
			Type: "devices",
			Attributes: DeviceAttributes{
				Name:     input.DeviceName,
				UDID:     input.DeviceUdid,
				Platform: string(input.DevicePlatform),
			},
		},
	}
	return &newDeviceData
}

func ConvertCreateBundleId(input entity.CreateBundleId) ApiEntity {
	newBid := ApiCreateBundleId{
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
	return &newBid
}

func ConvertEnableCapabilityType(enableCap entity.EnableCapabilityType) ApiEntity {
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

	apiEnableCapability := ApiEnableCapability{
		Data: dataForCapability,
	}

	return &apiEnableCapability
}
