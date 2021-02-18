package model

// MobileApplications struct
type MobileApplications []MobileApplication

// MobileApplication struct
type MobileApplication struct {
	Name              string             `gorm:"primaryKey" json:"name"`
	Platform          string             `json:"platform"`
	BundleID          string             `json:"bundle_id"`
	MobileAppVersions []MobileAppVersion `gorm:"foreignKey:MobileApplicationName" json:"versions"`
}
