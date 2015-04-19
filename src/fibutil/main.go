// fibutil project main.go
package main

import (
	"fibutil/fib"
	"fmt"
)

func main() {
	nextFibo := fib.Fibo()
	for i := 0; i < 20; i++ {
		fmt.Printf("%d : %d\n", i, nextFibo())
	}

	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case c <- 0: // no statement, no fall through
		case c <- 1:
		}
	}
}
