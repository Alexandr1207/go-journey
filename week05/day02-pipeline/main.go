package main

import (
	"fmt"
)

func CommaOk(nums ...int) {
	ch := make(chan int, 3)
	go func() {
		defer close(ch)
		for _, v := range nums {
			ch <- v
		}
	}()

	for i := 0; i < 4; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}
	// 0 true
	// 1 true
	// 0 false
	// 0 false
	// After 2 iterations, channel is closed.
}

func SentChanel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	ch := make(chan int, 2)
	close(ch)
	ch <- 3
}

func ClosedChannel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	ch := make(chan int, 2)
	close(ch)
	close(ch)
}

func Produce(n int, out chan<- int) {
	defer close(out)
	for i := 1; i <= n; i++ {
		out <- i
	}
}

func Consume(in <-chan int) int {
	s := 0
	for v := range in {
		s += v
	}
	return s
}

func Generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func Filter(in <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if pred(v) {
				out <- v
			}
		}
	}()

	return out
}

func pred(n int) bool {
	return n%2 == 0
}

func main() {
	CommaOk(1, 2)

	SentChanel()    // Recovered: send on closed channel
	ClosedChannel() // Recovered: close of closed channel

	n := 100
	out := make(chan int, n)
	go Produce(100, out)
	fmt.Println(Consume(out)) // 5050

	for v := range Filter(Square(Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)), pred) {
		fmt.Println(v)
	}
}
