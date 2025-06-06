package repository

import (
	"fmt"
	"rest-api/interfaces"
	"rest-api/models"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(database *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{db: database}
}

func (r *orderRepository) CreateOrder(order *models.Order) (*models.Order, error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, fmt.Errorf("could not create order: %w", err)
	}
	return order, nil
}

func (r *orderRepository) FindOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("could not find orders for user %d: %w", userID, err)
	}
	return orders, nil
}

func (r *orderRepository) FindAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("could not find all orders: %w", err)
	}
	return orders, nil
}

func (r *orderRepository) FindOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("order with ID %d not found", orderID)
		}
		return nil, fmt.Errorf("could not find order %d: %w", orderID, err)
	}
	return &order, nil
}

func (r *orderRepository) UpdateOrder(order *models.Order) (*models.Order, error) {
	if err := r.db.Save(order).Error; err != nil {
		return nil, fmt.Errorf("could not update order %d: %w", order.ID, err)
	}
	return order, nil
}
