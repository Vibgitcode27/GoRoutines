package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func Once() {
	foundAir1 := false
	async := sync.WaitGroup{}
	once := sync.Once{}
	async.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			if findAir1() {
				once.Do(func() { markAir1Found(&foundAir1, i) })
			}
			async.Done()
		}()
	}

	async.Wait()

	checkAir1found(&foundAir1)
}

func checkAir1found(foundAir1 *bool) {
	if *foundAir1 {
		fmt.Println("Air1 found")
	} else {
		fmt.Println("Air1 not found")
	}
}

func markAir1Found(foundAir1 *bool, i int) {
	*foundAir1 = true
	fmt.Println("Air1 found by ", i)
}

func findAir1() bool {
	return 0 == rand.Intn(10)
}
