package db

import (
	"context"
	"time"

	"github.com/tapiaw38/concejapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifiedRegister(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("concedb")
	col := db.Collection("users")

	register := make(map[string]interface{})

	if len(user.FirstName) > 0 {
		register["firstName"] = user.FirstName
	}
	if len(user.LastName) > 0 {
		register["lastName"] = user.LastName
	}

	register["dateBirth"] = user.DateBirth

	if len(user.Email) > 0 {
		register["email"] = user.Email
	}
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}

	updateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil

}
