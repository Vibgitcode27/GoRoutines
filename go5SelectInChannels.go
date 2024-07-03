package main

import (
	"fmt"
)

func SelectInChannels() {
	candidate1, candidate2 := make(chan string), make(chan string)

	go electPresident(candidate1, "Candidate 1")
	go electPresident(candidate2, "Candidate 2")

	select {
	case candidate1Message := <-candidate1:
		fmt.Println(candidate1Message)
	case candidate2Message := <-candidate2:
		fmt.Println(candidate2Message)
	}

	// fmt.Println(<-candidate1)
	// fmt.Println(<-candidate2)

	almostRand()
}

func electPresident(candidate chan string, message string) {
	candidate <- message
}

func almostRand() {
	candidate1 := make(chan string)
	close(candidate1)
	candidate2 := make(chan string)
	close(candidate2)

	var can1 int = 0
	var can2 int = 0

	for i := 0; i < 1000; i++ {
		select {
		case <-candidate1:
			can1++
		case <-candidate2:
			can2++
		}
	}

	fmt.Println(can1, " , ", can2)
}
