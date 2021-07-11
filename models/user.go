package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	Id        primitive.ObjectID `bson: "_id,omitempty" json:"id"`
	FirstName string             `bson: "firstName" json:"fistName,omitempty"`
	LastName  string             `bson: "lastName" json:"lastName,omitempty"`
	DateBirth time.Time          `bson: "birthday" json:"dateBirth,omitempty"`
	Email     string             `bson: "email" json:"email"`
	Password  string             `bson: "password" json:"password,omitempty"`
	Avatar    string             `bson: "avatar" json:"avatar,omitempty"`
	Banner    string             `bson: "banner" json:"banner,omitempty"`
}
