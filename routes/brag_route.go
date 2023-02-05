package routes

import (
	"thebrag/controllers"

	"github.com/gin-gonic/gin"
)

func BragRoute(router *gin.Engine) {
	//All routes related to brags comes here
	router.POST("/brag", controllers.AddBrag())
	router.GET("/:userId/brags", controllers.GetAllUserBrags())
}
