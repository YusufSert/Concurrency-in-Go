package stage_take

func take(done, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case v := <-valueStream:
				takeStream <- v
			}
		}
	}()
	return takeStream
}
