package routes

import (
	controller "resturnat-management/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(incommingRoutes *gin.Engine) {

	// Use middleware if needed
	incommingRoutes.POST("/signup", controller.Sign())
	incommingRoutes.POST("/login", controller.Login())

	incommingRoutes.GET("/user/:id", controller.GetUser())
	incommingRoutes.GET("/users", controller.GetAllUsers())

	// if i want to i will in future if it is needed
	// incommingRouter.PUT("/user/:id", controller.UpdateUser)
	// incommingRouter.DELETE("/user/:id", controller.DeleteUser)
}
