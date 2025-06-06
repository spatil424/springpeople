package services

import (
	"fmt"
	"rest-api/interfaces"
	"rest-api/models"

	"github.com/sirupsen/logrus"
)

type OrderService interface {
	GetOrders(requestingUserID uint, requestingUserRole string) ([]models.Order, error)
	UpdateOrderByID(orderID uint, orderData models.OrderUpdateRequest, requestingUserID uint, requestingUserRole string) (*models.Order, error)
	// CreateOrder(order *models.Order, requestingUserID uint) (*models.Order, error)
}

type orderService struct {
	repo interfaces.OrderRepository
}

func NewOrderService(r interfaces.OrderRepository) OrderService {
	return &orderService{repo: r}
}

func (s *orderService) GetOrders(requestingUserID uint, requestingUserRole string) ([]models.Order, error) {
	logrus.Infof("User %d with role %s fetching orders", requestingUserID, requestingUserRole)
	if requestingUserRole == "admin" {
		return s.repo.FindAllOrders()
	}
	if requestingUserRole == "user" {
		return s.repo.FindOrdersByUserID(requestingUserID)
	}
	return nil, fmt.Errorf("unknown user role: %s", requestingUserRole)
}

func (s *orderService) UpdateOrderByID(orderID uint, orderData models.OrderUpdateRequest, requestingUserID uint, requestingUserRole string) (*models.Order, error) {
	logrus.Infof("User %d with role %s attempting to update order %d", requestingUserID, requestingUserRole, orderID)
	if requestingUserRole != "admin" {
		return nil, fmt.Errorf("user role '%s' cannot update orders", requestingUserRole)
	}

	existingOrder, err := s.repo.FindOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	if orderData.Product != nil {
		existingOrder.Product = *orderData.Product
	}
	if orderData.Quantity != nil {
		existingOrder.Quantity = *orderData.Quantity
	}
	if orderData.Price != nil {
		existingOrder.Price = *orderData.Price
	}

	return s.repo.UpdateOrder(existingOrder)
}
