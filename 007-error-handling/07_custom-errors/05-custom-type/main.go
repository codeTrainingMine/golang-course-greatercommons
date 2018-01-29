package main

import (
	"log"
	"fmt"
)

type norgateMathError struct {
	lat string
	long string
	err error
}

func (n norgateMathError) Error() string {
	fmt.Println("in Error()")
	return fmt.Sprintf("a norgate math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		fmt.Println("before the log")
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	defer fmt.Println("in sqrt's defer")
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number")
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}