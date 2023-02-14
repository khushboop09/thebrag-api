package main

import (
	"fmt"
	"os"
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
	db.AutoMigrate(&models.User{})

	//routes
	routes.BragRoute(router)
	routes.CategoryRoute(router)
	routes.UserRoute(router)

	listen_url := fmt.Sprintf("%s:8080", os.Getenv("SERVER_IP"))
	router.Run(listen_url)
}
