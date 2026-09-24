package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

// Here I solve tasks from CodeWars

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(EachCons(arr, 3))

	arr2 := []int{68, -1, 1, -7, 10, 10}
	fmt.Println(multipleOfIndex(arr2))

	fmt.Println(AmIWilson(11))

	fmt.Println(PowersOfTwo(2))

	fmt.Println(Char(48))

	fmt.Println(Feast("brown bear", "bear claw"))

	fmt.Println(DNAStrand("GTAT"))

	test := []string{"Telescopes", "Glasses", "Eyes", "Monocles"}
	fmt.Println(SortByLength(test))
}

// Task 1. Enumerable Magic #20 - Cascading Subsets
func EachCons(arr []int, n int) [][]int {
	var matrix [][]int
	for i := 0; i < len(arr); i++ {
		end := i + n

		if end > len(arr) {
			break
		}
		window := arr[i:end]
		matrix = append(matrix, window)
	}
	return matrix
}

// Task 2. Multiple of index
func multipleOfIndex(ints []int) []int {
	result := []int{}
	for i := 1; i < len(ints); i++ {
		if ints[i]%i == 0 {
			result = append(result, ints[i])
		}
	}
	return result
}

// Task 3.
func Factorial(number int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(1); i <= number; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func AmIWilson(n int64) bool {
	if n <= 1 {
		return false
	}

	p := Factorial(n - 1)
	p.Add(p, big.NewInt(1))

	nSquared := big.NewInt(n * n)

	remainder := new(big.Int)
	remainder.Mod(p, nSquared)

	return remainder.Cmp(big.NewInt(0)) == 0
}

// Task 4. Powers of 2
// Complete the function that takes a non-negative integer n as input, and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).
func PowersOfTwo(n int) []uint64 {
	result := []uint64{}
	for i := 0; i <= n; i++ {
		result = append(result, uint64(1)<<i)
	}
	return result
}

// Task 5. Get character from ASCII Value
// Write a function which takes a number and returns the corresponding ASCII char for that value.
func Char(code int) string {
	return string(rune(code))
}

// Task 6. Is the string uppercase?
// Create a method to see whether the string is ALL CAPS.
type MyString string

func (s MyString) IsUpperCase() bool {
	return string(s) == strings.ToUpper(string(s))
}

// Task 7. The Feast of Many Beasts
// All of the animals are having a feast! Each animal is bringing one dish. There is just one rule: the dish must start and end with the same letters as the animal's name. For example, the great blue heron is bringing garlic naan and the chickadee is bringing chocolate cake.

// Write a function feast that takes the animal's name and dish as arguments and returns true or false to indicate whether the beast is allowed to bring the dish to the feast.

// Assume that beast and dish are always lowercase strings, and that each has at least two letters. beast and dish may contain hyphens and spaces, but these will not appear at the beginning or end of the string. They will not contain numerals.
func Feast(beast string, dish string) bool {
	return beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1]
}

// Task 8. Complementary DNA
// In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". Your function receives one side of the DNA (string, except for Haskell); you need to return the other complementary side. DNA strand is never empty or there is no DNA at all (again, except for Haskell).
func DNAStrand(dna string) string {
	var sb strings.Builder
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'A':
			sb.WriteByte('T')
		case 'T':
			sb.WriteByte('A')
		case 'C':
			sb.WriteByte('G')
		case 'G':
			sb.WriteByte('C')
		}
	}
	return sb.String()
}

// Task 9. Shortest Word
// Simple, given a string of words, return the length of the shortest word(s).
// String will never be empty and you do not need to account for different data types.
func FindShort(s string) int {
	words := strings.Fields(s)
	shortest := len(words[0])
	for _, word := range words {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	return shortest
}

// Task 10. Sort array by string length
// Write a function that takes an array of strings as an argument and returns a sorted array containing the same strings, ordered from shortest to longest.
func SortByLength(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	return arr
}
