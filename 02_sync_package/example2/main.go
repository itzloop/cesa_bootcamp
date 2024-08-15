package main

import (
	"flag"
	"fmt"
	"sync"
)


func main() {
    var (
        num = 0
        wg = sync.WaitGroup{}
        mu = sync.Mutex{}
    )

    n := flag.Int("n", 10, "number of interations")
    flag.Parse()
   
    for i := 0; i < *n; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            num++
            mu.Unlock()
        }()
    }

    wg.Wait()

    fmt.Println(num)
    fmt.Println("program exiting")
}
