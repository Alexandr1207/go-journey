package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	mu sync.Mutex
	n  map[string]int
}

func (s *SafeCounter) Inc(name string) {
	s.mu.Lock()
	s.n[name]++
	s.mu.Unlock()
}

func (s *SafeCounter) Value(name string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.n[name]
}

type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Withdraw(n int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance >= n { // проверили
		a.balance -= n // списали
		return true
	}
	return false
}

type Cache struct {
	mu   sync.Mutex
	data map[string]int
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]int)}
}

func (c *Cache) SafeCache(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

type SixTask struct {
	mu sync.Mutex
}

func (s *SixTask) inner() {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println("inner")
}

func (s *SixTask) outer() {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println("outer")
	s.inner() // здесь дедлок: mu уже захвачен нами же
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter) // 10000

	var newCounter atomic.Int64
	for range 10 {
		wg.Go(func() {
			for range 1000 {
				newCounter.Add(1)
			}
		})
	}
	wg.Wait()
	fmt.Println(newCounter.Load()) // 10000

	count := SafeCounter{n: make(map[string]int)}
	for range 1000 {
		wg.Go(func() {
			count.Inc("SomeKey")
		})
	}
	wg.Wait()
	fmt.Println(count.Value("SomeKey")) // 1000

	acc := Account{balance: 100}
	for range 10 {
		wg.Go(func() {
			fmt.Println(acc.Withdraw(50)) // true true, others is false
		})
	}
	wg.Wait()
	fmt.Println(acc.balance) // 0

	// fatal error: concurrent map writes
	task := make(map[int]int)
	for i := range 100 {
		wg.Go(func() {
			task[i] = i
		})
	}
	wg.Wait()

	newTask := NewCache()
	for i := range 100 {
		wg.Go(func() {
			newTask.SafeCache(strconv.Itoa(i), i)
		})
	}
	wg.Wait()

	s := SixTask{}
	s.inner()
	s.outer() // fatal error: all goroutines are asleep - deadlock!
}
