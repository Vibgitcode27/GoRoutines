package main

import (
	"fmt"
	"sync"
)

func waitGroup() {
	var async sync.WaitGroup
	async.Add(4)
	nina := []string{"Tommy", "Chopper", "Nina", "Kalasenikov"}
	for i := 0; i < 4; i++ {
		go attackTommy(nina[i], &async)
	}
	async.Wait()

	var beeper sync.WaitGroup
	beeper.Add(1)
	beeper.Done()
	beeper.Done()
	beeper.Wait()
}

func attackTommy(tommy string, async *sync.WaitGroup) {
	fmt.Println("Attack on :- ", tommy)
	async.Done()
}
