package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	// "resturnat-management/config"
	"resturnat-management/database"
	"resturnat-management/models"
	"strconv"
	"time"

	config "resturnat-management/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetAllFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Retrieve pagination params as before
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		if s := c.Query("startIndex"); s != "" {
			startIndex, _ = strconv.Atoi(s)
		}

		// Create a unique cache key for this query
		cacheKey := fmt.Sprintf("foods:page=%d:perPage=%d:startIndex=%d", page, recordPerPage, startIndex)

		// Check Redis cache
		cached, err := config.RDB.Get(ctx, cacheKey).Result()
		if err == nil {
			// Cache hit - unmarshal and return cached data
			var cachedData bson.M
			if err := json.Unmarshal([]byte(cached), &cachedData); err == nil {
				c.JSON(http.StatusOK, cachedData)
				return
			}
			// If unmarshal fails, continue to fetch from DB
		}

		// Cache miss - fetch from MongoDB
		matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
		groupStage := bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			}},
		}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "total_count", Value: 1},
				{Key: "food_items", Value: bson.D{
					{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}},
				}},
			}},
		}

		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing food items"})
			return
		}

		var allFoods []bson.M
		if err = result.All(ctx, &allFoods); err != nil {
			log.Fatal(err)
		}
		if len(allFoods) == 0 {
			response := bson.M{"total_count": 0, "food_items": []interface{}{}}
			// Cache empty result with shorter TTL
			jsonData, _ := json.Marshal(response)
			config.RDB.Set(ctx, cacheKey, jsonData, 1*time.Minute)
			c.JSON(http.StatusOK, response)
			return
		}

		// Cache the result with TTL
		jsonData, _ := json.Marshal(allFoods[0])
		config.RDB.Set(ctx, cacheKey, jsonData, 5*time.Minute)

		c.JSON(http.StatusOK, allFoods[0])
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("food_id")
		cacheKey := fmt.Sprintf("food:%s", foodId)

		// Try to get cached food
		cached, err := config.RDB.Get(ctx, cacheKey).Result()
		if err == nil {
			// Cache hit - unmarshal and return cached item
			var food models.Food
			if err := json.Unmarshal([]byte(cached), &food); err == nil {
				c.JSON(http.StatusOK, food)
				return
			}
			// If unmarshal fails, query DB
		}

		// Cache miss - fetch from MongoDB
		var food models.Food
		err = foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the food item"})
			return
		}

		// Cache the food data
		jsonData, _ := json.Marshal(food)
		config.RDB.Set(ctx, cacheKey, jsonData, 10*time.Minute)

		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var food models.Food
		var menu models.Menu
		defer cancel()

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validating
		validateError := validate.Struct(food)
		if validateError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateError.Error()})
			return
		}

		// finding if the nemu exist or not
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu did not found"})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Food_id = primitive.NewObjectID().Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "food item was not created"})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}
func round(num float64) int {
	return int(num + 0.5)
}

func toFixed(num float64, precision int) float64 {
	return float64(round(num*math.Pow(10, float64(precision)))) / math.Pow(10, float64(precision))
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var menu models.Menu
		var food models.Food
		defer cancel()

		foodId := c.Param("food_id")

		if err := c.BindJSON((&food)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D

		if food.Name != nil {
			updateObj = append(updateObj, bson.E{Key: "name", Value: food.Name})
		}
		if food.Price != nil {
			updateObj = append(updateObj, bson.E{Key: "price", Value: food.Price})
		}
		if food.Food_image != nil {
			updateObj = append(updateObj, bson.E{Key: "food_image", Value: food.Food_image})
		}
		if food.Menu_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "menu not found"})
				defer cancel()
				return
			}
		}
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: food.Updated_at})

		upsert := true

		filter := bson.M{"food_id": foodId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := foodCollection.UpdateOne(ctx, filter, bson.D{
			{Key: "$set", Value: updateObj},
		}, &opt,
		)
		if err != nil {
			msg := "food item update failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"message": "food item updated successfully", "data": result})

	}
}
