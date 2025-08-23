package routes

import (
	"resturnat-management/controller"

	"github.com/gin-gonic/gin"
)

func OrderItemRouter(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/orderitem", controller.CreateOrderItem())
	incommingRoutes.GET("/orderitem/:id", controller.GetOrderItem())
	incommingRoutes.GET("/orderitems", controller.GetAllOrderItems())
	incommingRoutes.GET("/orderitems-order/:id", controller.GetOrderItemsByOrder())
	incommingRoutes.PATCH("/orderitem/:id", controller.UpdateOrderItem())
}
