package controllers

import (
	"net/http"
	"strconv"
	"thebrag/models"
	"thebrag/responses"

	"github.com/gin-gonic/gin"
)

func AddCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		userId, err := strconv.Atoi(c.Params.ByName("userId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: "Invalid user"})
			return
		}

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
			Name:   category.Name,
			UserId: userId,
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
		userId := c.Params.ByName("userId")
		var categories []models.Category
		var err error

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}
		db.Where("user_id = ?", userId).Find(&categories)
		var response []responses.CategoryResponse
		for _, category := range categories {
			item := responses.CategoryResponse{
				ID:   category.ID,
				Name: category.Name,
			}
			response = append(response, item)
		}
		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: response})
	}
}

func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		userId := c.Params.ByName("userId")

		//validate the request body
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}
		if category.Name == "" {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: "name cannot be empty"})
			return
		}
		result := db.Model(&category).Where("user_id = ?", userId).Updates(models.Category{Name: category.Name})
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
