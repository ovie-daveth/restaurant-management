package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", controller.GetNotes())
	incomingRoutes.GET("/notes/:notes_id", controller.GetNote())
	incomingRoutes.POST("/notes/create", controller.CreateNote())
	incomingRoutes.PATCH("/notes/:notes_id", controller.UpdateNote())
}
