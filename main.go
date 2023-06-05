package main

import (
	"gin.com/gin/models"
	"gin.com/gin/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Student{})

	r := routes.SetupRoutes(db)
	r.Run()
}
