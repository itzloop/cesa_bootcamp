package main

import (
	"fmt"
	"sync"
)

var art = `
         ,_---~~~~~----._         
  _,,_,*^____      _____''*g*\"*, 
 / __/ /'     ^.  /      \ ^@q   f 
[  @f | @))    |  | @))   l  0 _/  
 \'/   \~____ / __ \_____/    \   
  |           _l__l_           I   
  }          [______]           I  
  ]            | | |            |  
  ]             ~ ~             |  
  |                            |   
   |                           |  
`

func main() {
    wg := sync.WaitGroup{}

    wg.Add(1)
	go func() {
        defer wg.Done()
		fmt.Println(art)
        panic("nil pointer dereference")
	}()

    wg.Wait()
	fmt.Println("Program exiting")
}
