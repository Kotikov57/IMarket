package utils

import (
	"IMarket/config"
	"IMarket/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"
)

var JwtSecret = []byte("your_secret_key")

func GenerateJWT(userID int) (string, error) { // GenerateJWT генерирует уникальный JWT
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func LoginHandler(c *gin.Context) { // LoginHandler генерирует ключ для конкретного пользователя
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var users []models.User
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	result := config.Db.Select("id").Where("email = ? AND password = ?", loginData.Email, loginData.Password).Find(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
	}
	token, err := GenerateJWT(users[0].ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
