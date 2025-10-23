package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_id   string             `json:"order_id"`
	Order_date time.Time          `json:"order_date" validate:"required"`
	Table_id   *string            `json:"table_id" validate:"required"`
	// Food_id      *string            `json:"food_id" validate:"required"`
	// Quantity     *int               `json:"quantity" validate:"required,min=1"`
	// Price        *float64           `json:"price" validate:"required"`
	// Order_Status *string            `json:"order_status" validate:"required,eq=PENDING|eq=COMPLETED|eq=CANCELLED"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
