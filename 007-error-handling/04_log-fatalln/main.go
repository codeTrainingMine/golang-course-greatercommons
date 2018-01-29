package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	defer foo()

	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("err happened", err)
	}
}

func foo()  {
	fmt.Print("when os.exit is called, deferred functions dont run")
}