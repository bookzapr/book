package controller

import (
	"github.com/bookzapr/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /medActivity //เพิ่มข้อมูลใน DB

func CreateMedicalEquipment(c *gin.Context) {

	var medicalEquipment entity.MedicalEquipment

	if err := c.ShouldBindJSON(&medicalEquipment); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&medicalEquipment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": medicalEquipment})

}

// GET /medActivity/:id ดึงข้อมูลเฉพาะตัวที่ต้องการ

func GetMedicalEquipment(c *gin.Context) {

	var medicalEquipment entity.MedicalEquipment

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM medical_equipments WHERE id = ?", id).Scan(&medicalEquipment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": medicalEquipment})

}

// GET /doctor ดึงทั้งหมดใน DB ของ หมอ
func ListMedicalEquipments(c *gin.Context) {

	var medicalEquipment []entity.MedicalEquipment

	if err := entity.DB().Raw("SELECT * FROM medical_equipments").Scan(&medicalEquipment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": medicalEquipment})

}

// DELETE /doctor/:id

func DeleteMedicalEquipment(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM medical_equipments WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "medical_equipments not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /doctor

func UpdateMedicalEquipment(c *gin.Context) {

	var medicalEquipment entity.MedicalEquipment

	if err := c.ShouldBindJSON(&medicalEquipment); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", medicalEquipment.ID).First(&medicalEquipment); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "medical_equipments not found"})

		return

	}

	if err := entity.DB().Save(&medicalEquipment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": medicalEquipment})

}
