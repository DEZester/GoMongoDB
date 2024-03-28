package main

import (
	"GoMongoDB/internal/database"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Can't load .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")

	connectToDB := database.Init(mongoURI, "development")
	if connectToDB != nil {
		panic(connectToDB)
	}
	fmt.Println("Connected to MongoDB")

	defer func() {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}()
}
