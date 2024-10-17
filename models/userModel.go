package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`                               // Unique identifier for the user
	FirstName *string            `json:"first_name" validate:"required,min=2,max=100"` // User's first name
	LastName  *string            `json:"last_name" validate:"required,min=2,max=100"`  // User's last name
	Email     *string            `json:"email" validate:"email,required"`              // User's email (unique for login)
	Password  *string            `json:"password" validate:"required,min=6"`           // Hashed password for authentication
	Phone     *string            `json:"phone" validate:"required"`                    // User's contact number
	Address   *string            `json:"address" validate:"required"`                  // User's home/shipping address
	CreatedAt time.Time          `json:"created_at"`                                   // Timestamp for when the user was created
	UpdatedAt time.Time          `json:"updated_at"`                                   // Timestamp for the last update of the user info
	UserID    string             `json:"user_id"`                                      // User's unique ID for the app (human-readable)
}
