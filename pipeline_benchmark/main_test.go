package pipeline_benchmark

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valuesStream := make(chan interface{})
	go func() {
		defer close(valuesStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valuesStream <- v:
				}
			}
		}
	}()
	return valuesStream
}
func take(
	done <-chan interface{},
	valueStream <-chan string,
	num int,
) <-chan string {
	takeStream := make(chan string)
	go func() {
		defer close(takeStream)
		for i := num; i > 0 || i == -1; {
			if i != -1 {
				i--
			}
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func toString(done, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)

		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}
