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
		fmt.Println("closure-lock")
		defer c.L.Unlock()
		c.Wait()
		res, err := http.Get("https://www.youtube.com/")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.Status)
		fmt.Println("closure-unlock")
		
		wg.Done()
	}


	wg.Add(1)
	go r()

	os.Stdin.Read(make([]byte, 1))
	c.Signal()
	fmt.Println("signal")
	
	wg.Wait()
}