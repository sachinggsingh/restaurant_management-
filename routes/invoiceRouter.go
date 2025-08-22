package routes

import (
	"resturnat-management/controller"

	"github.com/gin-gonic/gin"
)

func InvoiceRouter(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/invoice", controller.CreateInvoice())
	incommingRoutes.GET("/invoice/:id", controller.GetInvoice())
	incommingRoutes.GET("/invoices", controller.GetAllInvoices())
	incommingRoutes.PATCH("/invoice/:id", controller.UpdateInvoice())
}
