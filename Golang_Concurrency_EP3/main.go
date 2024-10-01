// buffered channel
package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "first message"
	channel <- "second message"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
