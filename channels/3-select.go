package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func(c chan string) {
		time.Sleep(300 * time.Millisecond)
		c <- "chan2"
	}(chan2)

	timeout := time.After(200 * time.Millisecond)

	select {
	case msg := <-chan1:
		fmt.Println("from channel 1", msg)

	case <-chan2:
		fmt.Println("from channel 2")

	case v := <-timeout:
		fmt.Println("time exceed", v)
	}
}

// timeout
