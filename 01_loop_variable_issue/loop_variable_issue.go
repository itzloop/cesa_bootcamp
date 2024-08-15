package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	var i int
	for {
		if i == 10 {
			break
		}

		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()

		i++
	}

	wg.Wait()

	fmt.Println("program exiting")
}
