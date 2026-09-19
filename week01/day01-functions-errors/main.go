package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	res := a / b
	return res, nil
}

func main() {
	fmt.Println(divide(4, 2)) //2 <nil>
	fmt.Println(divide(3, 2)) //1 <nil>
	fmt.Println(divide(4, 0)) //0 Деление на ноль запрещено

	var i int
	var s string
	var b bool
	var f float64
	fmt.Println(i, s, b, f) //0  false 0

	var x int = 5
	var y = 5
	z := 5
	fmt.Printf("%T\n", x) //int
	fmt.Printf("%T\n", y) //int
	fmt.Printf("%T\n", z) //int
}
