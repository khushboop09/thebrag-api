package main

import (
	"thebrag/configs"
	"thebrag/models"
	"thebrag/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = configs.ConnectDB()
)

func main() {
	router := gin.Default()

	configs.LoadEnv()
	defer configs.DisconnectDB(db)
	db.AutoMigrate(&models.Brag{})
	db.AutoMigrate(&models.Category{})

	//routes
	routes.BragRoute(router)
	routes.CategoryRoute(router)

	router.Run("localhost:8080")
}
