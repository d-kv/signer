package entity

type Tenant struct {
	ID   string
	Name string
}

type Integration struct {
	ID       string `gorm:"primary_key"`
	IssuerId string
	TeamId   string
	Tenant   Tenant   `gorm:"references:ID"`
	Devices  []Device `gorm:"many2many:integration_devices"`
	// Ids
	TenantId string
}

type BundleId struct {
	ID          string `gorm:"primary_key"`
	Identifier  string
	Name        string
	Integration Integration `gorm:"references:ID"`
	// Ids
	IntegrationId string
}

type Capability struct {
	BundleId BundleId `gorm:"references:ID"`
	Type     string
	// Ids
	BundleIdId string
}

type User struct {
	ID   string `gorm:"primary_key"`
	Name string
}

type Device struct {
	UDID         string `gorm:"primary_key"`
	Name         string
	Platform     string
	User         User          `gorm:"references:ID"`
	Profiles     []Profile     `gorm:"many2many:profile_devices;"`
	Integrations []Integration `gorm:"many2many:integration_devices"`
	// Ids
	UserID string
}

type Certificate struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Type        string
	Integration Integration `gorm:"references:ID"`
	Profiles    []Profile   `gorm:"many2many:profile_certificates;"`
	// Ids
	IntegrationId string
}

type Profile struct {
	ID           string `gorm:"primary_key"`
	Name         string
	BundleId     BundleId      `gorm:"references:ID"`
	Integration  Integration   `gorm:"references:ID"`
	Devices      []Device      `gorm:"many2many:profile_devices;"`
	Certificates []Certificate `gorm:"many2many:profile_certificates;"`
	// Ids
	BundleIdId    string
	IntegrationId string
}
