package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/Table", controller.GetTables())
	incomingRoutes.GET("/Table/:Table_id", controller.GetTable())
	incomingRoutes.POST("/Table/create", controller.CreateTable())
	incomingRoutes.PATCH("/Table/:Table_id", controller.UpdateTable())
}
