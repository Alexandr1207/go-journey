package main

import "fmt"

func safeCopy(hugeArr []int) []int {
	n := 5
	newSlice := make([]int, n)
	copy(newSlice, hugeArr[:n])
	return newSlice
}

func CloneSlice(s []int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	return newSlice
}

func main() {
	// Task 1
	var s []int
	fmt.Println("Len: ", len(s), ", cap: ", cap(s))
	for i := range 20 {
		s = append(s, i)
		fmt.Println("Len: ", len(s), ", cap: ", cap(s))
	}

	// Task 2
	huge := make([]int, 1000)
	for i := range 1000 {
		huge[i] = i
	}
	small := safeCopy(huge) // this variant is lighter, then just slicing huge array, because there are no pointers to huge arr.
	fmt.Println(small)

	// Task 3
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[2:5]
	fmt.Println("Len: ", len(b), ", cap: ", cap(b))
	c := a[2:5:5]
	fmt.Println("Len: ", len(c), ", cap: ", cap(c))
	b = append(b, 4)
	fmt.Println(a)
	fmt.Println("Len: ", len(a), ", cap: ", cap(a))
	c = append(c, 5)
	fmt.Println(a)
	fmt.Println("Len: ", len(a), ", cap: ", cap(a))

	// Task 4
	fArr := []int{1, 2, 3, 4, 5}
	sArr := CloneSlice(fArr)
	fmt.Println(fArr)
	fmt.Println(sArr)
	sArr = append(sArr, 6)
	fmt.Println(fArr)
	fmt.Println(sArr)
}
