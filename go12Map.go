package main

import (
	"fmt"
	"sync"
)

func Map() {
	regularMap := make(map[int]interface{}) // regular map do not support concurrent operations operations
	syncMap := sync.Map{}

	regularMap[0] = 0
	regularMap[0] = 0
	regularMap[0] = 0
	syncMap.Store(0, 0)
	syncMap.Store(1, 1)
	syncMap.Store(2, 2)

	regularValue, regularOK := regularMap[0]
	fmt.Println(regularValue, regularOK)
	syncValue, syncOK := syncMap.Load(0)
	fmt.Println(syncValue, syncOK)
}
