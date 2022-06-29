package main

import (
	"fmt"
)

func test(ch chan string) {
	ch <- "test"
}

func main() {

	ch := make(chan string)

	go test(ch)

	v := <-ch
	fmt.Println(v)
}

// basics
// syntax <- ->
// deadlock
// basic usage
