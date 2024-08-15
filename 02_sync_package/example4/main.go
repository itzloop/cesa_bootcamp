package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	var (
		wg   = sync.WaitGroup{}
		nums = map[int]bool{}
	)

	n := flag.Int("n", 10, "number of interations")
	flag.Parse()

	for i := 0; i < *n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nums[i] = true
		}()
	}

	wg.Wait()

	fmt.Println(len(nums))
	fmt.Println("program exiting")
}
