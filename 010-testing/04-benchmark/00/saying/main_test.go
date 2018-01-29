package saying

import (
	"testing"
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/009-testing/04-benchmark/00-live/saying"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome my dear James" {
		t.Error("Got", s, "Expected", "Welcome my dear James")
	}
}

func ExampleGreet() {
	fmt.Println(saying.Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}