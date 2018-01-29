package main

import (
	"strings"
	"fmt"
	"golang-course-greatercommons/010-testing/04-benchmark/01/mystr"
)

const s = `
Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the “go test” command, which automates execution of any function of the form

func TestXxx(*testing.T)
where Xxx can be any alphanumeric string (but the first letter must not be in [a-z]) and serves to identify the test routine.

Within these functions, use the Error, Fail or related methods to signal failure.

To write a new test suite, create a file whose name ends _test.go that contains the TestXxx functions as described here. Put the file in the same package as the one being tested. The file will be excluded from regular package builds but will be included when the “go test” command is run. For more detail, run “go help test” and “go help testflag”.

Tests and benchmarks may be skipped if not applicable with a call to the Skip method of *T and *B:
`

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}
