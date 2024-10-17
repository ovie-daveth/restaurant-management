package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`                             // Unique identifier for the note
	OrderID    string             `json:"order_id" validate:"required"`    // Reference to the order the note is attached to
	CustomerID string             `json:"customer_id" validate:"required"` // Reference to the customer leaving the note
	Content    *string            `json:"content" validate:"required"`     // The note content
	CreatedAt  time.Time          `json:"created_at"`                      // Timestamp when the note was created
	UpdatedAt  time.Time          `json:"updated_at"`                      // Timestamp when the note was last updated
}
