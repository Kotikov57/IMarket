package models

type Product struct { // Product структура продукта
	ID       int     `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" binding:"required,min=2,max=100"`
	Price    float64 `json:"price" binding:"required,gt=0"`
	Quantity int     `json:"quantity" binding:"required,gte=0"`
}

type Order struct { // Order структура заказа
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
	Address   string `json:"address"`
	UserID    int    `json:"user_id"`
	Status    string `json:"status" binding:"required"`
}

type User struct { // User структура данных о пользователе
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
