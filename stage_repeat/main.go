package main

import (
	"fmt"
	"time"
)

func main() {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valuesStream := make(chan interface{})
		go func() {
			defer close(valuesStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valuesStream <- v:
					}
				}
			}
		}()
		return valuesStream
	}

	done := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Millisecond)
		defer close(done)
	}()

	for v := range repeat(done, []int{1, 2, 3, 4, 5}) {
		fmt.Println(v)
	}
}
