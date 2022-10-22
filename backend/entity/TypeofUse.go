package entity

import (
	"gorm.io/gorm"
)

type TypeofUse struct {
	gorm.Model
	Name string

	// MedicalEquipmentID *uint
	MedicalEquipment []MedicalEquipment `gorm:"foreignKey:TypeofUseID"`
}
