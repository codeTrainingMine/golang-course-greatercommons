package main

import (
	"os"
	"fmt"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
	}
}
