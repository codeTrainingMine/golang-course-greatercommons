package main

import "testing"

func TestMySum(t *testing.T)  {
	x := mySum(2, 3 )
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

func TestMultiply(t *testing.T)  {
	x := multiply(2, 3)
	if x != 6 {
		t.Error("Expected", 6, "Got", x)
	}
}
