package db

import (
	"context"
	"fmt"
	"time"

	"github.com/tapiaw38/concejapi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnect.Database("concedb")
	col := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)

	//clean the password
	profile.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado")

		return profile, err
	}

	return profile, nil

}
