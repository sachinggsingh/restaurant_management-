package controller

import (
	"context"
	"net/http"
	"resturnat-management/database"
	"resturnat-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var noteCollection *mongo.Collection = database.OpenCollection(database.Client, "note")

func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var note models.Note

		if err := c.BindJSON(&note); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(note)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		note.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		note.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		note.ID = primitive.NewObjectID()
		note.Note_id = note.ID.Hex()

		result, insertErr := noteCollection.InsertOne(ctx, note)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "note item was not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		result, err := noteCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing notes"})
		}

		var allNotes []bson.M

		if err = result.All(ctx, &allNotes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing notes"})
		}

		c.JSON(http.StatusOK, allNotes)
	}
}

// func GetNote()gin.HandlerFunc{
// 	return func(c *gin.Context){
// }
