package controllers

import (
	"errors"
	"net/http"
	"thebrag/models"
	"thebrag/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		//validate the request body
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}
		if user.Name == "" {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: "name cannot be empty"})
			return
		}
		newUser := models.User{
			Name:  user.Name,
			Email: user.Email,
		}

		result := db.Create(&newUser)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: "something went wrong"})
			return
		}
		c.JSON(http.StatusCreated, responses.APIResponse{Status: http.StatusCreated, Message: "success", Data: newUser.ID})
	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Params.ByName("id")
		var user models.User

		result := db.First(&user, userId)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: "user not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, responses.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: result.Error})
			return
		}

		userResponse := responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		c.JSON(http.StatusOK, responses.APIResponse{Status: http.StatusOK, Message: "success", Data: userResponse})
	}
}
