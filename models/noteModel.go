package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID         primitive.ObjectID `bson:"_id"`
	Note_id    string             `json:"note_id"`
	Order_id   string             `json:"order_id" validate:"required"`
	Title      string             `json:"title" validate:"required,min=2,max=100"`
	Text       string             `json:"note" validate:"required,min=2"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
