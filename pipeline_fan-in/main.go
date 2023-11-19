package pipeline_fan_in

import "sync"

func main() {
	fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {

		var wg sync.WaitGroup
		muxStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				select {
				case <-done:
					return
				case muxStream <- v:
				}
			}
		}

		// Select from all the channels, Add all fan-out stages to wait group,
		//so you can close the outgoing channel when all fan-out stages exits
		for _, ch := range channels {
			go multiplex(ch)
		}
		//EN
		//Go's defer statement schedules a function call to be run
		//immediately before the function executing the defer returns

		//TR
		//Go'nun defer ifadesi, defer işlevini yürüten işlev geri dönmeden
		//hemen önce bir işlev çağrısının çaliştırılmasını planlar

		// Wait for all the reads to complete
		go func() {
			wg.Wait()
			close(muxStream)
		}()

		return muxStream
	}
}
