package controllers

import (
	"context"
	"log"
	"net/http"
	"thebrag/configs"
	"thebrag/models"
	"thebrag/responses"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bragsCollection *mongo.Collection = configs.GetCollection(configs.DB, "brags")

func AddBrag() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var brag models.Brag
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&brag); err != nil {
			c.JSON(http.StatusBadRequest, responses.BragResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		newBrag := models.Brag{
			Title:      brag.Title,
			Details:    brag.Details,
			User_Id:    brag.User_Id,
			Created_At: time.Now(),
			Updated_At: time.Now(),
		}

		result, err := bragsCollection.InsertOne(ctx, newBrag)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BragResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.BragResponse{Status: http.StatusCreated, Message: "success", Data: result.InsertedID})
	}
}

func GetAllUserBrags() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		userId := c.Param("userId")
		var brags []models.Brag
		defer cancel()

		opts := options.Find()
		opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
		sortCursor, err := bragsCollection.Find(ctx, bson.D{{Key: "user_id", Value: userId}}, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BragResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}
		// var episodesSorted []bson.M
		if err = sortCursor.All(ctx, &brags); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, responses.BragResponse{Status: http.StatusOK, Message: "success", Data: brags})

	}
}
