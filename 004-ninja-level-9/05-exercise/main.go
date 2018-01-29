package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var counter int64
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
	atomic.AddInt64(&counter, 1)
	runtime.Gosched()
	//fmt.Println("counter", counter)
	wg.Done()
}