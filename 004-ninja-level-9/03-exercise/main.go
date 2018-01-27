package main

import (
	"sync"
	"runtime"
	"fmt"
)

var counter int
var wg sync.WaitGroup

func main() {

	routines := 100

	fmt.Println("counter", counter)
	fmt.Println("routines", runtime.NumGoroutine())

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go incrementor()
		fmt.Println("routines", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("counter", counter)
	fmt.Println("routines", runtime.NumGoroutine())
}

func incrementor() {
	v := counter
	v++
	runtime.Gosched()
	counter = v
	wg.Done()
	fmt.Println("counter", counter)
}