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

func GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderId := c.Param("order_id")
		noteId := c.Param("note_id")
		var note models.Note

		objID, err := primitive.ObjectIDFromHex(noteId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note id"})
			return
		}
		filter := bson.M{
			"_id":      objID,
			"order_id": orderId,
		}

		err = noteCollection.FindOne(ctx, filter).Decode(&note)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching note"})
			return
		}
		c.JSON(http.StatusOK, note)
	}
}

// once the note is made for an order then hardly any chance that the customer will update it
// so no update note function is created here
