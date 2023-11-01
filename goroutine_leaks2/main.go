package main

import (
	"fmt"
	"time"
)

func main() {

	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} { //1
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done: //2
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() { //3
		// Cancel the operation after 1 second
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated //4
	fmt.Println("Done.")
}

//1 Here we pass the done channel to the doWork function. As a convention, this
//channel is the first parameter.

//2 On this line we see the ubiquitous for-select pattern in use. One of our case statements
//is checking whether our done channel has been signaled. If it has, we
//return from the goroutine.

//3 Here we create another goroutine that will cancel the goroutine spawned in
//doWork if more than one second passes

//4 This is where we join the goroutine spawned from doWork with the main goroutine.
