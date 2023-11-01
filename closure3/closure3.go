package closure3

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(sasalutation string) {
			defer wg.Done()
			fmt.Println(sasalutation)
		}(salutation)
	}
	wg.Wait()
}

