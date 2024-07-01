package main

import (
	"fmt"
	"time"
)

func attack2(target string) {
	time.Sleep(1 * time.Second)
	fmt.Println(target)
}

func channelsImplementation() {
	// Create a channel
	ninja := "Vibhor"

	go attack2(ninja)

	time.Sleep(2 * time.Second)
}
