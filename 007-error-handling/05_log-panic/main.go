package main

import (
	"os"
	"log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln("err happened", err)
	}
}
