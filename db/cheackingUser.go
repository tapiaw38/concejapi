package db

import (
	"context"
	"time"

	"github.com/tapiaw38/concejapi/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheackingUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("concedb")
	col := db.Collection("users")

	conditions := bson.M{"email": email}

	var result models.User
	err := col.FindOne(ctx, conditions).Decode(&result)

	ID := result.Id.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
