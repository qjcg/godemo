package hello

import (
	"testing"
)

var testCases = []struct {
	Name     string
	Expected string
}{
	{"John", "Howdy John!"},
	{"Jennifer", "Howdy Jennifer!"},
	{"", "Howdy Diego!"},
}

func TestGreet(t *testing.T) {
	for _, tc := range testCases {
		got := Greet(tc.Name)
		if got != tc.Expected {
			t.Fatalf("Expected: %s, Got: %s\n", tc.Expected, got)
		}
	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("")
	}

}

func ExampleGreet() {
	Greet("")
	Greet("John")
}
