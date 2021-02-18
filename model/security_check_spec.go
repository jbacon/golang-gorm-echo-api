package model

import "gorm.io/datatypes"

// SecurityCheckSpecs struct
type SecurityCheckSpecs []SecurityCheckSpec

// SecurityCheckSpec struct
type SecurityCheckSpec struct {
	ShortName   string         `gorm:"primaryKey" json:"short_name"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Severity    string         `json:"severity"`
	Score       float32        `json:"score"`
	Regulatory  datatypes.JSON `json:"regulatory"`
	Fields      datatypes.JSON `json:"fields"`
}
