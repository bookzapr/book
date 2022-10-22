package entity

import (
	"gorm.io/gorm"
)

type MedicalEquipment struct {
	gorm.Model
	Name string

	TypeofUseID *uint
	TypeofUse   TypeofUse

	// BorrowID *uint
	Borrow []Borrow `gorm:"foreignKey:MedicalEquipmentID"`
}
