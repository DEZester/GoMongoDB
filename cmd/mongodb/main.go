package main

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
}
