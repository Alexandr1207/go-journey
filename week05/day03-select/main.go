package main

import (
	"errors"
	"fmt"
	"time"
)

func WhoIsFaster() {
	ch := make(chan string, 1)
	ch2 := make(chan string, 1)
	go func(ch chan string) {
		ticker := time.NewTicker(800 * time.Millisecond)
		defer ticker.Stop()

		<-ticker.C
		ch <- "fast"
	}(ch)

	go func(ch chan string) {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()

		<-ticker.C
		ch <- "slow"
	}(ch2)

	select {
	case v := <-ch:
		fmt.Println("получено из ch1:", v)
	case v := <-ch2:
		fmt.Println("получено из ch2:", v)
	}
}

func work() string {
	return "FINISHED!"
}

func WithTimeout(work func() string, limit time.Duration) (string, error) {
	ch := make(chan string, 1)
	go func() { ch <- work() }()

	select {
	case v := <-ch:
		return v, nil
	case <-time.After(limit):
		return "", errors.New("timeout")
	}
}

func TryReceive(ch <-chan int) (int, bool) {
	select {
	case v, ok := <-ch:
		return v, ok
	default:
		return 0, false
	}
}

func TickerAndStop(done <-chan struct{}, finished chan<- struct{}) {
	defer close(finished)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-done:
			fmt.Println("worker: stop")
			return
		}
	}
}

func BreakTrap() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	count := 0
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		}
		count++
		if count > 10 {
			return
		}
	}
}

func main() {
	WhoIsFaster()
	fmt.Println(WithTimeout(work, 500))
	ch := make(chan int, 2)
	close(ch)
	fmt.Println(TryReceive(ch))

	done := make(chan struct{})
	finished := make(chan struct{})
	go TickerAndStop(done, finished)

	time.Sleep(1 * time.Second)
	close(done)
	<-finished
	fmt.Println("main: all done")

	BreakTrap()
}
