package entity

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	BorrowDate time.Time
	ReturnDate time.Time
	Quant      int

	DoctorID *uint
	Doctor   Doctor

	MedicalEquipmentID *uint
	MedicalEquipment   MedicalEquipment `gorm:"foreignKey:MedicalEquipmentID;references:ID"`

	WorkplaceID *uint
	Workplace   Workplace
}
