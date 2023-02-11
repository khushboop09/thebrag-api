package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"thebrag/configs"
	"thebrag/models"
	"thebrag/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = configs.ConnectDB()

func AddBrag() gin.HandlerFunc {
	return func(c *gin.Context) {
		var brag models.Brag

		//validate the request body
		if err := c.ShouldBindJSON(&brag); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}
		if brag.Title == "" {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: "title cannot be empty"})
			return
		}
		newBrag := models.Brag{
			Title:      brag.Title,
			Details:    brag.Details,
			CategoryID: brag.CategoryID,
		}

		result := db.Create(&newBrag)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: "something went wrong"})
			return
		}
		c.JSON(http.StatusCreated, responses.APIResponse{Status: http.StatusCreated, Message: "success", Data: newBrag.ID})
	}
}

func GetAllBrags() gin.HandlerFunc {
	return func(c *gin.Context) {
		var brags []models.Brag
		skip := 0
		limit := 10
		var err error
		if c.Query("skip") != "" {
			skip, err = strconv.Atoi(c.Query("skip"))
		}
		if c.Query("limit") != "" {
			limit, err = strconv.Atoi(c.Query("limit"))
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}
		db.Limit(limit).Offset(skip).Preload("Category").Find(&brags)

		var response []responses.BragResponse
		for _, brag := range brags {
			item := responses.BragResponse{
				ID:           brag.ID,
				CategoryName: brag.Category.Name,
				CategoryID:   brag.CategoryID,
				Title:        brag.Title,
				Details:      brag.Details,
			}
			response = append(response, item)
		}
		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: response})
	}
}

func GetABrag() gin.HandlerFunc {
	return func(c *gin.Context) {
		bragId := c.Param("bragId")
		var brag models.Brag

		result := db.Preload("Category").First(&brag, bragId)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: "brag not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: result.Error})
			return
		}

		bragResponse := responses.BragResponse{
			ID:           brag.ID,
			CategoryName: brag.Category.Name,
			CategoryID:   brag.CategoryID,
			Title:        brag.Title,
			Details:      brag.Details,
		}
		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: bragResponse})
	}
}

func DeleteBrag() gin.HandlerFunc {
	return func(c *gin.Context) {
		bragId := c.Param("bragId")

		result := db.Delete(&models.Brag{}, bragId)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: result.Error})
			return
		}
		if result.RowsAffected == 1 {
			c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: "deleted"})
		} else {
			c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: "brag not found"})
		}

	}
}

func UpdateBrag() gin.HandlerFunc {
	return func(c *gin.Context) {
		var brag models.Brag

		//validate the request body
		if err := c.ShouldBindJSON(&brag); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}
		if brag.Title == "" || brag.ID == 0 {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: "ID or title cannot be empty"})
			return
		}
		result := db.Model(&brag).Updates(models.Brag{Title: brag.Title, Details: brag.Details, CategoryID: brag.CategoryID})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: result.Error})
			return
		}

		if result.RowsAffected == 1 {
			c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: brag.ID})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: "brag not updated, please try again!"})
	}
}
