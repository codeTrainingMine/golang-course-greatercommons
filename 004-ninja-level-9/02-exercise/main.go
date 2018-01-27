package main

import "fmt"

//This exercise will reinforce our understanding of method sets:
//
//create a type person struct
//	attach a method speak to type person using a pointer receiver
//	*person
//	create a type human interface
//	to implicitly implement the interface, a human must have the speak method
//	create func “saySomething”
//	have it take in a human as a parameter
//	have it call the speak method
//	show the following in your code
//	you CAN pass a value of type *person into saySomething
//	you CANNOT pass a value of type person into saySomething

type person struct {
	first string
	last string
	age int
}

func (p *person) speak(s string){
	fmt.Println(p.first, "says", s)
}

type human interface{
	speak(s string)
}

func saySomething(h human)  {
	h.speak("I'm saying something")
}

func main() {
	p := person {
		first: "john",
		last: "smith",
		age: 21,
	}

	saySomething(&p)
	saySomething(p)
}