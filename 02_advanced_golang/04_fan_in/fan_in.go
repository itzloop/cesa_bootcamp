package main

import (
	"flag"
	"fmt"
)

func main() {
	// define some command line arguments
	n := flag.Int("n", 10, "number of elements sent on ch1 and ch2")
	flag.Parse()

	// make channel 1 and channel 2
	ch1, ch2 := make(chan string), make(chan string)
	ch3 := runFanIn(ch1, ch2)

	// send some data on ch1 and ch2
	go func() {
		for i := 0; i < *n; i++ {
			ch1 <- fmt.Sprintf("data %d sent on channel 1", i)
			ch2 <- fmt.Sprintf("data %d sent on channel 2", i)
		}

		close(ch1)
		close(ch2)
	}()

	for result := range ch3 {
		fmt.Println(result)
	}

	fmt.Println("program exiting")
}

func runFanIn(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)

	close1, close2 := false, false
	go func() {

		defer close(ch)
		for {
			if close1 && close2 {
				fmt.Println("both channels are closed")
				return
			}

			select {
			case s, ok := <-ch1:
				if !ok {
					close1 = true
					continue
				}

				ch <- s
			case s, ok := <-ch2:
				if !ok {
					close2 = true
					continue
				}
				ch <- s
			}
		}
	}()
	return ch
}
