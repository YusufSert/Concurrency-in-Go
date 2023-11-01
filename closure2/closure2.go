package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation) // 1
		}()
	}
	wg.Wait()
}

// 1 Here we reference the loop variable salutation
// created by ranging over a string slice.
