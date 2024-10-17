package main

import (
	"golang-restaurant-management/database"
	"golang-restaurant-management/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	_ "golang-restaurant-management/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Initialize the Gin router
	router := gin.New()
	router.Use(gin.Logger())

	// Swagger setup (no middleware)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Middleware for authentication on other routes
	// router.Use(middlewares.Authentication()) // Correctly apply middleware

	// Set up application-specific routes
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	// Start the server on the specified port
	router.Run(":" + port)
}
