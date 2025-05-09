package main

import "fmt"

func main() {
	generator := Fibonacci()

	for range 5 {
		fmt.Println("Número gerado: ", generator())
	}

	func() {
		fmt.Println("Sou uma função anônima")
	}()

	calculate(10, 20, add, func(resultado int) {
		fmt.Println("O resultado foi ", resultado)
	})

	calculate(10, 20, sub, func(resultado int) {
		fmt.Println("O resultado foi ", resultado)
	})
}

func Fibonacci() func() int {
	prev := 0
	n := 1
	return func() int {
		i := prev + n
		prev = n
		n = i

		return i
	}
}

func calculate(a, b int, f func(int, int) int, callback func(int)) {
	callback(f(a, b))
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
