package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	"log"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

var db *gorm.DB
var validStatuses = map[string]bool{
	"pending": true,
	"shipped": true,
	"delivered": true,
	"cancelled": true,
}
var jwtSecret = []byte("your_secret_key")

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
	Address string `json:"address"`
	UserID int `json:"user_id"`
	Status string `json:"status" binding:"required"`
}

type User struct { // User структура данных о пользователе
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func GenerateJWT(userID int) (string, error) { // GenerateJWT генерирует уникальный JWT
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func AuthMiddleware() gin.HandlerFunc { // AuthMiddleware проверяет корректность запроса с авторизацией и корректность ключа
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("userID", claims["userID"])
		c.Next()
	}
}

func LoginHandler(c *gin.Context) { // LoginHandler генерирует ключ для конкретного пользователя
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var users []User
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	result := db.Select("id").Where("email = ? AND password = ?", loginData.Email, loginData.Password).Find(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error" : result.Error.Error()})
	}
	token, err := GenerateJWT(users[0].ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
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
		if !IsValidStatus(newOrder.Status) {
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

func getOrder(c *gin.Context) { // getOrder возвращает данные о заказе по его ID
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

func putOrder(c *gin.Context) { // putOrder обновляет данные о заказе
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	var updatedOrder Order
	if !IsValidStatus(updatedOrder.Status) {
		c.JSON(400, gin.H {"error": "Invalid status"})
		return
	}
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

func getUsers(c *gin.Context) { // getUsers возвращает данные о пользователях
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, users)
}

func postUser(c *gin.Context) { // postUser добавляет нового пользователя
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H {"error": err.Error()})
		return
	}
	result := db.Create(&newUser)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message" : "User post succesfully",
//		"user" : newUser,
	})
}

func getUser(c *gin.Context) { // getUser возвращает данные о пользователе по его ID
	var user User
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	result := db.Find(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(200, user)
}

func putUser(c *gin.Context) { // putUser обновляет данные о пользователе
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	} 
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := db.Model(&updatedUser).Where("id = ?",id).Updates(updatedUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, result.Error.Error())
		return
	}
	c.JSON(200, gin.H{"message": "User succesfully updated"})
}

func deleteUser(c *gin.Context) { // deleteUser удаляет пользователя
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error" : "Invalid user ID"})
		return
	}
	result := db.Delete(&User{}, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(200, gin.H{"message": "User deleted"})
}

func IsValidStatus(status string) bool { // IsValidStatus проверяет, корректен ли статус заказа
    return validStatuses[status]
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
	router.POST("/login", LoginHandler)
	authRoutes := router.Group("/auth", AuthMiddleware())
	{
		authRoutes.POST("/products", postProduct)
		authRoutes.PUT("/products/:id", putProduct)
		authRoutes.DELETE("/products/:id", deleteProduct)
		authRoutes.POST("/orders", postOrder)
		authRoutes.PUT("/orders/:id", putOrder)
		authRoutes.DELETE("/orders/:id", deleteOrder)
		authRoutes.GET("/users", getUsers)
		authRoutes.POST("/users", postUser)
		authRoutes.GET("/users/:id", getUser)
		authRoutes.PUT("/users/:id", putUser)
		authRoutes.DELETE("/users/:id", deleteUser)
	}
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.GET("/orders", getOrders)
	router.GET("/orders/:id", getOrder)
	
	router.Run(":8080")
}