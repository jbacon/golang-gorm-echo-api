package model

import "gorm.io/datatypes"

// SecurityReportEntry struct
type SecurityReportEntry struct {
	ID                         uint               `gorm:"primaryKey" json:"id"`
	SecurityReportID           uint               `json:"security_report_id"`
	SecurityReport             *SecurityReport    `json:"security_report,omitempty"`
	SecurityCheckSpecShortName string             `json:"short_name"`
	SecurityCheckSpec          *SecurityCheckSpec `gorm:"foreignKey:SecurityCheckSpecShortName" json:"security_check_spec,omitempty"`
	Data                       datatypes.JSON     `json:"data"`
	Vulnerable                 bool               `json:"vulnerable"`
}
