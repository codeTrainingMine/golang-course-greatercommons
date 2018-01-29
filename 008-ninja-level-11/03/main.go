package main

import (
	"fmt"
	"log"
)

//Create a struct “customErr” which implements the builtin.error interface.
//	Create a func “foo” that has a value of type error as a parameter.
//		Create a value of type “customErr” and pass it into “foo”.

type customErr struct {
	first string
	err error
}

func (c customErr) Error() string  {
	return fmt.Sprintf("custom err: %v", c.err)
}

func main() {
	c := customErr{
		first: "James",
		err: fmt.Errorf("divide by 0 error"),
	}
	foo(c)
}

func foo(err error)  {
	log.Fatalln(err)
}