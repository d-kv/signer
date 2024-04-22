package entity

import "d-kv/signer/command-executor/pkg/entity"

func (input *CreateDevice) Convert() entity.ApiEntity {
	newDeviceData := entity.ApiCreateDevice{
		Data: entity.DeviceData{
			Type: "devices",
			Attributes: entity.DeviceAttributes{
				Name:     input.DeviceName,
				UDID:     input.DeviceUdid,
				Platform: string(input.DevicePlatform),
			},
		},
	}
	return &newDeviceData
}

func (input *CreateBundleId) Convert() entity.ApiEntity {
	newBid := entity.ApiCreateBundleId{
		Data: entity.BundleIdData{
			Type: "bundleIds",
			Attributes: entity.BundleIdAttributes{
				Name:       input.BundleName,
				Identifier: input.BundleIdentifier,
				SeedID:     input.SeedId,
				Platform:   string(input.BundlePlatform),
			},
		},
	}
	return &newBid
}

func (input *EnableCapabilityType) Convert() entity.ApiEntity {
	settingData := entity.SettingData{}

	optionData := entity.OptionData{}

	settingData.Options = append(settingData.Options, optionData)

	attributesData := &entity.AttributesData{
		CapabilityType: string(input.CapabilityType),
		Settings:       []entity.SettingData{settingData},
	}

	bundleIdForCapabilityData := &entity.BundleIdForCapabilityData{
		Id:   input.BundleId,
		Type: "bundleIds",
	}

	bundleIdRelationship := &entity.BundleIdRelationship{
		Data: *bundleIdForCapabilityData,
	}

	relationshipsData := &entity.RelationshipsData{
		BundleId: *bundleIdRelationship,
	}

	dataForCapability := &entity.DataForCapability{
		Relationships: *relationshipsData,
		Attributes:    *attributesData,
		Type:          "bundleIdCapabilities",
	}

	apiEnableCapability := entity.ApiEnableCapability{
		Data: *dataForCapability,
	}

	return &apiEnableCapability
}
