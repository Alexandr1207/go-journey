package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped:", name, ctx.Err())
			return
		case <-ticker.C:
			fmt.Println("Working:", name)
		}
	}
}

func work() string {
	return "FINISHED!"
}

func Run(ctx context.Context, work func() string) (string, error) {
	ch := make(chan string, 1)
	go func() {
		ch <- work()
	}()

	select {
	case v := <-ch:
		return v, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	for _, name := range []string{"worker1", "worker2"} {
		wg.Add(1)
		go Worker(ctx, name, &wg)
	}

	time.Sleep(1 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("main: done") // context canceled a tree of operations, and done cancelled only one operation

	slow := func() string { time.Sleep(time.Second); return "done" }
	fast := func() string { time.Sleep(100 * time.Millisecond); return "done" }

	ctx1, c1 := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer c1()
	fmt.Println(Run(ctx1, slow)) //  "" context deadline exceeded

	ctx2, c2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer c2()
	fmt.Println(Run(ctx2, fast)) //  done <nil>

	parent, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()
	child, cancelChild := context.WithTimeout(parent, 5*time.Second)
	defer cancelChild()

	done := make(chan struct{})
	go func() {
		defer close(done)
		<-child.Done()
		fmt.Println("child err:", child.Err()) // child err: context canceled
	}()

	time.Sleep(300 * time.Millisecond)
	cancelParent()

	<-done
	fmt.Println("main: done")
}
