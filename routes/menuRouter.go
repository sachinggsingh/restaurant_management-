package routes

import (
	controller "resturnat-management/controller"

	"github.com/gin-gonic/gin"
)

func MenuRouter(incomingRoutes *gin.Engine) {
	// Use middleware if needed
	incomingRoutes.POST("/menu", controller.CreateMenu())
	incomingRoutes.GET("/menu/:id", controller.GetMenu())
	incomingRoutes.GET("/menus", controller.GetAllMenus())
	incomingRoutes.PATCH("/menu/:id", controller.UpdateMenu())
	// incomingRoutes.DELETE("/menu/:id", controller.DeleteMenu())
}
