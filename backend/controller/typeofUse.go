package controller

import (
	"github.com/bookzapr/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /doctor //เพิ่มข้อมูลใน DB

func CreateTypeofUse(c *gin.Context) {

	var typeofUse entity.TypeofUse

	if err := c.ShouldBindJSON(&typeofUse); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&typeofUse).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": typeofUse})

}

// GET /doctor/:id ดึงข้อมูลเฉพาะตัวที่ต้องการ

func GetTypeOfUse(c *gin.Context) {

	var typeofUse entity.TypeofUse

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Typeof_uses WHERE id = ?", id).Scan(&typeofUse).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": typeofUse})

}

// GET /doctor ดึงทั้งหมดใน DB ของ หมอ
func ListTypeOfUses(c *gin.Context) {

	var typeOfUse []entity.TypeofUse

	if err := entity.DB().Raw("SELECT * FROM Typeof_uses").Scan(&typeOfUse).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": typeOfUse})

}

// DELETE /doctor/:id

func DeleteTypeOfUse(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Typeof_uses WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Typeof_uses not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /doctor

func UpdateTypeOfUse(c *gin.Context) {

	var typeOfUse entity.TypeofUse

	if err := c.ShouldBindJSON(&typeOfUse); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", typeOfUse.ID).First(&typeOfUse); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&typeOfUse).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": typeOfUse})

}
