package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	FoodID   primitive.ObjectID `bson:"food_id" json:"food_id" validate:"required"` // Reference to the food item
	Quantity int                `json:"quantity" validate:"required,min=1"`         // Quantity of the food item
	Price    float64            `json:"price" validate:"required"`                  // Price of the food item
}
