package routes

import (
	"thebrag/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	//All routes related to brags comes here
	router.POST("/:userId/category", controllers.AddCategory())
	router.GET("/:userId/categories", controllers.GetAllCategories())
	router.PUT("/:userId/category", controllers.UpdateCategory())
}
