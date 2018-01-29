package main

import (
	"fmt"
	"golang-course-greatercommons/009-ninja-level-12/01/dog"
)

//Create a dog package. The dog package should have an exported func “Years” which takes human years and turns them into dog years
// (1 human year = 7 dog years). Document  your code with comments. Use this code in func main.
//
//run your program and make sure it works
//run a local server with godoc and look at your documentation.

func main() {
	n := 3
	fmt.Println(n, "human years equals to", dog.Years(3), "dog years.")
}
