package controllers

import (
	"golang-restaurant-management/database"

	"go.mongodb.org/mongo-driver/mongo"
)

// Centralizing collection references to avoid duplication
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")
