package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var x int = 10

	//* Switch básico
	switch x {
	case 10:
		fmt.Println("x é 10")
	case 20:
		fmt.Println("x é 20")
	}

	switch x {
	case 10:
		fmt.Println("x é 10")
		fallthrough
	case 20:
		fmt.Println("x é 10, mas caiu por causa do fallthrough")
	}

	//* Select
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		d := time.Duration(rand.Intn(1))
		time.Sleep(d * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		d := time.Duration(rand.Intn(1))
		time.Sleep(d * time.Millisecond)
		ch2 <- 2
	}()

	//* O select vai esperar o primeiro canal que receber dados e vai executar o case correspondente.
	select {
	case <-ch1:
		fmt.Println("Canal 1 recebeu dados primeiro")
	case <-ch2:
		fmt.Println("Canal 2 recebeu dados primeiro")
	}
}
