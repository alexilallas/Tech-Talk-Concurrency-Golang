package main

import (
	"fmt"
	"time"
)

func goroutineTest(count *int) {
	*count++
}

func main() {
	var count int
	go goroutineTest(&count)

	if count == 0 {
		time.Sleep(time.Millisecond)
		fmt.Println("Zero", count)
	}

	fmt.Println("main")
}

// 2 ways to start one goroutine
// main goroutine
// race condition
