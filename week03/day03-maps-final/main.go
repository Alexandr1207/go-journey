package main

import (
	"fmt"
	"sort"
)

type Cache struct {
	data map[string]int
}

func newCache() *Cache {
	return &Cache{data: make(map[string]int)}
}

func (c Cache) Set(key string, value int) {
	c.data[key] = value
}

func (c Cache) Get(key string) (int, bool) {
	v, ok := c.data[key]
	return v, ok
}

func MapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, ok := b[key]
		if !ok || valA != valB {
			return false
		}
	}
	return true
}

func GroupAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		runes := []rune(word)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedWord := string(runes)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	return anagramMap
}

func main() {
	// Task 1
	myCache := newCache()
	myCache.Set("Age", 21)
	fmt.Println(myCache.Get("Age")) // 21 true. Map has a pointer to real data.

	// Task 2
	a := map[string]int{"Age": 21, "Course": 3}
	b := map[string]int{"Age": 21, "Course": 3}
	fmt.Println(MapsEqual(a, b))

	// Task 3
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := GroupAnagrams(words)

	fmt.Println("Group result:")
	for key, group := range result {
		fmt.Printf("Key:%s -> Words:%v\n", key, group)
	}
}
