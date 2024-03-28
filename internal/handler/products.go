package handler

import (
	"GoMongoDB/internal/database"
	"GoMongoDB/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts(c *gin.Context) {
	cursor, err := database.Products.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get products"})
		return
	}

	var products []model.Product
	if err := cursor.All(c, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get products"})
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

func AddProducts(c *gin.Context) {

}

func GetProductById(c *gin.Context) {

}

func UpdateProductScockById(c *gin.Context) {

}

func UpdateProductPricekById(c *gin.Context) {

}

func DeleteProductById(c *gin.Context) {

}
