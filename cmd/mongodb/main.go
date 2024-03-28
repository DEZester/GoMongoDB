package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Can't load .env file")
		return
	}
	mongoURI := os.Getenv("MONGO_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)                  //create object that gives API, access to MongoDB
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI) //options for user for connect to DB
	client, err := mongo.Connect(context.Background(), opts)                   //connect to DB
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database started successfully")
}
