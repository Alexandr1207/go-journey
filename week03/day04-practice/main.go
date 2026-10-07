package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

// Task 1. Find the middle element. DONE!
func Gimme(array [3]int) int {
	arr := array[:]
	maxi := slices.Max(arr)
	mini := slices.Min(arr)
	res := 0
	for i := range arr {
		if arr[i] != maxi && arr[i] != mini {
			res = i
		}
	}
	return res
}

// Task 2. Find the next perfect square! DONE!
func FindNextSquare(n int64) int64 {
	if n < 0 {
		return -1
	}

	r := int64(math.Sqrt(float64(n)))

	if r*r != n {
		return -1
	}
	next := r + 1
	return next * next
}

// Task 3. Fibonacci. DONE!
func Fib(n int) int {
	if n <= 1 {
		return n
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

// Task 4. Indexed capitalization. DONE!
func Capitalize(st string, arr []int) string {
	runes := []rune(st)

	for _, idx := range arr {
		if idx >= 0 && idx < len(runes) {
			runes[idx] = rune(rune(runes[idx]) - 'a' + 'A')
		}
	}

	return string(runes)
}

// Task 5. Help Suzuki rake his garden! DONE!
func RakeGarden(garden string) string {
	items := strings.Fields(garden)

	for i, item := range items {
		if item != "rock" && item != "gravel" {
			items[i] = "gravel"
		}
	}

	return strings.Join(items, " ")
}

func main() {
	fmt.Println(FindNextSquare(121))
	fmt.Println(Fib(5))
}
