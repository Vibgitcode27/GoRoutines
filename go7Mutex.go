package main

import (
	"fmt"
	"sync"
	"time"
)

func mutex() {
	lock := sync.Mutex{}
	iteration := 10000
	count := 0
	for i := 0; i < iteration; i++ {
		go increment(&count, &lock)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(count)
}

func increment(count *int, lock *sync.Mutex) {
	lock.Lock()
	*count++ // this statement is not atomic
	// temp := *count
	// temp++
	// *count = temp
	lock.Unlock()
}

// there are read and write locks in go. Where open for all for
