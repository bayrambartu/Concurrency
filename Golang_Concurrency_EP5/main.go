package main

import (
	"fmt"
)

func main() {
	n1, n2 := make(chan string), make(chan string)

	go ce(n1, "n1")
	go ce(n2, "n2")

	select {
	case message := <-n1:
		fmt.Println(message)
	case message := <-n2:
		fmt.Println(message)

	default:
		fmt.Println("Neither")
	}
	rf()
}

func ce(n chan string, message string) {
	n <- message
}

func rf() {
	n1 := make(chan interface{})
	close(n1)
	n2 := make(chan interface{})
	close(n2)

	var n1Count, n2Count int
	for i := 0; i < 1000; i++ {

		select {
		case <-n1:
			n1Count++

		case <-n2:
			n2Count++
		}
	}

	fmt.Printf("n1Count: %d, n2Count: %d\n", n1Count, n2Count)

}
