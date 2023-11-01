package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	sync.Mutex
	value int
}

func(v *value) Work() {
	fmt.Println(v.value)
}

func main() {

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.Lock()
		defer v1.Unlock()

		// Here we sleep for a period of time to simulate work(and trigger a deadlock)
		time.Sleep(2*time.Second) 

		v2.Lock()
		defer v2.Unlock()
		fmt.Printf("sum=%v\n", v1.value + v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
