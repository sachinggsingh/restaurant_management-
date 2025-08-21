package routes

import (
	controller "resturnat-management/controller"

	"github.com/gin-gonic/gin"
)

func FoodRouter(incommingRoutes *gin.Engine) {
	// Use middleware if needed
	incommingRoutes.POST("/food", controller.CreateFood())
	incommingRoutes.GET("/food/:id", controller.GetFood())
	incommingRoutes.GET("/foods", controller.GetAllFoods())
	incommingRoutes.PUT("/food/:id", controller.UpdateFood())
	incommingRoutes.DELETE("/food/:id", controller.DeleteFood())
}
