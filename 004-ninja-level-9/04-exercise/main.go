package main

import (
	"sync"
	"runtime"
	"fmt"
)

var counter int
var wg sync.WaitGroup
var m sync.Mutex

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
	m.Lock()
	v := counter
	v++
	//runtime.Gosched()
	counter = v
	fmt.Println("counter", counter)
	m.Unlock()
	wg.Done()
}