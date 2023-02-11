package controllers

import (
	"net/http"
	"thebrag/models"
	"thebrag/responses"

	"github.com/gin-gonic/gin"
)

func AddCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category

		//validate the request body
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}
		if category.Name == "" {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: "title cannot be empty"})
			return
		}
		newCategory := models.Category{
			Name: category.Name,
		}

		result := db.Create(&newCategory)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: "something went wrong"})
			return
		}
		c.JSON(http.StatusCreated, responses.APIResponse{Status: http.StatusCreated, Message: "success", Data: newCategory.ID})
	}
}

func GetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.Category
		var err error

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}
		db.Find(&categories)
		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: categories})
	}
}

func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category

		//validate the request body
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}
		if category.Name == "" {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: "name cannot be empty"})
			return
		}
		result := db.Model(&category).Updates(models.Category{Name: category.Name})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: result.Error})
			return
		}

		if result.RowsAffected == 1 {
			c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: category.ID})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: "category not updated, please try again!"})
	}
}
