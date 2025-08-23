package routes

import (
	"resturnat-management/controller"

	"github.com/gin-gonic/gin"
)

// modi
func TableRouter(incomingRoutes *gin.Engine) {
	// Use middleware if needed
	incomingRoutes.POST("/table", controller.CreateTable())
	incomingRoutes.GET("/table/:id", controller.GetTable())
	incomingRoutes.GET("/tables", controller.GetAllTables())
	incomingRoutes.PATCH("/table/:id", controller.UpdateTable())
	// incomingRoutes.DELETE("/table/:id", controller.DeleteTable())
}
