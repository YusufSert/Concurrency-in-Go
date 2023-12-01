package main

import "fmt"

func bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)

	loop:
		for {
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if ok != true {
					return
				}
				stream = maybeStream
			case <-done:
				break loop
			}
			for val := range orDone(done, stream) {
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func orDone(done, in <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
	loop:
		for {
			select {
			case <-done:
				break loop
			case val, ok := <-in:
				if ok == false {
					break loop
				}
				select {
				case valStream <- val:
				case <-done:
				}

			}
		}
	}()
	return valStream
}

func main() {
	genVal := func() <-chan <-chan interface{} {
		chStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chStream)
			for i := 1; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chStream <- stream
			}
		}()
		return chStream
	}

	for v := range bridge(nil, genVal()) {
		fmt.Printf("%v ", v)
	}
}
