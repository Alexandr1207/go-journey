package main

import (
	"fmt"
	"sync"
)

// worker passes lock by value: sync.WaitGroup contains sync.noCopy
func worker(wg sync.WaitGroup) { defer wg.Done() }

func ParallelSum(nums []int, parts int) int {
	var wg sync.WaitGroup
	chunks := make([][]int, 0)
	chunkSize := (len(nums) + parts - 1) / parts
	for i := 0; i < len(nums); i += chunkSize {
		end := i + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		chunks = append(chunks, nums[i:end])
	}

	results := make([]int, len(chunks))
	for i, chunk := range chunks {
		wg.Add(1)
		go func(i int, c []int) {
			defer wg.Done()
			s := 0
			for _, v := range c {
				s += v
			}
			results[i] = s
		}(i, chunk)
	}
	wg.Wait()
	sum := 0
	for _, n := range results {
		sum += n
	}

	return sum
}

type Account struct{ balance int }

func (a *Account) Withdraw(n int) bool {
	if a.balance >= n { // проверили
		a.balance -= n // списали
		return true
	}
	return false
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello from goroutine")
	}()
	fmt.Println("main finished")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("goroutine", n)
		}(i)
	}

	counter := 0
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	// After 5 runs:
	// 9654, 9753, 9590, 9481, 9570
	// WaitGroup waiting when all goroutines complete, but data race dont give 10000

	//main.go:9:16: worker passes lock by value: sync.WaitGroup contains sync.noCopy
	//main.go:42:9: call of worker copies lock value: sync.WaitGroup contains sync.noCopy
	// wg.Add(1)
	// worker(wg)

	// if Add inside goroutine, wait can complete earlier
	// cnt := 0
	// go func() {
	// 	defer wg.Done()
	// 	wg.Add(1)
	// 	cnt++
	// }()
	// fmt.Println(cnt) // 0

	nums := make([]int, 1000000)
	fmt.Println(ParallelSum(nums, 50))

	acc := Account{100}
	cTrue := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := acc.Withdraw(50)
			if res {
				cTrue++
			}
		}()
	}
	wg.Wait()
	fmt.Println(cTrue)       // 2
	fmt.Println(acc.balance) // 0
}
