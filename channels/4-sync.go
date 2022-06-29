package main

import (
	"fmt"
	"sync"
)

func goroutine(wg *sync.WaitGroup) {
	fmt.Println("goroutine!")
	wg.Done()
}

func main() {
	qnt := 10

	var wg sync.WaitGroup
	wg.Add(qnt)

	for i := 0; i < qnt; i++ {
		go goroutine(&wg)
	}

	wg.Wait()
	fmt.Println("finish")

}

// challenge: wait for all goroutines finish
