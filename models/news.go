package models

import "time"

type News struct {
	User     string    `bson:"user" json:"user,omitempty"`
	Title    string    `bson:"title" json:"title,omitempty"`
	Category string    `bson:"category" json:"category,omitempty"`
	Body     string    `bson:"body" json:"body,omitempty"`
	Picture  string    `bson:"picture" json:"picture,omitempty"`
	Date     time.Time `bson:"date" json:"date,omitempty"`
}
