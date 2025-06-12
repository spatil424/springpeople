package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	totalItems := 100
	numWorkers := 10

	var wg sync.WaitGroup

	itemChan := make(chan int, totalItems)

	fmt.Println("Populating conveyor belt with items...")
	for i := 0; i < totalItems; i++ {
		itemChan <- i
	}

	fmt.Println("Conveyor belt populated.")

	close(itemChan)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go worker(i+1, itemChan, &wg)
	}

	wg.Wait()
	fmt.Println("\nAll items have been processed.")
	fmt.Println("hello")
}

func worker(workerID int, itemChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range itemChan {
		// fmt.Printf("Worker %d: Picked item %d\n", workerID, item)
		// time.Sleep(100 * time.Millisecond)
		// fmt.Printf("Worker packed item %d\n\n", item)

		fmt.Printf("Worker %d: Picked item %d\n", workerID, item)
		time.Sleep(990 * time.Millisecond)
		fmt.Printf("Worker %d: Packed item %d\n\n", workerID, item)
	}
	fmt.Printf("Worker %d: Finished. No more items to process.\n", workerID)

}
