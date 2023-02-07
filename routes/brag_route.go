package routes

import (
	"thebrag/controllers"

	"github.com/gin-gonic/gin"
)

func BragRoute(router *gin.Engine) {
	//All routes related to brags comes here
	router.POST("/brag", controllers.AddBrag())
	router.GET("/brags", controllers.GetAllBrags())
	router.GET("/brag/:bragId", controllers.GetABrag())
	router.DELETE("/brag/:bragId", controllers.DeleteBrag())
	router.POST("/brag/:bragId", controllers.UpdateBrag())
}
