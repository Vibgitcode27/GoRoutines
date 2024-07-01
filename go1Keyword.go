package main

import (
	"fmt"
	"time"
)

func attack(goNinja string) {
	fmt.Println(goNinja)
	time.Sleep(1 * time.Second)
}

func goKeywordImplementation() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	goNinjas := []string{"Vibhor", "Varun", "Adul", "Rahul", "Rohit"}

	for _, goNinja := range goNinjas {
		go attack(goNinja)
	}
	time.Sleep(2 * time.Second)
}
