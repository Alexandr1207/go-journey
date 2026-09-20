package main

import (
	"errors"
	"fmt"
)

func FizzBuzz(n int) {
	for i := 0; i < n; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func CheckAge(age int) error {
	if age < 0 {
		return errors.New("validation error")
	}
	return nil
}

func DeferCheck() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // Last in - First out (2, 1, 0)
	}
}

func SumFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func SumWhile() {
	i := 0
	sum := 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func SumEmpty() {
	sum := 0
	i := 0
	for {
		if sum > 50 {
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	count := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count <= 1
}

func main() {
	FizzBuzz(20)

	age := 3
	if err := CheckAge(age); err == nil {
		fmt.Println("Age is correct")
	} else {
		fmt.Println("Wrong age")
	}

	DeferCheck()

	SumFor()   // 45
	SumWhile() // 45
	SumEmpty() // 55

	fmt.Println(IsPrime(59))
}
