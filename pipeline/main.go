package main

import "fmt"

func main() {
	for n := range sq(gen(2, 3)) {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, n := range nums {
			c <- n
		}
		close(c)
	}()

	return c
}

func sq(in <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for n := range in {
			c <- n * n
		}
		close(c)
	}()

	return c
}
