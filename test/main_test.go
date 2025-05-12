package test

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{5, 5}, 10},
		{[]int{5, 3}, 8},
	}

	for _, test := range tests {
		sum := Add(test.data[0], test.data[1])
		if sum != test.answer {
			t.Error("Expected", test.answer, "Got", sum)
		}
	}
}

func ExampleAdd() {
	fmt.Println(Add(2, 3))
	// Output:
	// 7
}

func BenchmarkAdd(b *testing.B) {
	for b.Loop() {
		Add(5, 5)
	}
}
