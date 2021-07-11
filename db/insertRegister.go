package db

import (
	"context"
	"time"

	"github.com/tapiaw38/concejapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRegister
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConnect.Database("concedb")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil

}
