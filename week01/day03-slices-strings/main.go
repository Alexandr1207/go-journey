package main

import (
	"fmt"
	"strings"
)

func main() {
	BaseAppend()
	fmt.Println(ReverseString("привет"))

	for key, val := range CountWords("маша даша саша гоша даша") {
		fmt.Println(key, val)
	}

	DoubleSlice()

	Experiment()
}

func BaseAppend() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	s := a[2:7]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	a = append(a, 11)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CountWords(s string) map[string]int {
	m := make(map[string]int)
	str := strings.Fields(s)
	for _, w := range str {
		if m[w] == 0 {
			m[w] = 1
		} else {
			m[w]++
		}
	}
	return m
}

func DoubleSlice() {
	matrix := make([][]int, 3)
	num := 1
	for i := range 3 {
		matrix[i] = make([]int, 3)
	}

	for i := range 3 {
		for j := range 3 {
			matrix[i][j] = num
			num += 1
		}
	}

	for i := range 3 {
		fmt.Println(matrix[i])
	}
}

func Experiment() {
	a := []int{1, 2, 3}
	b := a // b refers to a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}
