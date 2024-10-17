package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/Menu", controller.GetMenus())
	incomingRoutes.GET("/Menu/:Menu_id", controller.GetMenu())
	incomingRoutes.POST("/Menu/create", controller.CreateMenu())
	incomingRoutes.PATCH("/Menu/:Menu_id", controller.UpdateMenu())
}
