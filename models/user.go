package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
	DateBirth time.Time          `bson:"dateBirth,omitempty"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password,omitempty"`
	Avatar    string             `bson:"avatar,omitempty"`
	Banner    string             `bson:"banner,omitempty"`
}
