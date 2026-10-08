package main

import (
	"fmt"
	"runtime"
	"time"
)

func sayHello() {
	fmt.Println("hello from goroutine")
}

func main() {
	go sayHello()
	fmt.Println("main finished")
	time.Sleep(100 * time.Millisecond) // "hello from goroutine" only if time sleep

	for i := 0; i < 5; i++ {
		go func() { fmt.Println("goroutine", i) }() // Every time order is different
	}
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("goroutine", n) // in new version every goroutine has own i variable
		}(i)
	}
	time.Sleep(time.Second)

	// Goroutines is light, then it can complete a lot of times before main ends.
	counter := 0
	fmt.Println(runtime.NumGoroutine()) // 1
	for i := 0; i < 10000; i++ {
		go func() { counter++ }()
	}
	fmt.Println(counter)                // 9902
	fmt.Println(runtime.NumGoroutine()) // 1
	fmt.Println(runtime.NumCPU())       // 8

	// after go run -race main.go:
	// 	1
	// ==================
	// WARNING: DATA RACE
	// Read at 0x00c00000c108 by goroutine 9:
	//   main.main.func1()
	//       D:/Go_projects/Golang_learn/week04/day01-goroutines/main.go:33 +0x2e

	// Previous write at 0x00c00000c108 by goroutine 8:
	//   main.main.func1()
	//       D:/Go_projects/Golang_learn/week04/day01-goroutines/main.go:33 +0x44

	// Goroutine 9 (running) created at:
	//   main.main()
	//       D:/Go_projects/Golang_learn/week04/day01-goroutines/main.go:33 +0xc4

	// Goroutine 8 (finished) created at:
	//   main.main()
	//       D:/Go_projects/Golang_learn/week04/day01-goroutines/main.go:33 +0xc4
	// ==================
}
