package main

import (
	"fmt"
	"sync"
)

func main() {
	var onceA, onceB sync.Once

	var initB func()

	initA := func() {
		fmt.Println("init a")
		onceB.Do(initB)
		fmt.Println("init a exit")
	}
	initB = func() {
		fmt.Println("init b")
		onceA.Do(initA)
		fmt.Println("init b exit")
	}
	onceA.Do(initA)
}

