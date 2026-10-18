package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Id    int
	Value int
}

type Result struct {
	JobId int
	Out   int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- Result{JobId: j.Id, Out: j.Value * j.Value}
		time.Sleep(200 * time.Millisecond)
	}
}

func SquareAll(nums []int, workers int) []int {
	jobs := make(chan Job, len(nums))
	results := make(chan Result, len(nums))
	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func() {
		for i, v := range nums {
			jobs <- Job{Id: i, Value: v}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results) // Wait for waitgroup of worker and close results channel
	}()

	out := make([]int, len(nums))
	for r := range results {
		out[r.JobId] = r.Out
	}
	return out
}

type JobWithErr struct {
	Id    int
	Value int
}

type ResultWithErr struct {
	JobId int
	Out   int
	Err   error
}

func workerWithErr(id int, jobs <-chan JobWithErr, results chan<- ResultWithErr, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		func() {
			defer func() {
				if r := recover(); r != nil {
					results <- ResultWithErr{
						JobId: j.Id,
						Out:   0,
						Err:   fmt.Errorf("worker panic: %v", r),
					}
				}
			}()

			if j.Value == 0 {
				results <- ResultWithErr{
					JobId: j.Id,
					Out:   0,
					Err:   errors.New("division by zero"),
				}
				return
			}

			out := 100 / j.Value
			results <- ResultWithErr{JobId: j.Id, Out: out, Err: nil}
		}()

		time.Sleep(200 * time.Millisecond)
	}
}

func DivideAll(nums []int, workers int) ([]int, []error, int, int) {
	jobs := make(chan JobWithErr, len(nums))
	results := make(chan ResultWithErr, len(nums))
	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go workerWithErr(i, jobs, results, &wg)
	}

	go func() {
		for i, v := range nums {
			jobs <- JobWithErr{Id: i, Value: v}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	out := make([]int, len(nums))
	errs := make([]error, len(nums))
	successCount := 0
	errorCount := 0

	for r := range results {
		if r.Err != nil {
			errs[r.JobId] = r.Err
			errorCount++
		} else {
			out[r.JobId] = r.Out
			successCount++
		}
	}

	return out, errs, successCount, errorCount
}

func RunWithErrors() {
	nums := []int{5, 0, 10, 2, 0, 4, 25, 0, 50, 1}

	start := time.Now()
	out, errs, success, failed := DivideAll(nums, 3)
	elapsed := time.Since(start)

	fmt.Println("Результаты (100 / n):")
	for i, v := range nums {
		if errs[i] != nil {
			fmt.Printf("  [%d] 100 / %d = ERROR: %v\n", i, v, errs[i])
		} else {
			fmt.Printf("  [%d] 100 / %d = %d\n", i, v, out[i])
		}
	}
	fmt.Printf("\nУспешных: %d, С ошибкой: %d\n", success, failed)
	fmt.Printf("Время выполнения: %v\n", elapsed)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	start := time.Now()
	fmt.Println(SquareAll(nums, 3)) // [1 4 9 16 25 36 49 64 81 100 121 144]
	end := time.Since(start)
	fmt.Println(end) // 802.6761ms

	start = time.Now()
	fmt.Println(SquareAll(nums, 6)) // [1 4 9 16 25 36 49 64 81 100 121 144]
	end = time.Since(start)
	fmt.Println(end) // 401.2207ms

	start = time.Now()
	fmt.Println(SquareAll(nums, 12)) // [1 4 9 16 25 36 49 64 81 100 121 144]
	end = time.Since(start)
	fmt.Println(end) // 200.5761ms

	RunWithErrors()
}
