package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)


func main() {
	c := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	r :=  func() {
		c.L.Lock()
		fmt.Println("func-lock")
		res, err := http.Get("https://www.youtube.com/")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.Status)
		c.L.Unlock()
		fmt.Println("func-unlock")
		c.Signal()
		wg.Done()
	}

	c.L.Lock()
	fmt.Println("main-lock")

	os.Stdin.Read(make([]byte, 1))

	wg.Add(1)
	go r()

	fmt.Println("wait-unlock")	
	c.Wait() // unlock and wait for event
	fmt.Println("wait-lock")	

	c.L.Unlock()
	fmt.Println("main-unluck")
	
	wg.Wait()
}