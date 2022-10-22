package entity

import (
	// _ "fmt"
	_ "time"

	"gorm.io/gorm"
)

type Workplace struct {
	gorm.Model
	Name    string
	Address string

	// BorrowID *uint
	Borrow []Borrow `gorm:"foreignKey:WorkplaceID"`
}
