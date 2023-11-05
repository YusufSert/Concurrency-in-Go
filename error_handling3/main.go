package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{err, resp}
				select {
				case <-done:
					fmt.Println("cancelled")
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	// close(done) doesn't make sense to call close on main goroutine

	errCount := 0
	urls := []string{"a", "https://www.google.com", "b", "c", "d"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				close(done) // better but because only executing on main goroutine otherwise use defer close(chan)
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
	time.Sleep(time.Second)
}
