package entity

type Tenant struct {
	ID   uint
	Name string
}

type Integration struct {
	ID       uint
	IssuerId string
	TeamId   string
	Tenant   Tenant
	Devices  []Device `gorm:"many2many:integration_devices"`
}

type BundleId struct {
	ID          uint
	Name        string
	Integration Integration
}

type Capability struct {
	BundleId BundleId
	Type     string
}

type User struct {
	ID   uint
	Name string
}

type Device struct {
	ID           uint
	Name         string
	User         User
	Profiles     []Profile     `gorm:"many2many:profile_devices;"`
	Integrations []Integration `gorm:"many2many:integration_devices"`
}

type Certificate struct {
	ID          uint
	Name        string
	Type        string
	Integration Integration
	Profiles    []Profile `gorm:"many2many:profile_certificates;"`
}

type Profile struct {
	ID           uint
	Name         string
	BundleId     BundleId
	Integration  Integration
	Devices      []Device      `gorm:"many2many:profile_devices;"`
	Certificates []Certificate `gorm:"many2many:profile_certificates;"`
}
