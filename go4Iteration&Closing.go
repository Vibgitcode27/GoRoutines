package main

import (
	"fmt"
	"math/rand"
)

func randamizor(channel chan int) {
	// rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		score := rand.Intn(10)
		channel <- score
	}

	close(channel)
}

func iterationClosing() {
	channel := make(chan int) // We can extract all the vlues from the go routine using a non buffered channel here
	// for i := 0; i < 5; i++ {
	go randamizor(channel)
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-channel)
	// }

	for value := range channel {
		fmt.Println(value)
	}
}
