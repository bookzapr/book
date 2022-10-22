package main

import (
	"github.com/bookzapr/sa-65-example/controller"

	"github.com/bookzapr/sa-65-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// Doctor Routes

	r.GET("/doctors", controller.ListDoctors)
	r.GET("/doctor/:id", controller.GetDoctor)
	r.POST("/doctors", controller.CreateDoctor)

	//workplace
	r.GET("/workPlaces", controller.ListWorkPlaces)
	r.GET("/workPlace/:id", controller.GetWorkPlace)

	r.GET("/TypeofUses", controller.ListTypeOfUses)
	r.GET("/TypeofUse/:id", controller.GetTypeOfUse)

	//medactivity
	r.GET("/medicalEquipments", controller.ListMedicalEquipments)
	r.GET("/medicalEquipment/:id", controller.GetMedicalEquipment)

	//schedule
	r.GET("/borrows", controller.ListBorrows)
	r.GET("/borrow/:id", controller.GetBorrow)
	r.POST("/borrow", controller.CreateBorrow)

	// Run the server

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
