package main

import "sync"

func main() {
	// var beeper sync.WaitGroup
	// evilNinjas := []string{"tommy", "johnz", "bobby"}
	// beeper.Add(len(evilNinjas))

	// for _, evilNinja := range evilNinjas {
	// 	go attack(evilNinja, &beeper)

	// }
	// beeper.Wait()
	// fmt.Println("Mission completed")
	var beeper sync.WaitGroup
	beeper.Add(2)
	beeper.Done()

	beeper.Done()

	beeper.Wait()

}

// func attack(evilNinja string, beeper *sync.WaitGroup) {
// 	fmt.Println("Evil Ninja ..: ", evilNinja)
// 	beeper.Done()
// }
