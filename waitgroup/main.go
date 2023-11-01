package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(time.Second)
	}()

	wg.Wait()
	fmt.Println("All goroutines completes")

}