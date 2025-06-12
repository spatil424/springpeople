package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter is %d", counter)

}
