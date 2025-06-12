package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch1 <- "message from channel 1"
	}()
	go func() {
		defer wg.Done()
		ch2 <- "message from channel 2"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case msg1 := <-ch1:
			fmt.Println(msg1)

		}
	}
	wg.Wait()
	fmt.Println("I am done")
}
