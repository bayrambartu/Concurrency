package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}
	time.Sleep(time.Second * 2)

}

func attack(target string) {
	fmt.Println("Throwing ninja starts at", target)
	time.Sleep(time.Second)
}
