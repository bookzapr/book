package entity

import (
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`

	// BorrowID *uint
	Borrow []Borrow `gorm:"foreignKey:DoctorID"`
}
