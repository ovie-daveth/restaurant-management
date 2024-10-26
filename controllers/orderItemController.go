package controllers

import (
	"context"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetorderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := orderItemCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var allOrderItem []bson.M

		if err = result.All(ctx, &allOrderItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, allOrderItem)
	}
}

func GetorderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemId := c.Param("orderItem_id")

		objID, err := primitive.ObjectIDFromHex(orderItemId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "orderItemId is wrong"})
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Define a variable to hold the result
		var order models.OrderItem

		// Query the database for a document with the given ObjectID
		err = orderItemCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Food item not found"})
			return
		}

		// Return the food item as JSON
		c.JSON(http.StatusOK, order)
	}
}

func GetorderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
