package routes

import (
	"thebrag/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to brags comes here
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:id", controllers.GetAUser())
}
