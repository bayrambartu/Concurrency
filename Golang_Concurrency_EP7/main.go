package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int // default -> 0
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func main() {

	// fmt.Println(count) // -> 0
	// iterations := 1000
	// for i := 0; i < iterations; i++ {
	// 	go increment()
	// }
	// time.Sleep(1 * time.Second)
	// fmt.Println("Resulted count is : ", count)
	readAndWrite()

}

/* func basics() {
	fmt.Println(count) // -> 0
	iterations := 1000
	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is : ", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
} */

func readAndWrite() {

	go write()
	go read()
	go read()

	go write()
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}
func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Reading unlocking")
}
func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlocking")
}
