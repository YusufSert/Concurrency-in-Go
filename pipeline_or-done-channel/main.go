package pipeline_or_done_channel

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
