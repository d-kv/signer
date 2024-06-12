package entity

import "d-kv/signer/db-common/entity"

type CreateDevice struct {
	Outer entity.CreateDevice
}

func (input *CreateDevice) GetId() uint {
	return input.Outer.ID
}
func (input *CreateBundleId) GetId() uint {
	return input.Outer.ID
}
func (input *EnableCapabilityType) GetId() uint {
	return input.Outer.ID
}
func (input *CreateProfile) GetId() uint {
	return input.Outer.ID
}
func (input *CreateCertificate) GetId() uint {
	return input.Outer.ID
}

func (input *CreateDevice) GetIntegrationId() string {
	return input.Outer.IntegrationId
}
func (input *CreateBundleId) GetIntegrationId() string {
	return input.Outer.IntegrationId
}
func (input *EnableCapabilityType) GetIntegrationId() string {
	return input.Outer.IntegrationId
}
func (input *CreateProfile) GetIntegrationId() string {
	return input.Outer.IntegrationId
}
func (input *CreateCertificate) GetIntegrationId() string {
	return input.Outer.IntegrationId
}

func (input *CreateDevice) Convert() ApiEntity {
	newDeviceData := ApiCreateDevice{
		Data: DeviceData{
			Type: "devices",
			Attributes: DeviceAttributes{
				Name:     input.Outer.DeviceName,
				UDID:     input.Outer.DeviceUdid,
				Platform: string(input.Outer.DevicePlatform),
			},
		},
	}
	return &newDeviceData
}

type CreateBundleId struct {
	Outer entity.CreateBundleId
}

func (input *CreateBundleId) Convert() ApiEntity {
	newBid := ApiCreateBundleId{
		Data: BundleIdData{
			Type: "bundleIds",
			Attributes: BundleIdAttributes{
				Name:       input.Outer.BundleName,
				Identifier: input.Outer.BundleIdentifier,
				SeedID:     input.Outer.SeedId,
				Platform:   string(input.Outer.BundlePlatform),
			},
		},
	}
	return &newBid
}

type EnableCapabilityType struct {
	Outer entity.EnableCapabilityType
}

func (input *EnableCapabilityType) Convert() ApiEntity {
	settingData := SettingData{}
	optionData := OptionData{}
	settingData.Options = append(settingData.Options, optionData)
	attributesData := &AttributesData{
		CapabilityType: string(input.Outer.CapabilityType),
		Settings:       []SettingData{settingData},
	}
	bundleIdForCapabilityData := &BundleIdForCapabilityData{
		Id:   input.Outer.BundleId,
		Type: "bundleIds",
	}
	bundleIdRelationship := &BundleIdRelationship{
		Data: *bundleIdForCapabilityData,
	}
	relationshipsData := &RelationshipsData{
		BundleId: *bundleIdRelationship,
	}
	dataForCapability := &DataForCapability{
		Relationships: *relationshipsData,
		Attributes:    *attributesData,
		Type:          "bundleIdCapabilities",
	}
	apiEnableCapability := ApiEnableCapability{
		Data: *dataForCapability,
	}
	return &apiEnableCapability
}

type CreateProfile struct {
	Outer entity.CreateProfile
}

func (input *CreateProfile) Convert() ApiEntity {
	var certificateData []ProfileArgumentData
	for _, id := range input.Outer.CertificateIds {
		certificateData = append(certificateData, ProfileArgumentData{
			ID:   id,
			Type: "certificates",
		})
	}
	var deviceData []ProfileArgumentData
	for _, id := range input.Outer.DeviceIds {
		deviceData = append(deviceData, ProfileArgumentData{
			ID:   id,
			Type: "devices",
		})
	}
	var bundleIdData []ProfileArgumentData
	bundleIdData = append(bundleIdData, ProfileArgumentData{
		ID:   input.Outer.BundleId,
		Type: "bundleIds",
	})
	newProfile := ApiCreateProfile{Data: ProfileData{
		Type: "profiles",
		Attributes: ProfileAttributes{
			Name:        input.Outer.Name,
			ProfileType: string(input.Outer.ProfileType),
		},
		Relationships: ProfileRelationships{
			BundleID:     ProfileArgument{Data: bundleIdData},
			Certificates: ProfileArgument{Data: certificateData},
			Devices:      ProfileArgument{Data: deviceData},
		},
	},
	}
	return &newProfile
}

type CreateCertificate struct {
	Outer entity.CreateCertificate
}

func (input *CreateCertificate) Convert() ApiEntity {
	newCertificate := ApiAddCertificate{Data: CertificateData{
		Attributes: CertificateAttributes{
			CsrContent:      input.Outer.CsrContent,
			CertificateType: string(input.Outer.Type),
		},
		Type: "certificates",
	}}
	return &newCertificate
}
