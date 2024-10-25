package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	"log"
//	"fmt"
	"errors"
)

type Product struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
	Status string `json:"status"`
}

func InProducts(id int, products []Product) int { // InProducts возвращает позицию объекта в списке
	for i, product := range products {
		if product.ID == id{
			return  i
		}
	}
	return -1
}

func InOrders(id int, orders []Order) int { // InOrders возвращает позицию объекта в списке
	for i, order := range orders {
		if order.ID == id{
			return  i
		}
	}
	return -1
}

func getProducts(c *gin.Context) {
	var products []Product
	result := db.Find(&products)
	if result.Error != nil {
		c.JSON(400, gin.H{"error" : result.Error.Error()})
		return
	}
	c.JSON(200, products)
}

func postProduct(c *gin.Context) {
	var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(400, gin.H {"error": err.Error()})
			return
		}
		result := db.Create(&newProduct)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message" : "Products post succesfully",
			"product" : newProduct,
		})
}

func getProduct(c *gin.Context){
	var product Product
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	result := db.Find(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Product not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(200, product)
}

func putProduct(c *gin.Context){
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
	result := db.Model(&updatedProduct).Where("id = ?", id).Updates(updatedProduct)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Product not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(200, gin.H{
		"message": "Product succesfully updated",
		"product": updatedProduct,
	})
}

func deleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error" : "Invalid product ID"})
		return
	}
	result := db.Delete(&Product{}, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Product not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(200, gin.H{"message": "Product deleted"})
}

func getOrders(c *gin.Context) {
	var orders []Order
	result := db.Find(&orders)
	if result.Error != nil {
		c.JSON(400, gin.H{"error" : result.Error.Error()})
		return
	}
	c.JSON(200, orders)
}

func postOrder(c *gin.Context) {
	var newOrder Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			c.JSON(400, gin.H {"error": err.Error()})
			return
		}
		result := db.Create(&newOrder)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message" : "Order post succesfully",
			"order" : newOrder,
		})
}

func getOrder(c *gin.Context){
	var order Order
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}
	result := db.Find(&order, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Order not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(200, order)
}

func putOrder(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	var updatedOrder Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	result := db.Model(&updatedOrder).Where("id = ?", id).Updates(updatedOrder)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Order not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(200, gin.H{
		"message": "Order succesfully updated",
		"order": updatedOrder,
	})
}

func deleteOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error" : "Invalid order ID"})
		return
	}
	result := db.Delete(&Order{}, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Order not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(200, gin.H{"message": "Order deleted"})
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=fkla5283 dbname= imarket_db port=5432 sslmode=disable TimeZone=UTC"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	router := gin.Default()

	router.GET("/products", getProducts)
	router.POST("/products", postProduct)
	router.GET("/products/:id", getProduct)
	router.PUT("/products/:id", putProduct)
	router.DELETE("/products/:id", deleteProduct)
	router.GET("/orders", getOrders)
	router.POST("/orders", postOrder)
	router.GET("/orders/:id", getOrder)
	router.PUT("/orders/:id", putOrder)
	router.DELETE("/orders/:id", deleteOrder)

	router.Run(":8080")
}