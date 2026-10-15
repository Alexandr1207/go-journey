package main

import (
	"fmt"
	"time"
)

// Task 1.
func HelloChannel(c chan string) {
	c <- "Hello from goroutine"
}

// Task 2.
func SumParallel(nums []int) int {
	mid := len(nums) / 2

	ch := make(chan int)

	go func() {
		s := 0
		for _, i := range nums[:mid] {
			s += i
		}
		ch <- s
	}()

	go func() {
		s := 0
		for _, i := range nums[mid:] {
			s += i
		}
		ch <- s
	}()

	a := <-ch
	b := <-ch

	return a + b
}

func BuferChan(a, b, c, d int) {
	ch := make(chan int, 3)
	ch <- a
	ch <- b
	ch <- c
	ch <- d
}

func main() {
	ch := make(chan string)
	go HelloChannel(ch)
	res := <-ch
	fmt.Println(res) // Hello from goroutine

	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(SumParallel(nums)) // 21

	// BuferChan(1, 2, 3, 4) // fatal error: all goroutines are asleep - deadlock! Channel one without help

	c := make(chan string)
	go func() {
		fmt.Println("A: before send")
		c <- "data"
		fmt.Println("A: after send")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main: before receive")
	msg := <-c
	fmt.Println("main: got", msg)
	time.Sleep(100 * time.Millisecond)
}
