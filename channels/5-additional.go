package main

import "fmt"

func goroutineTest(c chan<- string) {
	//c <- "sd"
	//<-c
}

func main() {
	ch := make(chan string)

	go goroutineTest(ch)

	fmt.Println(<-ch)
}

// Using struct in channel.
// Channel Directions.
