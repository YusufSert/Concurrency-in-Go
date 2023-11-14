package main

import (
	"fmt"
	"math/rand"
)

type Fn func() interface{}

func main() {
	repeatFn := func(done <-chan interface{}, fn Fn) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	done := make(chan interface{})
	defer close(done)

	rando := func() interface{} { return rand.Int() }

	for num := range take(done, repeatFn(done, rando), 10) {
		fmt.Println(num)
	}

}

func take(done, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case v := <-valueStream:
				takeStream <- v
			}
		}
	}()
	return takeStream
}
