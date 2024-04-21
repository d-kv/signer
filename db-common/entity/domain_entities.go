package entity

type Tenant struct {
	ID   string
	Name string
}

type Integration struct {
	ID       string
	IssuerId string
	TeamId   string
	Tenant   Tenant
	Devices  []Device `gorm:"many2many:integration_devices"`
}

type BundleId struct {
	ID          string
	Name        string
	Integration Integration
}

type Capability struct {
	BundleId BundleId
	Type     string
}

type User struct {
	ID   string
	Name string
}

type Device struct {
	ID           string
	Name         string
	User         User
	Profiles     []Profile     `gorm:"many2many:profile_devices;"`
	Integrations []Integration `gorm:"many2many:integration_devices"`
}

type Certificate struct {
	ID          string
	Name        string
	Type        string
	Integration Integration
	Profiles    []Profile `gorm:"many2many:profile_certificates;"`
}

type Profile struct {
	ID           string
	Name         string
	BundleId     BundleId
	Integration  Integration
	Devices      []Device      `gorm:"many2many:profile_devices;"`
	Certificates []Certificate `gorm:"many2many:profile_certificates;"`
}
