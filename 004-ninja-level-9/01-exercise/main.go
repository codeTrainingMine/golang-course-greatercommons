package main

import (
	"sync"
	"fmt"
	"runtime"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("fired up main")
	fmt.Println("goroutines: ", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	fmt.Println("goroutines: ", runtime.NumGoroutine())

	wg.Add(1)
	go bar()
	fmt.Println("goroutines: ", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("main is shutting down")
	fmt.Println("goroutines: ", runtime.NumGoroutine())
}

func foo()  {
	fmt.Println("foo is running")
	wg.Done()
}

func bar()  {
	fmt.Println("bar is running")
	wg.Done()
}