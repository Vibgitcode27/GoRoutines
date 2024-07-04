package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	pID  int
	data string
}

var (
	queue   []Task
	mux     sync.Mutex
	condVar sync.Cond
)

func producer(pId int, task string) {
	mux.Lock()
	queue = append(queue, Task{pID: 1, data: task + string(pId)})
	fmt.Println("Producer added a task")
	condVar.Signal()
	mux.Unlock()
}

func consumer() {
	mux.Lock()
	for len(queue) == 0 {
		condVar.Wait()
	}
	var task Task = queue[0]
	mux.Unlock()
	fmt.Println("Consumer consumed a task : ", task.data, "with pID: ", task.pID)
	queue = queue[1:]
}

func SandB() {
	condVar.L = &mux
	go producer(1, "Task")
	go producer(2, "Task")
	go consumer()
	go consumer()
	go consumer()

	go func() {
		// Simulate adding more tasks later
		time.Sleep(1 * time.Second)
		producer(3, "Task 3 data")
	}()

	time.Sleep(5 * time.Second)
}
