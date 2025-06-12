package repository

import "restapi/models"

type OrderRepositoryImpl struct{}

func (o *OrderRepositoryImpl) BuyOrders(order models.Orders) (orderId int, err error) {
	buyedOrder := order.OrderId + 1
	return buyedOrder, nil
}

func (o *OrderRepositoryImpl) ExecuteOrders() {
	return
}
