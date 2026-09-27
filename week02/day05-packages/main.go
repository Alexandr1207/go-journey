package main

import (
	"fmt"
	"myfirstgo/week02/day05-packages/library"
	"myfirstgo/week02/day05-packages/shapes"
)

func main() {
	b1 := library.Book{Title: "Martin Iden", Author: "Jack London", Year: 1908, Price: 166}
	fmt.Println(b1.Describe()) // Martin Iden by Jack London (1908), ID:

	newBook := library.NewBook("Women", "Charles Bukowski", 1978, 168)
	fmt.Println(newBook.Describe()) // Women by Charles Bukowski (1978), ID: WomenCharles Bukowski

	c1 := shapes.Circle{Radius: 6}
	shapes.PrintShapeInfo(c1) // 113.09733552923255 37.69911184307752
	r1 := shapes.Rectangle{Width: 4, Height: 3}
	shapes.PrintShapeInfo(r1) // 12 14
}
