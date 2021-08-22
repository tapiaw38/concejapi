package models

import "time"

type News struct {
	UserId   string    `bson:"userId" json:"userId,omitempty"`
	Title    string    `bson:"title" json:"title,omitempty"`
	Category string    `bson:"category" json:"category,omitempty"`
	Body     string    `bson:"body" json:"body,omitempty"`
	Date     time.Time `bson:"date" json:"date,omitempty"`
}
