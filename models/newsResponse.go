package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This is the structure which I return the news.
type NewsResponse struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	User     string             `bson:"user" json:"user,omitempty"`
	Category string             `bson:"category" json:"category,omitempty"`
	Body     string             `bson:"body" json:"body,omitempty"`
	Picture  string             `bson:"picture" json:"picture,omitempty"`
	Date     time.Time          `bson:"date" json:"date,omitempty"`
}
