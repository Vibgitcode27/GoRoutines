package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ready bool = false

func cond() {
	go gettingReady()
	for !ready {
		println("I am not ready")
	}

	fmt.Println("Now I am ready")
}

func gettingReady() {
	ready = false
	Sleep()
}

func Sleep() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(5)
	time.Sleep(time.Duration(randNum) * time.Second)
	ready = true
}

func condImplementation() {

	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond) // Here we are not passing address of cond because it is already a pointer

	cond.L.Lock() // This is incorrect
	for !ready {
		println("I am not ready")
		cond.Wait() // Wait on the condition
	}

	fmt.Println("Now I am ready")
	cond.L.Unlock()
}

func gettingReadyWithCond(cond *sync.Cond) {
	ready = false
	Sleep()
	cond.Signal() // Signal waiting goroutines after setting ready to true
}
