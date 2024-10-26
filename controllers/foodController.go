package controllers

import (
	"context"
	"fmt"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := foodCollection.Find(ctx, bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing the FOOD items"})
			return
		}
		var allMenus []bson.M

		if err = result.All(ctx, &allMenus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding the FOOD items"})
			return
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract food ID from the URL parameters
		foodID := c.Param("food_id")

		// Convert the food ID string to a MongoDB ObjectID
		objID, err := primitive.ObjectIDFromHex(foodID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		// Create a context with a timeout for the MongoDB operation
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Define a variable to hold the result
		var food models.Food

		// Query the database for a document with the given ObjectID
		err = foodCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&food)
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

		fmt.Printf("Received food: %+v\n", food)

		// Validate the menuId provided in the request
		menuID := food.Menu_id //menu_id is a string

		if menuID == nil || *menuID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "menuId is required"})
			return
		}

		objID, err := primitive.ObjectIDFromHex(*menuID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		// Check if the provided menuId is valid
		var menu models.Menu
		err = menuCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menuId"})
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
		_, err = foodCollection.InsertOne(ctx, food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating the food item"})
			return
		}

		// Return the inserted food item along with its ID
		c.JSON(http.StatusOK, gin.H{
			"message":   "Food created successfully",
			"food_item": food,
		})
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

		foodId := c.Param("food_id")

		objID, err := primitive.ObjectIDFromHex(foodId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Id you entered is incorrect"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var updatedFood models.Food

		// Bind the incoming JSON body to the food model

		if err = c.BindJSON(&updatedFood); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		}

		defer cancel()

		update := bson.M{
			"$set": bson.M{
				"name":       updatedFood.Name,
				"updated_at": time.Now(),
				"food_image": updatedFood.Food_Image,
				"price":      updatedFood.Price,
			},
		}

		// Update the menu in the collection
		result := foodCollection.FindOneAndUpdate(
			ctx,
			bson.M{"_id": objID},
			update,
			options.FindOneAndUpdate().SetReturnDocument(options.After),
		)

		// Check for errors during the update
		if err := result.Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the FoodItem"})
			}
			return
		}

		// Successfully updated the menu, return the updated menu item
		var updatedFoodResponse models.Food
		if err := result.Decode(&updatedFoodResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while retrieving the updated menu"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "Food updated successfully",
			"updated_food": updatedFoodResponse,
		})
	}

}
func round(num float64) int {
	return 0
}

func toFixed(num float64, precision int) float64 {
	return 3.3
}
