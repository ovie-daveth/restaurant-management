package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order", controller.Getorders())
	incomingRoutes.GET("/order/:order_id", controller.Getorder())
	incomingRoutes.POST("/order/create", controller.CreateOrder())
	incomingRoutes.PATCH("/order/:order_id", controller.UpdateOrder())
}
