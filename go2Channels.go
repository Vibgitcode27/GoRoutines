package main

import (
	"fmt"
	"time"
)

func attack2(target string, ensure chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println(target)
	ensure <- true
}

func channelsImplementation() {
	// Create a channel
	ninja := "Vibhor"
	channel := make(chan bool)

	go attack2(ninja, channel)
	channel <- false // Will lead to dealock
	fmt.Println("Waiting for the goroutine to finish")
	for !<-channel {
	}
	// Wait for the goroutine to finish
}
