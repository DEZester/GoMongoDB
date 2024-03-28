package handler

import (
	"GoMongoDB/internal/database"
	"GoMongoDB/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var body model.CreateProductRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := database.Products.InsertOne(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create product"})
		return
	}
	product := model.Product{
		ID:       res.InsertedID.(primitive.ObjectID),
		Name:     body.Name,
		Category: body.Category,
		Price:    body.Price,
		Stock:    body.Stock,
	}
	c.IndentedJSON(http.StatusCreated, product)
}

func GetProductById(c *gin.Context) {
	id := c.Param(":id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := database.Products.FindOne(c, primitive.M{"_id": _id})
	product := model.Product{}
	err = result.Decode(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find product"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)

}

func UpdateProductScockById(c *gin.Context) {

}

func UpdateProductPricekById(c *gin.Context) {

}

func DeleteProductById(c *gin.Context) {

}
