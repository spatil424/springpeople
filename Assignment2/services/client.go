package services

import (
	"fmt"
	"math/rand"
	"src/src/springpeople/Assignment2/models"
	"sync"
	"time"
)

func Client(clientId int, orderPerClient int, orderChan chan<- models.Order, wg *sync.WaitGroup) {
	// generate buy and sell orders randomy and pass it to the orderschannel

	defer wg.Done()
	fmt.Printf("Client %d: Starting to submit %d orders.\n", clientId, orderPerClient)

	for i := 0; i < orderPerClient; i++ {
		orderId := fmt.Sprintf("Order id is %d_%d", clientId, i)
		orderType := ""
		if rand.Intn(2) == 0 {
			orderType = "BUY"
		} else {
			orderType = "SELL"
		}
		orderPrice := rand.Intn(101) + 50
		orderQuantity := rand.Intn(10) + 1

		order := models.Order{
			OrderId:  orderId,
			Type:     orderType,
			Price:    orderPrice,
			Quantity: orderQuantity,
		}

		orderChan <- order
		fmt.Printf("Client %d: Submitted %s order %s (Price: %d, Qty: %d)\n", clientId, order.Type, order.OrderId, order.Price, order.Quantity)

		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Finished submitting orders")

}
