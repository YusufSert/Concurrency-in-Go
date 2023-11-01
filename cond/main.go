package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay) // pre work: connection, anything than takes time
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	c.L.Lock()
	for i := 0; i < 10; i++ {
		//c.L.Lock() orginal lock form book
		for len(queue) == 2 {
			fmt.Println("wait")
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second)
		// c.L.UnLock() orginal Unlock from book
		//time.Sleep(1*time.Second) not from book
	}
	c.L.Unlock()
}