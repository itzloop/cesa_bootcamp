package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	var (
		wg   = sync.WaitGroup{}
		nums []int // nil
        mu = sync.Mutex{}
	)

	n := flag.Int("n", 10, "number of interations")
	flag.Parse()

	for i := 0; i < *n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
            mu.Lock()
			nums = append(nums, i)
            mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(len(nums))
	fmt.Println("program exiting")
}
