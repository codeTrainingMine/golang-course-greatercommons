package main

import (
	"sync"
	"fmt"
)

//write a program that
//launches 10 goroutines
//each goroutine adds 10 numbers to a channel
//pull the numbers off the channel and print them

func main() {
	c := gen()

	for v := range c {
		fmt.Println("final print is printing", v, "from the channel")
	}
}

func gen() <-chan int {
	c := make(chan int)

	numchans := 10

	go func(){
		var wg sync.WaitGroup
		for i := 0; i < numchans; i++ {
			wg.Add(1)
			go func(routineid int){
				for j := 0; j < 10; j++ {
					c <- j
					fmt.Println("routine id", routineid, "is adding", j, "to channel")
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()

	return c
}
