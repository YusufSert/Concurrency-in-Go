package main

import (
	"fmt"
)

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int, len(integers))
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan interface{}, inputStream <-chan int, multiplier int) <-chan int {
		outputStream := make(chan int)
		go func() {
			defer close(outputStream)
			for i := range inputStream {
				select {
				case <-done:
					return
				case outputStream <- i * multiplier:
				}
			}
		}()
		return outputStream
	}

	add := func(done <-chan interface{}, inputStream <-chan int, additive int) <-chan int {
		outputStream := make(chan int)
		go func() {
			defer close(outputStream)
			for i := range inputStream {
				select {
				case <-done:
					return
				case outputStream <- i + additive:
				}
			}
		}()
		return outputStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
