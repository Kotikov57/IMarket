package routes

import (
	"IMarket/config"
	"IMarket/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

var validStatuses = map[string]bool{ // validStatuses набор доступных статусов заказа
	"pending":   true,
	"shipped":   true,
	"delivered": true,
	"cancelled": true,
}

func GetProducts(c *gin.Context) { // getProducts возвращает список всех продуктов
	var products []models.Product
	result := config.Db.Find(&products)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, products)
}

func PostProduct(c *gin.Context) { // postProducts добавляет продукт
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := config.Db.Create(&newProduct)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Products post succesfully",
		"product": newProduct,
	})
}

func GetProduct(c *gin.Context) { // getProduct возвращает данные о продукте по его ID
	var product models.Product
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	result := config.Db.Find(&product, id)
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

func PutProduct(c *gin.Context) { // putProduct обновляет данные о продукте
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	result := config.Db.Model(&updatedProduct).Where("id = ?", id).Updates(updatedProduct)
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

func DeleteProduct(c *gin.Context) { // deleteProduct удаляет продукт
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	result := config.Db.Delete(&models.Product{}, id)
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

func GetOrders(c *gin.Context) { // getOrders возвращает список всех заказов
	var orders []models.Order
	result := config.Db.Find(&orders)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, orders)
}

func PostOrder(c *gin.Context) { // postOrder добавляет новый заказ
	var newOrder models.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !IsValidStatus(newOrder.Status) {
		c.JSON(400, gin.H{"error": "Invalid status"})
		return
	}
	var product models.Product
	if err := config.Db.First(&product, newOrder.ProductID).Error; err != nil {
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
	if err := config.Db.Save(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	result := config.Db.Create(&newOrder)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Order post succesfully",
		"order":   newOrder,
	})
}

func GetOrder(c *gin.Context) { // getOrder возвращает данные о заказе по его ID
	var order models.Order
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}
	result := config.Db.Find(&order, id)
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

func PutOrder(c *gin.Context) { // putOrder обновляет данные о заказе
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	var updatedOrder models.Order
	if !IsValidStatus(updatedOrder.Status) {
		c.JSON(400, gin.H{"error": "Invalid status"})
		return
	}
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	result := config.Db.Model(&updatedOrder).Where("id = ?", id).Updates(updatedOrder)
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
		"order":   updatedOrder,
	})
}

func DeleteOrder(c *gin.Context) { // deleteOrder удаляет заказ
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}
	result := config.Db.Delete(&models.Order{}, id)
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

func GetUsers(c *gin.Context) { // getUsers возвращает данные о пользователях
	var users []models.User
	result := config.Db.Find(&users)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, users)
}

func PostUser(c *gin.Context) { // postUser добавляет нового пользователя
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := config.Db.Create(&newUser)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "User post succesfully",
		//		"user" : newUser,
	})
}

func GetUser(c *gin.Context) { // getUser возвращает данные о пользователе по его ID
	var user models.User
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	result := config.Db.Find(&user, id)
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

func PutUser(c *gin.Context) { // putUser обновляет данные о пользователе
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := config.Db.Model(&updatedUser).Where("id = ?", id).Updates(updatedUser)
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

func DeleteUser(c *gin.Context) { // deleteUser удаляет пользователя
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	result := config.Db.Delete(&models.User{}, id)
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
