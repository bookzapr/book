package entity

import (
	_ "fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		// &User{},
		&Doctor{},
		&TypeofUse{},
		&Workplace{},
		&MedicalEquipment{},
		&Borrow{},
	)

	db = database

	Phonsak := Doctor{
		Name:  "Phonsak Songsang",
		Email: "Phonsak@gmail.com",
	}
	db.Model(&Doctor{}).Create(&Phonsak)

	Hanoi := Doctor{
		Name:  "Hanoi Slotmachine",
		Email: "Hanoi@gmail.com",
	}
	db.Model(&Doctor{}).Create(&Hanoi)

	Kanothip := Doctor{
		Name:  "Kanothip Lamai",
		Email: "Kanok@hotmail.com",
	}
	db.Model(&Doctor{}).Create(&Kanothip)

	//type

	db.Model(&TypeofUse{}).Create(&TypeofUse{
		Name: "licensed",
	})
	db.Model(&TypeofUse{}).Create(&TypeofUse{
		Name: "inform details",
	})
	db.Model(&TypeofUse{}).Create(&TypeofUse{
		Name: "general",
	})

	var licensed TypeofUse
	var informdetails TypeofUse
	var general TypeofUse
	db.Raw("SELECT * FROM Typeof_uses WHERE name = ?", "licensed").Scan(&licensed)
	db.Raw("SELECT * FROM Typeof_uses WHERE name = ?", "inform details").Scan(&informdetails)
	db.Raw("SELECT * FROM Typeof_uses WHERE name = ?", "general").Scan(&general)

	humanbloodbag := MedicalEquipment{
		Name: "human blood bag",
		TypeofUse: licensed,
	}
	db.Model(&MedicalEquipment{}).Create(&humanbloodbag)

	drugtestingkit := MedicalEquipment{
		Name: "drug testing kit",
		TypeofUse: informdetails,
	}
	db.Model(&MedicalEquipment{}).Create(&drugtestingkit)

	surgicalequipment := MedicalEquipment{
		Name: "surgical equipment",
		TypeofUse: general,
	}
	db.Model(&MedicalEquipment{}).Create(&surgicalequipment)
	//-----------------------------------------------------------------------
	Outpatient := Workplace{
		Name:    "Outpatient",
		Address: "Suranaree 1 floor",
	}
	db.Model(&Workplace{}).Create(&Outpatient)

	Inpatient := Workplace{
		Name:    "Inpatient",
		Address: "Suranaree 2,3 floor",
	}
	db.Model(&Workplace{}).Create(&Inpatient)

	Emergency := Workplace{
		Name:    "Emergency",
		Address: "Suranaree 1 floor",
	}
	db.Model(&Workplace{}).Create(&Emergency)

	Surgery := Workplace{
		Name:    "Surgery",
		Address: "Suranaree 4 floor",
	}
	db.Model(&Workplace{}).Create(&Surgery)
	//-------------------------------------------------------------------------
	Borrow1 := time.Date(2021, 05, 25, 45, 35, 0, 0, time.Now().Location())
	Borrow2 := time.Date(2021, 05, 25, 45, 35, 0, 0, time.Now().Location())
	Borrow3 := time.Date(2021, 05, 25, 45, 35, 0, 0, time.Now().Location())
	Return1 := time.Date(2021, 05, 25, 45, 35, 0, 0, time.Now().Location())
	Return2 := time.Date(2021, 05, 25, 45, 35, 0, 0, time.Now().Location())
	Return3 := time.Date(2021, 05, 25, 45, 35, 0, 0, time.Now().Location())

	var Doctor1 Doctor
	var Doctor2 Doctor
	var Doctor3 Doctor
	db.Raw("SELECT * FROM Doctors WHERE Email = ?", "Phonsak@gmail.com").Scan(&Doctor1)
	db.Raw("SELECT * FROM Doctors WHERE Email = ?", "Hanoi@gmail.com").Scan(&Doctor2)
	db.Raw("SELECT * FROM Doctors WHERE Email = ?", "Kanok@hotmail.com").Scan(&Doctor3)

	db.Model(&Borrow{}).Create(&Borrow{
		BorrowDate:       Borrow1,
		ReturnDate:       Return1,
		Quant:            1,
		Doctor:           Doctor1,
		MedicalEquipment: humanbloodbag,
		Workplace:        Outpatient,
	})

	db.Model(&Borrow{}).Create(&Borrow{
		BorrowDate:       Borrow2,
		ReturnDate:       Return2,
		Quant:            2,
		Doctor:           Doctor2,
		MedicalEquipment: drugtestingkit,
		Workplace:        Inpatient,
	})

	db.Model(&Borrow{}).Create(&Borrow{
		BorrowDate:       Borrow3,
		ReturnDate:       Return3,
		Quant:            3,
		Doctor:           Doctor3,
		MedicalEquipment: surgicalequipment,
		Workplace:        Emergency,
	})
}
