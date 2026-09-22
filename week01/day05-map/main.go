package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]int)
	m["word"] = 1
	fmt.Println(SafeGet(m, "word"))  // 1 true
	fmt.Println(SafeGet(m, "shark")) // 0 false

	univer := make(map[string]map[string]int)
	univer["Alex"] = make(map[string]int)
	univer["Alex"]["Algebra"] = 3
	univer["Alex"]["Russian"] = 4
	univer["Alex"]["Literature"] = 5

	univer["Dasha"] = make(map[string]int)
	univer["Dasha"]["Algebra"] = 5
	univer["Dasha"]["Russian"] = 5
	univer["Dasha"]["Literature"] = 5

	univer["Niyaz"] = make(map[string]int)
	univer["Niyaz"]["Algebra"] = 2
	univer["Niyaz"]["Russian"] = 4
	univer["Niyaz"]["Literature"] = 3

	fmt.Println(AverageGrade(univer, "Niyaz")) // 3

	nums := []int{1, 2, 2, 3, 4, 5, 6}
	fmt.Println(HasDuplicates(nums)) // true

	people := map[string]int{"Alex": 21, "Dasha": 18, "Niyaz": 16}
	fmt.Println(people) // map[Alex:21 Dasha:18 Niyaz:16]
	DeleteChildren(people)
	fmt.Println(people) // map[Alex:21 Dasha:18]

	s := "go go gopher go"
	words := CountWords(s)
	fmt.Println(words)         // map[go:3 gopher:1]
	fmt.Println(GetMax(words)) // go 3
}

func SafeGet(m map[string]int, key string) (int, bool) {
	elem, ok := m[key]
	return elem, ok
}

func AverageGrade(m map[string]map[string]int, student string) float64 {
	grades, ok := m[student]
	if !ok || len(grades) == 0 {
		return 0
	}
	sum := 0.0
	for _, mark := range grades {
		sum += float64(mark)
	}
	return sum / float64(len(grades))
}

func HasDuplicates(nums []int) bool {
	m := make(map[int]struct{})
	isDup := false
	for _, n := range nums {
		if _, ok := m[n]; ok {
			isDup = true
			break
		} else {
			m[n] = struct{}{}
		}
	}
	return isDup
}

func DeleteChildren(m map[string]int) {
	for name, age := range m {
		if age < 18 {
			delete(m, name) // we can delete elem with range because "delete" works with map and key
		}
	}
}

func CountWords(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		m[word]++
	}
	return m
}

func GetMax(m map[string]int) (string, int) {
	name := ""
	counts := 0
	for key, value := range m {
		if m[key] >= counts {
			name = key
			counts = value
		}
	}
	return name, counts
}
