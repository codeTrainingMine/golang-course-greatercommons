package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	n := 2
	s := oddOrEven(n)
	fmt.Println("the number", n, "is", s)
}

func oddOrEven(n int) string {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

	if n %2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}