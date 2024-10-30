package main

import (
	"IMarket/config"
	"IMarket/middlewares"
	"IMarket/routes"
	"IMarket/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDatabase()
	router.POST("/login", utils.LoginHandler)
	authRoutes := router.Group("/auth", middlewares.AuthMiddleware())
	{
		authRoutes.POST("/products", routes.PostProduct)
		authRoutes.PUT("/products/:id", routes.PutProduct)
		authRoutes.DELETE("/products/:id", routes.DeleteProduct)
		authRoutes.POST("/orders", routes.PostOrder)
		authRoutes.PUT("/orders/:id", routes.PutOrder)
		authRoutes.DELETE("/orders/:id", routes.DeleteOrder)
		authRoutes.GET("/users", routes.GetUsers)
		authRoutes.POST("/users", routes.PostUser)
		authRoutes.GET("/users/:id", routes.GetUser)
		authRoutes.PUT("/users/:id", routes.PutUser)
		authRoutes.DELETE("/users/:id", routes.DeleteUser)
	}
	router.GET("/products", routes.GetProducts)
	router.GET("/products/:id", routes.GetProduct)
	router.GET("/orders", routes.GetOrders)
	router.GET("/orders/:id", routes.GetOrder)

	router.Run(":8080")
}
