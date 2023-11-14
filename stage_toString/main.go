package main

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

/*Example of how to use it:*/
func main() {
	generator := func(done <-chan interface{}) <-chan interface{} {
		genStream := make(chan interface{})

		go func() {
			defer close(genStream)

			for _, s := range []string{"K", "u", "d", "i"} {
				select {
				case <-done:
					return
				case genStream <- s:
				}
			}
		}()
		return genStream
	}

	done := make(chan interface{})
	defer close(done)

	stream := generator(done)
	var message string

	for token := range toString(done, stream) {
		message += token
	}
}
