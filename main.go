package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"resturnat-management/config"
	"resturnat-management/routes"

	middleware "resturnat-management/middleware"

	"github.com/gin-gonic/gin"
)

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "Food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		fmt.Println("Welcome to the Restaurant Management System")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routes.UserRouter(router)
	router.Use(middleware.Authentication())

	if err := config.InitRedis(); err != nil {
		log.Fatal(err)
	}

	routes.FoodRouter(router)
	routes.MenuRouter(router)
	routes.OrderRouter(router)
	routes.OrderItemRouter(router)
	routes.TableRouter(router)
	routes.InvoiceRouter(router)
	routes.NoteRouter(router)

	router.Run(":" + port)
}
