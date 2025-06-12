package services

import (
	"fmt"
	"src/src/springpeople/Assignment2/models"
	"time"
)

func MatchingEngine(orderChan <-chan models.Order) {
	//array of buy order
	buyOrders := []models.Order{}

	// array of sell orders
	sellOrders := []models.Order{}

	fmt.Println("Matching Engine: Started.")

	// read the orderChan if the ordertypes is by add to buy-array and sell add to sell array
	for receivedOrder := range orderChan {
		fmt.Printf("Matching Engine: Received order %s (Type: %s, Price: %d, Qty: %d)\n", receivedOrder.OrderId, receivedOrder.Type, receivedOrder.Price, receivedOrder.Quantity)

		if receivedOrder.Type == "BUY" {
			buyOrders = append(buyOrders, receivedOrder)
		} else {
			sellOrders = append(sellOrders, receivedOrder)
		}
		// call matchOrders
		buyOrders, sellOrders = matchOrders(buyOrders, sellOrders)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Matching Engine: Order channel closed. ")

	if len(buyOrders) > 0 {
		fmt.Println("\nMatching Engine: Remaining unmatched BUY orders:")
		for _, order := range buyOrders {
			fmt.Printf("  - %s (Price: %d, Qty: %d)\n", order.OrderId, order.Price, order.Quantity)
		}
	}
	if len(sellOrders) > 0 {
		fmt.Println("\nMatching Engine: Remaining unmatched SELL orders:")
		for _, order := range sellOrders {
			fmt.Printf("  - %s (Price: %d, Qty: %d)\n", order.OrderId, order.Price, order.Quantity)
		}
	}

	fmt.Println("Matching Engine: Shut down gracefully.")

}

func matchOrders(currBuyOrders, currSellOrders []models.Order) (pendingBuy []models.Order, pendingSell []models.Order) {
	// pending buy and pending sell []
	pendingBuy = []models.Order{}
	pendingSell = []models.Order{}

	for i := 0; i < len(currBuyOrders); i++ {
		buyOrder := &currBuyOrders[i]

		if buyOrder.Quantity == 0 {
			continue
		}

		for j := 0; j < len(currSellOrders); j++ {
			sellOrder := &currSellOrders[j]
			if sellOrder.Quantity == 0 {
				continue
			}

			if (*buyOrder).Price >= (*sellOrder).Price {
				matchedQuantity := min((*buyOrder).Quantity, (*sellOrder).Quantity)
				fmt.Printf("BuyId %s SellId %s Price %d, Quantity %d\n", (*buyOrder).OrderId, (*sellOrder).OrderId, (*sellOrder).Price, matchedQuantity)
				fmt.Printf("  +++ TRADE: BuyID %s, SellID %s, Matched Price %d, Matched Quantity %d\n",
					buyOrder.OrderId, sellOrder.OrderId, sellOrder.Price, matchedQuantity)

				(*buyOrder).Quantity -= matchedQuantity
				(*sellOrder).Quantity -= matchedQuantity

				if (*buyOrder).Quantity == 0 {
					break
				}
			}
		}
	}

	for _, order := range currBuyOrders {
		if order.Quantity > 0 {
			pendingBuy = append(pendingBuy, order)
		}
	}

	for _, order := range currSellOrders {
		if order.Quantity > 0 {
			pendingSell = append(pendingSell, order)
		}
	}
	return pendingBuy, pendingSell
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
