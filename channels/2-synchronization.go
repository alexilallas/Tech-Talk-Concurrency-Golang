package main

import (
	"fmt"
)

func collectPhone(data []string, c chan string) {
	for _, v := range data {
		if v == "phone" {
			c <- v
		}
	}
	close(c)
}

func savePhone(c chan string, done chan bool) {
	for v := range c {
		fmt.Println("received", v)
	}

	done <- true
}

func main() {
	data := []string{"phone", "name", "info", "phone", "city", "phone", "age"}
	ch := make(chan string)
	done := make(chan bool)

	go collectPhone(data, ch)
	go savePhone(ch, done)

	<-done

}

// challenge: get and process phone concurrently
