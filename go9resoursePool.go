package main

import (
	"fmt"
	"sync"
)

func resoursePool() {
	var numMemPools int
	memPool := sync.Pool{
		New: func() interface{} {
			numMemPools++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	const numOfGoWorkers = 1024 * 1024

	var wg sync.WaitGroup
	wg.Add(numOfGoWorkers)
	for i := 0; i < numOfGoWorkers; i++ {
		go func() {
			mem := memPool.Get().(*[]byte)
			//                     ^Type assertion basically we are passing type to interface{}
			fmt.Sprintln("Taking some time here")
			memPool.Put(mem)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Number of memory pools: ", numMemPools)
}
