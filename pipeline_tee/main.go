package main

import (
	"fmt"
	"time"
)

func main() {
	tee := func(done, in <-chan interface{}) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range orDone(done, in) {
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	in := make(chan interface{})
	go func() {
		for i := range []int{1, 2, 3, 4} {
			in <- i
		}
	}()
	done := make(chan interface{})
	defer close(done)
	o1, o2 := tee(done, in)

	go func() {
		for v := range o1 {
			fmt.Println(v)
		}
	}()
	go func() {
		for v := range o2 {
			fmt.Println(v)
		}
	}()
	time.Sleep(1 * time.Second)

}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
	loop:
		for {
			select {
			case <-done:
				break loop
			case v, ok := <-c:
				if ok == false {
					break loop
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
