package routes

import (
	"resturnat-management/controller"

	"github.com/gin-gonic/gin"
)

func OrderRouter(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/order", controller.CreateOrder())
	incommingRoutes.GET("/order/:order_id", controller.GetOrder())
	incommingRoutes.GET("/orders", controller.GetAllOrders())
	incommingRoutes.PATCH("/order/:id", controller.UpdateOrder())
}
