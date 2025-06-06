package interfaces

import (
	"rest-api/models"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	FindOrdersByUserID(userID uint) ([]models.Order, error)
	FindAllOrders() ([]models.Order, error)
	FindOrderByID(orderID uint) (*models.Order, error)
	UpdateOrder(order *models.Order) (*models.Order, error)
}
