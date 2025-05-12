// Package main estuda channel
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Canal bidimensional
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go Sender(even, odd, quit)
	go receiver(even, odd, quit, &wg)

	wg.Wait()
}

// Sender Canal unidimensional onde enviamos dados para ele
func Sender(even, odd, quit chan<- int) {
	for i := range 10 {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	quit <- 0
}

// Canal unidimensional onde recebemos dados dele
func receiver(even, odd, quit <-chan int, wg *sync.WaitGroup) {
	for {
		// O canal que receber dado primeiro é escolhido
		select {
		case v := <-even:
			fmt.Printf("%d é par\n", v)
		case v := <-odd:
			fmt.Printf("%d é ímpar\n", v)
		case <-quit:
			fmt.Println("saindo do programa...")
			wg.Done()
			return
		}
	}
}
