package main

import (
	"fmt"
	"sync"
)


func main () {
var count int
increment := func() {
	fmt.Println("increment")
	count++}
decrement := func() {
		fmt.Println("decrement")
		count--}

var once sync.Once

once.Do(increment)
once.Do(decrement)
fmt.Printf("Count :%d\n", count)
}

