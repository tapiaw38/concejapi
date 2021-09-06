package db

import (
	"context"
	"log"
	"time"

	"github.com/tapiaw38/concejapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadNews(ID string, paginate int64) ([]*models.NewsResponse, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("concedb")
	col := db.Collection("news")

	var result []*models.NewsResponse

	condition := bson.M{
		"user": ID,
	}

	//filter && ordering
	option := options.Find()
	option.SetLimit(10)
	option.SetSort(bson.D{{Key: "date", Value: -1}})
	option.SetSkip((paginate - 1) * 10)

	cursor, err := col.Find(ctx, condition, option)

	if err != nil {
		log.Fatal(err.Error())

		return result, false
	}

	for cursor.Next(context.TODO()) {
		var register models.NewsResponse
		err := cursor.Decode((&register))

		if err != nil {
			return result, false
		}
		result = append(result, &register)

	}

	return result, true

}
