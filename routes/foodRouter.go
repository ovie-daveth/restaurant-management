package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/food", controller.GetFoods())
	incomingRoutes.GET("/food/:food_id", controller.GetFood())
	incomingRoutes.POST("/food/create", controller.CreateFood())
	incomingRoutes.PATCH("/food/:food_id", controller.UpdateFood())
}
