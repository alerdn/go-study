package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func add[T int | float64](a, b T) T {
	return a + b
}

// Pode ser int ou float64
type SimpleAddable interface {
	int | float64
}

func addGeneric[T SimpleAddable](a, b T) T {
	return a + b
}

// Pode ser int ou float64 além de qualquer alias
type BetterAddable interface {
	~int | ~float64
}

func addB[T BetterAddable](a, b T) T {
	return a + b
}

// Constraints são tipos que englobam outros tipos
type ConstraintAddable interface {
	constraints.Integer | constraints.Float
}

func addC[T ConstraintAddable](a, b T) T {
	return a + b
}

type Number int

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.5, 2.5))

	fmt.Println(addGeneric(1, 2.5))

	var n Number = 5
	fmt.Println(addB(n, n))

	fmt.Println(addC(1.1, -2))
}
