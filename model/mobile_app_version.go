package model

// MobileAppVersion struct
type MobileAppVersion struct {
	AppstoreVersion       string             `gorm:"primaryKey" json:"appstore_version"`
	InternalVersion       string             `json:"internal_version"`
	MobileApplicationName string             `gorm:"primaryKey" json:"app_name"`
	MobileApplication     *MobileApplication `gorm:"ForeignKey:MobileApplicationName" json:"app,omitempty"`
}
