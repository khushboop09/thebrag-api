package main

import (
	"thebrag/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//routes
	routes.UserRoute(router)
	routes.BragRoute(router)

	router.Run("localhost:8080")
}
