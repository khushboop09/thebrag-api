package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"thebrag/configs"
	"thebrag/models"
	"thebrag/responses"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			ID:      primitive.NewObjectID(),
			Title:   brag.Title,
			Details: brag.Details,
			// User_Id:    brag.User_Id,
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

func GetAllBrags() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		var brags []models.Brag
		skip, err := strconv.ParseInt(c.Query("skip"), 10, 64)
		limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
		defer cancel()

		opts := options.Find()
		opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
		opts.SetLimit(limit)
		opts.SetSkip(skip)

		sortCursor, err := bragsCollection.Find(ctx, bson.D{{}}, opts)
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

func GetABrag() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		bragId := c.Param("bragId")
		var brag models.Brag
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bragId)
		err := bragsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&brag)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BragResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.BragResponse{Status: http.StatusOK, Message: "success", Data: brag})

	}
}

func DeleteBrag() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bragId := c.Param("bragId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bragId)

		result, err := bragsCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BragResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.BragResponse{Status: http.StatusNotFound, Message: "error", Data: "Brag with specified ID not found!"},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.BragResponse{Status: http.StatusOK, Message: "success", Data: "Brag successfully deleted!"},
		)
	}
}

func UpdateBrag() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		bragId := c.Param("bragId")
		var brag models.Brag
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(bragId)

		//validate the request body
		if err := c.BindJSON(&brag); err != nil {
			c.JSON(http.StatusBadRequest, responses.BragResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		update := bson.M{"title": brag.Title, "details": brag.Details, "updated_at": time.Now()}
		result, err := bragsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BragResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		if result.MatchedCount == 1 {
			c.JSON(http.StatusOK, responses.BragResponse{Status: http.StatusOK, Message: "success", Data: bragId})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.BragResponse{Status: http.StatusInternalServerError, Message: "error", Data: "Brag not updated, please try again!"})
	}
}
