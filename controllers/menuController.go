package controllers

import (
	"context"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(ctx, bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing the menu items"})
			return
		}
		var allMenus []bson.M

		if err = result.All(ctx, &allMenus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding the menu items"})
			return
		}
		c.JSON(http.StatusOK, allMenus)

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuId := c.Param("Menu_id")

		objID, err := primitive.ObjectIDFromHex(menuId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		var menu models.Menu

		if err = menuCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Menu item not found"})
			return
		}

		c.JSON(http.StatusOK, menu)

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		//create a context
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu

		// Bind the incoming JSON body to the menu model
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the category field
		if !isValidCategory(menu.Category) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
			return
		}

		// Additional validations for other fields
		if menu.Name == "" || menu.PriceRange == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Price are required fields"})
			return
		}

		// Assign a new ObjectID (MongoDB unique identifier)
		menu.ID = primitive.NewObjectID()
		menu.CreatedAt = time.Now()
		menu.UpdatedAt = time.Now()

		// Insert the new menu item into the collection
		if _, err := menuCollection.InsertOne(ctx, menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating the menu item"})
			return
		}

		// Return the inserted menu item along with its ID
		c.JSON(http.StatusOK, gin.H{
			"message":   "menu created successfully",
			"menu_item": menu,
		})

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the menu ID from the URL parameters
		menuId := c.Param("Menu_id")

		// Convert the menu ID from string to ObjectID
		objID, err := primitive.ObjectIDFromHex(menuId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "MenuId format is invalid"})
			return
		}

		// Create a variable to hold the updated menu data
		var updatedMenu models.Menu

		// Bind the incoming JSON body to the updatedMenu model
		if err := c.BindJSON(&updatedMenu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Prepare the update document
		update := bson.M{
			"$set": bson.M{
				"name":         updatedMenu.Name,
				"description":  updatedMenu.Description,
				"category":     updatedMenu.Category,
				"availability": updatedMenu.Availability,
				"end_time":     updatedMenu.EndTime,
				"start_time":   updatedMenu.StartTime,
				"tags":         updatedMenu.Tags,
				"price_range":  updatedMenu.PriceRange,
				"image_url":    updatedMenu.ImageURL,
				"updated_at":   time.Now(),
			},
		}

		// Update the menu in the collection
		result := menuCollection.FindOneAndUpdate(
			context.TODO(),
			bson.M{"_id": objID},
			update,
			options.FindOneAndUpdate().SetReturnDocument(options.After),
		)

		// Check for errors during the update
		if err := result.Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the menu"})
			}
			return
		}

		// Successfully updated the menu, return the updated menu item
		var updatedMenuResponse models.Menu
		if err := result.Decode(&updatedMenuResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while retrieving the updated menu"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "Menu updated successfully",
			"updated_menu": updatedMenuResponse,
		})
	}
}

func AddItemsToMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the menu ID from the URL parameters
		menuId := c.Param("Menu_id")

		// Convert the menu ID from string to ObjectID
		objID, err := primitive.ObjectIDFromHex(menuId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "MenuId format is invalid"})
			return
		}

		// Create a variable to hold the updated menu data
		var updatedMenu models.Menu

		// Bind the incoming JSON body to the updatedMenu model
		if err := c.BindJSON(&updatedMenu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Prepare the update document
		update := bson.M{
			"$set": bson.M{
				"items":      updatedMenu.Items,
				"updated_at": time.Now(),
			},
		}

		// Update the menu in the collection
		result := menuCollection.FindOneAndUpdate(
			context.TODO(),
			bson.M{"_id": objID},
			update,
			options.FindOneAndUpdate().SetReturnDocument(options.After),
		)

		// Check for errors during the update
		if err := result.Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the menu"})
			}
			return
		}

		// Successfully updated the menu, return the updated menu item
		var updatedMenuResponse models.Menu
		if err := result.Decode(&updatedMenuResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while retrieving the updated menu"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "Menu updated successfully",
			"updated_menu": updatedMenuResponse,
		})
	}
}

// Function to check if the category is valid
func isValidCategory(category models.Category) bool {
	switch category {
	case models.MainCourse, models.Asides, models.Desert:
		return true
	default:
		return false
	}
}
