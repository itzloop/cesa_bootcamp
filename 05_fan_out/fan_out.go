package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	// define some command line arguments
	n := flag.Int("n", 10, "number of elements sent on ch1 and ch2")
	sleep := flag.Bool("sleep", false, "add a one second sleep for each data received")
	flag.Parse()

	// run fan out and get a single channel to read from
	ch := runFanOut(*n)

	// create a wait group to wait for two go-routines
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		// read the result from the channel
		for result := range ch {
			fmt.Printf("data %d read from loop 1\n", result)

			if *sleep {
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		// read the result from the channel
		for result := range ch {
			fmt.Printf("data %d read from loop 2\n", result)

			if *sleep {
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()
	fmt.Println("program exiting")
}

func runFanOut(n int) <-chan int {
	ch3 := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			ch3 <- i
		}

		close(ch3)
	}()

	return ch3
}
