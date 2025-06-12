package interfaces

import "restapi/models"

type OrderRepository interface {
	BuyOrders(models.Orders) (orderId int, err error)
	ExecuteOrders()
}
