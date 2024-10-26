package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`                                      // Unique ID for the menu
	Name         string               `bson:"name" validate:"required"`                                     // Name of the menu (e.g., "Lunch Specials")
	Description  string               `bson:"description,omitempty" json:"description" validate:"required"` // Optional description of the menu
	Category     Category             `bson:"category" json:"category" validate:"required"`                 // Category (e.g., "Main Course")
	Items        []primitive.ObjectID `bson:"items" json:"items"`                                           // List of food item IDs
	Availability bool                 `bson:"availability" json:"availability" validate:"required"`         // Menu availability (true/false)
	StartTime    string               `bson:"start_time,omitempty" json:"start_time,omitempty"`             // Menu start time (e.g., "11:00 AM")
	EndTime      string               `bson:"end_time,omitempty" json:"end_time,omitempty"`                 // Menu end time (e.g., "3:00 PM")
	CreatedAt    time.Time            `bson:"created_at" json:"created_at"`                                 // Timestamp when the menu was created
	UpdatedAt    time.Time            `bson:"updated_at" json:"updated_at"`                                 // Timestamp of the last menu update
	ImageURL     string               `bson:"image_url,omitempty" json:"image_url,omitempty"`               // Optional URL for menu image
	PriceRange   string               `bson:"price_range,omitempty" json:"price_range,omitempty"`           // Price range (e.g., "medium")
	Tags         []string             `bson:"tags,omitempty" json:"tags,omitempty"`                         // Optional tags like "Vegan", "Spicy", etc.
}

type Category int

const (
	MainCourse Category = iota
	Asides
	Desert
)

func (c Category) String() string {
	return [...]string{"Main Course", "Asides", "Dessert"}[c]
}
