package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
}

func main() {
	router := gin.Default()
	products := []Product{
		{1, "Utyug", 175.0, 5},
		{2, "Chainik", 123.56, 6},
		{3, "Televisor", 499.99, 2},
	}
	orders := []Order{
		{1, 1, 3},
		{2, 1, 2},
		{3, 3, 1},
	}

	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, products)
	})

	router.POST("/products", func (c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(400, gin.H {"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message" : "Products post succesfully",
			"product" : newProduct,
		})
	})

	router.GET("/products/:id", func (c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid product ID"})
			return
		}
		for _, product := range products {
			if product.ID == id {
				c.JSON(200, product)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Product not found"})
	})

	router.PUT("/products/:id", func (c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid product ID"})
			return
		}
		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		for _, product := range products {
			if product.ID == id {
				product = updatedProduct
				c.JSON(200, product)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Product not found"})
	})

	router.DELETE("/products/:id", func (c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error" : "Invalid product ID"})
			return
		}
		index := -1
		for i, _ := range products {
			if products[i].ID == id {
				index = i
				break
			}
		}
		if index == -1 {
			c.JSON(404, gin.H{"error" : "Product not found"})
			return
		}
		products = append(products[:index],products[index + 1:]...)
		c.JSON(200, gin.H{"message": "Product deleted"})
	})

	router.GET("/orders", func (c *gin.Context){
		c.JSON(200, orders)
	})

	router.POST("/orders", func (c *gin.Context){
		var newOrder Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message" : "Order succesfully added",
			"order" : newOrder,
		})
	})

	router.GET("/orders/:id")
	
	router.Run(":8080")
}