package main

import (
	"fmt"
)

func main() {
	producer := func() <-chan int {
		ownChan := make(chan int, 4)
		go func() {
			defer close(ownChan)
			for i := 0; i < 5; i++ {
				ownChan <- i
			}
			fmt.Printf("Done send. \n")
		}()
		return ownChan
	}
	consumer := func(ch <-chan int) {
		for result := range ch {
			fmt.Printf("Received %d. \n", result)
		}
		fmt.Printf("Done receieve")
	}
	results := producer()
	consumer(results)
}
