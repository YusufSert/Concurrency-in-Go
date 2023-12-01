package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.AfterFunc(2*time.Second, func() {
			fmt.Println("received", <-ch)
		})
	}()

loop:
	for {
		select {
		case ch <- 1:
			break loop
		case <-time.After(3 * time.Second):
			fmt.Println("cancelled")
			break loop
		}
	}
}
