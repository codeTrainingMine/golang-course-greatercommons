package main

import "testing"

func TestOddOrEven1(t *testing.T)  {
	if oddOrEven(2) != "even" {
		t.Fail()
	}
}

func TestOddOrEven2(t *testing.T)  {
	if oddOrEven(3) != "odd" {
		t.Fail()
	}
}

func BenchmarkOddOrEven1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oddOrEven(2)
	}
}