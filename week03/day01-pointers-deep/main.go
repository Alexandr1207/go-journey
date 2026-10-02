package main

import "fmt"

type Counter struct {
	value int
}

func newCoutner() *Counter {
	return &Counter{3}
}

func (c *Counter) changeCoutner() {
	c.value = 5
}

func nilCounter() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from painc: ", r)
		}
	}()
	var c *Counter
	c.changeCoutner()
}

func doublePointer(pp **int) {
	**pp += 10
}

func main() {
	newSlice := new([]int)
	*newSlice = append(*newSlice, 3)
	fmt.Println(newSlice) // &[3]
	makeSlice := make([]int, 0)
	makeSlice = append(makeSlice, 3)
	fmt.Println(makeSlice) // [3]
	newCoutner()           // &Counter{...} escapes to heap
	nilCounter()           // Recovered from painc:  runtime error: invalid memory address or nil pointer dereference
	x := 5
	p := &x
	pp := &p
	doublePointer(pp)
	fmt.Println(x)

}
