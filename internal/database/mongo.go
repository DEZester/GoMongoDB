package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init(uri, database string) error {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)             //create object that gives API, access to MongoDB
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI) //options for user for connect to DB

	localClient, err := mongo.Connect(context.Background(), opts) //connect to DB
	if err != nil {
		return err

	}

	client = localClient

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	return err
}

func Close() error {
	return client.Disconnect(context.Background())
}
