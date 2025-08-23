package main

import (
	"os"
	"resturnat-management/database"
	"resturnat-management/routes"

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
	routes.UserRouter(router)
	router.Use(middleware.AuthMiddleware())

	routes.FoodRouter(router)
	routes.MenuRouter(router)
	routes.OrderRouter(router)
	routes.OrderItemRouter(router)
	routes.TableRouter(router)
	routes.InvoiceRouter(router)
	// router.NoteRouter(router)

	router.Run(":" + port)
}
