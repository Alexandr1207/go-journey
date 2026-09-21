package main

import "fmt"

type Counter struct {
	value int
}

func main() {
	a := 0
	Increment_first(&a)
	Increment_first(&a)
	fmt.Println(a)

	b := 0
	ModifyValue(b)    // work with local variable and never changes var b
	fmt.Println(b)    // 0
	ModifyPointer(&b) // work with pointer and changes var b
	fmt.Println(b)    // 100

	v := Counter{}
	v.Increment_second()
	v.Increment_second()
	fmt.Println(v.value)

	arr := []int{1, 1, 1, 1, 1}
	fmt.Println(arr)
	DoubleAll(arr) // we need *[]int only when function must rewrite a header of slice(len, cap) or to force him to point to another array
	fmt.Println(arr)
}

func Increment_first(x *int) {
	*x++
}

func ModifyValue(x int) {
	x = 100
}

func ModifyPointer(x *int) {
	*x = 100
}

func (c *Counter) Increment_second() {
	c.value++
}

func DoubleAll(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] *= 2
	}
}
