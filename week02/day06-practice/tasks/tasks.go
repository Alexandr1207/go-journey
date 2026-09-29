package tasks

import (
	"strconv"
	"strings"
	"unicode"
)

// Task 1. Number of People in the Bus
func Number(busStops [][2]int) int {
	passangers := 0
	for _, stops := range busStops {
		passangers += stops[0] - stops[1]
	}
	return passangers
}

// Task 2. Count the Digit
func ContainsDigit(n int, d int) int {
	if n == 0 && d == 0 {
		return 1
	}
	counter := 0
	for n > 0 {
		if n%10 == d {
			counter++
		}
		n /= 10
	}
	return counter
}

func NbDig(n int, d int) int {
	counter := 0
	for i := 0; i <= n; i++ {
		counter += ContainsDigit(i*i, d)
	}
	return counter
}

// Band name generator
func capitalizeHyphenated(word string) string {
	parts := strings.Split(word, "-")
	for i, part := range parts {
		if len(part) == 0 {
			continue
		}
		r := []rune(part)
		r[0] = unicode.ToUpper(r[0])
		parts[i] = string(r)
	}
	return strings.Join(parts, "-")
}

func BandNameGenerator(word string) string {
	runes := []rune(word)
	if len(runes) == 0 {
		return word
	}

	first := unicode.ToLower(runes[0])
	last := unicode.ToLower(runes[len(runes)-1])

	if first == last {
		capitalized := capitalizeHyphenated(word)
		r := []rune(capitalized)
		return string(r) + string(r[1:])
	}

	return "The " + capitalizeHyphenated(word)
}

// Task 4. Sum of integers in string
func SumOfIntegersInString(strng string) int {
	sum := 0
	numStr := ""

	for _, ch := range strng {
		if ch >= '0' && ch <= '9' {
			numStr += string(ch)
		} else {
			if numStr != "" {
				val, _ := strconv.Atoi(numStr)
				sum += val
				numStr = ""
			}
		}
	}

	if numStr != "" {
		val, _ := strconv.Atoi(numStr)
		sum += val
	}

	return sum
}

// Task 5. Incrementer
func Incrementer(n []int) []int {
	for i := 0; i < len(n); i++ {
		n[i] = (n[i] + i + 1) % 10
	}
	return n
}
