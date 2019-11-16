package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(6)
	for i := 0; i < 6; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("T1 ", i)
		}()
	}
	wg.Wait()
	wg.Add(6)
	for i := 0; i < 6; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("T2 ", i)
		}()
	}
	wg.Wait()
}
