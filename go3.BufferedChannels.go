package main

import (
	"fmt"
	"time"
)

func bufferedChannels() {
	channel := make(chan string, 2)
	channel <- "Vibhor"
	channel <- "Vibhor"

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Vibhor"
	}()
	count := 0
	for count != 3 {
		fmt.Println(<-channel)
		count++
	}

}
