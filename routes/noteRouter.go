package routes

import (
	"resturnat-management/controller"

	"github.com/gin-gonic/gin"
)

func NoteRouter(incomingRoutes *gin.Engine) {

	// if want to put like /note/createNote like this then grouping can be used
	// noteRoutes := incomingRoutes.Group("/note")
	incomingRoutes.POST("/createNote", controller.CreateNote())
	incomingRoutes.GET("/getNotes", controller.GetNotes())
	incomingRoutes.GET("/getNote/:note_id", controller.GetNote())
}
