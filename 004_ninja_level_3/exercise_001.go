package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

 */

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	exercise5()
}

func exercise1() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	for i := 65; i < 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func exercise3() {
	bd := 1989

	for {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

func exercise4() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	rangeRandomNumber := 100
	randomNumber := r1.Intn(rangeRandomNumber)
	fmt.Println("Random number generated", randomNumber)
	if randomNumber%rangeRandomNumber > rangeRandomNumber/2 {
		fmt.Println("Second half of the random number range")
	} else {
		fmt.Println("First half of the random number range")
	}
}

func exercise5() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}

func exercise6() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
