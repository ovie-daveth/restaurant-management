package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID              primitive.ObjectID `bson:"_id"`              // Unique ID for the table
	TableNumber     *int               `json:"table_number"`     // Table number in the restaurant
	SeatingCapacity *int               `json:"seating_capacity"` // Number of people it can seat
	Status          *string            `json:"status"`           // Status: available, reserved, occupied
	Location        *string            `json:"location"`         // Where the table is located (indoor/outdoor)
	CreatedAt       time.Time          `json:"created_at"`       // Timestamp when the table was created
	UpdatedAt       time.Time          `json:"updated_at"`       // Timestamp when the table was last updated
	TableID         string             `json:"table_id"`         // Unique table ID (human-readable)
}
