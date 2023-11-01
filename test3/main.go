package main

import (
	"sync"
)

func main() {
	var mu sync.Mutex
	Do := func(f func()) {
		mu.Lock()
		f()
		mu.Unlock()
	}

	f := func() {
		Do(func() {})
	}

	Do(func() {
		Do(f)
	})
}
