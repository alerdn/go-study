package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	incrementor(ctx)
	c := incrementor(ctx)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Done")
}

func incrementor(ctx context.Context) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		i := 0

		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			default:
				c <- i
				i++
			}

			time.Sleep(1 * time.Second)
		}
	}()

	return c
}
