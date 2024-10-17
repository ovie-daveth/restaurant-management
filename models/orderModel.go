package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`                                  // Unique identifier for the order
	CustomerID string             `json:"customer_id" validate:"required"`      // Reference to the customer who placed the order
	OrderItems []OrderItem        `json:"order_items" validate:"required,dive"` // List of items in the order
	TotalPrice float64            `json:"total_price" validate:"required"`      // Total price of the order
	Status     string             `json:"status" validate:"required"`           // Status of the order (e.g., "pending", "completed")
	CreatedAt  time.Time          `json:"created_at"`                           // Timestamp for order creation
	UpdatedAt  time.Time          `json:"updated_at"`                           // Timestamp for the last order update
}
