package main

import (
	"GoMongoDB/internal/database"
	"GoMongoDB/internal/handler"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	router.GET("/products", handler.GetProducts)
	router.POST("/products", handler.AddProducts)

	router.GET("/products/:id", handler.GetProductById)
	router.PATCH("/products/:id/stock", handler.UpdateProductScockById)
	router.PATCH("/products/:id/price", handler.UpdateProductPricekById)
	router.DELETE("/products/:id", handler.DeleteProductById)

	router.Run(":8080")
}
