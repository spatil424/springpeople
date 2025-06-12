package main

import (
	"fmt"
	"src/src/springpeople/Assignment2/models"
	"src/src/springpeople/Assignment2/services"
	"sync"
)

// the main is the entry point for the program
func main() {

	numberOfClients := 10
	numberOfOrdersPerClient := 10
	var waitGroup sync.WaitGroup

	// create an order channel that takes the size as 10
	orderChan := make(chan models.Order, 100)

	//results channel
	//resultsChan := make(chan struct{}, 100)

	// main will call client and client will execute orders
	fmt.Println("starting clients")
	for i := 0; i < numberOfClients; i++ {
		waitGroup.Add(1)
		go services.Client(i, numberOfOrdersPerClient, orderChan, &waitGroup)
	}

	fmt.Println("starting matching engine")
	go services.MatchingEngine(orderChan)

	fmt.Println("\n--- Main: Waiting for all clients to finish submitting orders... ---")

	waitGroup.Wait()
	fmt.Println("--- Main: All clients have finished submitting orders. ---")

	close(orderChan)
	fmt.Println("--- Main: Order channel closed. Matching engine will process remaining orders and shut down. ---")

}
