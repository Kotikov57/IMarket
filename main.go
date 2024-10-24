package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
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

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=fkla5283 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	router := gin.Default()
	products := []Product{
		{1, "Utyug", 175.0, 5},
		{2, "Chainik", 123.56, 6},
		{3, "Televisor", 499.99, 2},
	}
	orders := []Order{
		{1, 1, 3, "Registered"},
		{2, 1, 2, "Complete"},
		{3, 3, 1, "Cancelled"},
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
		index := InProducts(id, products)
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
		index := InProducts(newOrder.ProductID, products)
		if  index == -1 {
			c.JSON(404, gin.H{"error": "Product not found"})
			return
		}
		remainder := products[index].Quantity - newOrder.Quantity
		if remainder < 0{
			c.JSON(400, gin.H{"error" : "There is not so much product"})
			return
		}
		products[index].Quantity = remainder
		c.JSON(200, gin.H{
			"message" : "Order succesfully added",
			"order" : newOrder,
		})
	})

	router.GET("/orders/:id", func (c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error" : err.Error()})
			return
		}
		for _, order := range orders {
			if order.ID == id {
				c.JSON(200, order)
				return
			}
		}
		c.JSON(404, gin.H{"error" : "Order not found"})
	})
	
	router.PUT("/orders/:id", func (c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error" : err.Error()})
			return
		}
		var updatedOrder Order
		if err := c.ShouldBindJSON(&updatedOrder); err != nil{
			c.JSON(400, gin.H{"error" : err.Error()})
			return
		}
		for _, order := range orders {
			if order.ID == id {
				order = updatedOrder
				c.JSON(200, gin.H{
					"message" : "Orded succesfulle updated",
					"order" : updatedOrder,
				})
				return
			}
		}
		c.JSON(404, gin.H{"error" : "Order not found"})
	})

	router.DELETE("/orders/:id", func(c *gin.Context){
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error" : err.Error()})
			return
		}
		index := InOrders(id, orders)
		if index == -1 {
			c.JSON(404, gin.H{"error" : "Order not found"})
			return
		}
		orders = append(orders[index:],orders[:index + 1]...)
		c.JSON(200, gin.H{"message" : "Order succesfully deleted"})
	})

	router.Run(":8080")
}