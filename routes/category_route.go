package routes

import (
	"thebrag/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	//All routes related to brags comes here
	router.POST("/category", controllers.AddCategory())
	router.GET("/categories", controllers.GetAllCategories())
	router.PUT("/category/", controllers.UpdateCategory())
}
