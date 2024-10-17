package controllers

import (
	"context"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

type Response struct {
	Message string `json:"message"`
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := foodCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing the menu items"})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract food ID from the URL parameters
		foodID := c.Param("id")

		// // Convert the food ID string to a MongoDB ObjectID
		// objID, err := primitive.ObjectIDFromHex(foodID)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		// 	return
		// }

		// Create a context with a timeout for the MongoDB operation
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Define a variable to hold the result
		var food models.Food
		var err error

		// Query the database for a document with the given ObjectID
		err = foodCollection.FindOne(ctx, bson.M{"_id": foodID}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Food item not found"})
			return
		}

		// Return the food item as JSON
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food

		// Bind the incoming JSON body to the food model
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the fields (you can add more validation as needed)
		if food.Name == nil || food.Price == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Price are required fields"})
			return
		}

		// Assign a new ObjectID (MongoDB unique identifier)
		food.ID = primitive.NewObjectID()
		food.Created_at = time.Now()
		food.Updated_at = time.Now()

		// Insert the new food item into the collection
		result, err := foodCollection.InsertOne(ctx, food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating the food item"})
			return
		}

		// Return the inserted food item along with its ID
		c.JSON(http.StatusOK, gin.H{
			"message":   "Food created successfully",
			"food_id":   result.InsertedID,
			"food_item": food,
		})
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func round(num float64) int {
	return 0
}

func toFixed(num float64, precision int) float64 {
	return 3.3
}
