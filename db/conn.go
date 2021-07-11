package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnect = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://tapiaw38:Walter153294@cluster0.bnlvy.mongodb.net/test")

//ConnectDB allows to connect to the db
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error(), "connection with db failed")
		return client
	}
	log.Println("Conection with db success!")

	return client
}

// CheckConnection allows check a db
func CheckConnection() bool {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
