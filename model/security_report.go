package model

import (
	"time"
)

// SecurityReports struct
type SecurityReports []SecurityReport

// SecurityReport struct
type SecurityReport struct {
	ID               uint                  `gorm:"primaryKey" json:"id"`
	App              string                `json:"app_name"`
	Version          string                `json:"app_version"`
	MobileAppVersion *MobileAppVersion     `gorm:"ForeignKey:Version,App" json:"app_version_details,omitempty"`
	CreatedDate      time.Time             `json:"created_date"`
	Tests            []SecurityReportEntry `json:"tests"`
}
