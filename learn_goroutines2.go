package main

import (
	"fmt"
	"sync"
)

func runMe2() {
	fmt.Println("Hello from a goroutine")
}

func main() {
     var wg sync.WaitGroup
    wg.Add(1)
    go func() {
	   runMe2()
	   wg.Done()
    }()

	wg.Wait()
}

	
