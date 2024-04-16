package entity

type ApiCreateBundleId struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Name           string `json:"name"`
			Identifier     string `json:"identifier"`
			SeedID         string `json:"seedId"`
			Platform       string `json:"platform"`
			BundleIDSuffix string `json:"bundleIdSuffix,omitempty"`
		} `json:"attributes"`
	} `json:"data"`
}

type ApiCreateDevice struct {
	Data struct {
		Attributes struct {
			Name     string `json:"name"`
			UDID     string `json:"udid"`
			Platform string `json:"platform"`
		} `json:"attributes"`
		Type string `json:"type"`
	} `json:"data"`
}

type ApiEnableCapability struct {
	Data struct {
		Relationships struct {
			BundleId struct {
				Data struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"bundleId"`
		} `json:"relationships"`
		Attributes struct {
			CapabilityType string `json:"capabilityType"`
			Settings       []struct {
				Key              string `json:"key,omitempty"`
				Name             string `json:"name,omitempty"`
				Description      string `json:"description,omitempty"`
				EnabledByDefault string `json:"enabledByDefault,omitempty"`
				Visible          string `json:"visible,omitempty"`
				AllowedInstances string `json:"allowedInstances,omitempty"`
				MinInstances     string `json:"minInstances,omitempty"`
				Options          []struct {
					Key              string `json:"key,omitempty"`
					Name             string `json:"name,omitempty"`
					Description      string `json:"description,omitempty"`
					EnabledByDefault string `json:"enabledByDefault,omitempty"`
					Enabled          string `json:"enabled,omitempty"`
					SupportsWildcard string `json:"supportsWildcard,omitempty"`
				} `json:"options,omitempty"`
			} `json:"settings,omitempty"`
		} `json:"attributes"`
		Type string `json:"type"`
	} `json:"data"`
}
