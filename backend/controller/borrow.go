package controller

import (
	"github.com/bookzapr/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /schedule //เพิ่มข้อมูลใน DB

func CreateBorrow(c *gin.Context) {

	var borrow entity.Borrow
	var medicalEquipment entity.MedicalEquipment
	var workplace entity.Workplace
	// var doctor entity.Doctor

	if err := c.ShouldBindJSON(&borrow); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// ค้นหา medActivity ด้วย id
	if tx := entity.DB().Where("id = ?", borrow.MedicalEquipmentID).First(&medicalEquipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalEquipment not found"})
		return
	}

	// ค้นหา workPlace ด้วย id
	if tx := entity.DB().Where("id = ?", borrow.WorkplaceID).First(&workplace); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Workplace not found"})
		return
	}

	// ค้นหา doctor ด้วย id
	// if tx := entity.DB().Where("id = ?", schedule.DoctorID).First(&doctor); tx.RowsAffected == 0 {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
	// 		return
	// }

		// 12: สร้าง schedule
		br := entity.Borrow{
			// Doctor:  doctor,             // โยงความสัมพันธ์กับ Entity doctor
			Workplace: 			workplace,                  // โยงความสัมพันธ์กับ workPlace
			MedicalEquipment:   medicalEquipment, 
			Quant:      		borrow.Quant,     // โยงความสัมพันธ์กับ Entity medactivity
			BorrowDate: 		borrow.BorrowDate,
			ReturnDate: 		borrow.ReturnDate,
		}
	

	if err := entity.DB().Create(&br).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": br})

}

// GET /schedule/:id ดึงข้อมูลเฉพาะตัวที่ต้องการ

func GetBorrow(c *gin.Context) {

	var borrow entity.Borrow

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM borrows WHERE id = ?", id).Scan(&borrow).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": borrow})

}

// GET /schedule ดึงทั้งหมดใน DB ของตารางเวลา
func ListBorrows(c *gin.Context) {

	var borrow []entity.Borrow

	if err := entity.DB().Preload("Workplace").Preload("MedicalEquipment.TypeofUse").Raw("SELECT * FROM borrows").Find(&borrow).Error; err != nil {
		//ดึงตารางย่อยมา .preload
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": borrow})

}

// DELETE /borrow/:id

func DeleteBorrow(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM borrows WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "borrow not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /doctor

func UpdateBorrow(c *gin.Context) {

	var borrow entity.Borrow

	if err := c.ShouldBindJSON(&borrow); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", borrow.ID).First(&borrow); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "borrows not found"})

		return

	}

	if err := entity.DB().Save(&borrow).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": borrow})

}
