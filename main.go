package main

import (
	"os"
	"resturnat-management/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	//  middle"resturant-management/middleware"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "Food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.UserRouter(router)
	// router.Use(middleware.AuthMiddleware())

	router.FoodRouter(router)
	router.MenuRouter(router)
	// router.OrderRouter(router)
	// router.OrderItemRouter(router)
	// router.TableRouter(router)
	// router.InvoiceRouter(router)
	// router.NoteRouter(router)

	router.Run(":" + port)
}
