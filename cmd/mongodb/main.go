package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)                                //create object that gives API, access to MongoDB
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI) //options for user for connect to DB

	client, err := mongo.Connect(context.Background(), opts) //connect to DB
	if err != nil {
		panic(err)
	}
}
