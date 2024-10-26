package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	"log"
//	"fmt"
	"errors"
	"github.com/go-playground/validator/v10"
)

type Product struct { // Product структура продукта
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required,min=2,max=100"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Quantity int `json:"quantity" binding:"required,gte=0"`
}

type Order struct { // Order структура заказа
	ID int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity" binding:"required,gt=0"`
	Status string `json:"status" binding:"required"`
}

func getProducts(c *gin.Context) { // getProducts возвращает список всех продуктов
	var products []Product
	result := db.Find(&products)
	if result.Error != nil {
		c.JSON(400, gin.H{"error" : result.Error.Error()})
		return
	}
	c.JSON(200, products)
}

func postProduct(c *gin.Context) { // postProducts добавляет продукт
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

func getProduct(c *gin.Context){ // getProduct возвращает данные о продукте по его ID
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

func putProduct(c *gin.Context){ // putProduct обновляет данные о продукте
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

func deleteProduct(c *gin.Context) { // deleteProduct удаляет продукт
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

func getOrders(c *gin.Context) { // getOrders возвращает список всех заказов
	var orders []Order
	result := db.Find(&orders)
	if result.Error != nil {
		c.JSON(400, gin.H{"error" : result.Error.Error()})
		return
	}
	c.JSON(200, orders)
}

func postOrder(c *gin.Context) { // postOrder добавляет новый заказ
	var newOrder Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			c.JSON(400, gin.H {"error": err.Error()})
			return
		}
		if !isValidStatus(newOrder.Status) {
			c.JSON(400, gin.H {"error": "Invalid status"})
			return
		}
	var product Product
	if err := db.First(&product, newOrder.ProductID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if product.Quantity < newOrder.Quantity {
		c.JSON(400, gin.H{"error": "Not enough product quantity"})
		return
	}
	product.Quantity -= newOrder.Quantity
	if err := db.Save(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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

func getOrder(c *gin.Context){ // getOrder возвращает данные о заказе по его ID
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

func putOrder(c *gin.Context){ // putOrder обновляет данные о заказе
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

func deleteOrder(c *gin.Context) { // deleteOrder удаляет заказ
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

/* func statusValidator(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	return validStatuses[status]
}

func setupValidators() {
	validate = validator.New()
	validate.RegisterValidation("status", statusValidator)
} */

func isValidStatus(status string) bool {
    return validStatuses[status]
}

var db *gorm.DB
var validate *validator.Validate
var validStatuses = map[string]bool{
	"pending": true,
	"shipped": true,
	"delivered": true,
	"cancelled": true,
}

func main() {
	dsn := "host=localhost user=postgres password=fkla5283 dbname= imarket_db port=5432 sslmode=disable TimeZone=UTC"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}
	router := gin.Default()
//	setupValidators()

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