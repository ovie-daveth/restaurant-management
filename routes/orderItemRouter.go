package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetorderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetorderItem())
	incomingRoutes.GET("/orderItems-order/:orderI_id", controller.GetorderItemsByOrder())
	incomingRoutes.POST("/orderItems/create", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
