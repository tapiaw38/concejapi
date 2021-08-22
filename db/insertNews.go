package db

import (
	"context"
	"time"

	"github.com/tapiaw38/concejapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertNews(n models.News) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("concedb")
	col := db.Collection("news")

	register := bson.M{
		"userId":   n.UserId,
		"title":    n.Title,
		"category": n.Category,
		"body":     n.Body,
		"date":     n.Date,
	}

	result, err := col.InsertOne(ctx, register)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil

}
