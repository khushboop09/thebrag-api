package routes

import (
	"thebrag/controllers"

	"github.com/gin-gonic/gin"
)

func BragRoute(router *gin.Engine) {
	//All routes related to brags comes here
	router.POST("/:userId/brag", controllers.AddBrag())
	router.GET("/:userId/brags", controllers.GetAllBrags())
	router.GET("/:userId/brag/:bragId", controllers.GetABrag())
	router.DELETE("/:userId/brag/:bragId", controllers.DeleteBrag())
	router.PUT("/:userId/brag", controllers.UpdateBrag())
}
