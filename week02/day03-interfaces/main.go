package main

import (
	"fmt"
	"math"
	"strconv"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

func PrintShapeInfo(s Shape) {
	fmt.Println(s.Area(), s.Perimeter())
}

type Descriable interface {
	Describe() string
}

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

type Monitor struct {
	Brand  string
	Matrix string
	Herz   int
}

func (b Book) Describe() string {
	res := b.Title + " by " + b.Author + " (" + strconv.Itoa(b.Year) + ")"
	return res
}

func (m Monitor) Describe() string {
	res := m.Brand + " with " + strconv.Itoa(m.Herz) + " Hz" + " and " + m.Matrix + " matrix"
	return res
}

func PrintAll(items []Descriable) {
	for _, i := range items {
		switch v := i.(type) {
		case Book:
			fmt.Println(v.Describe(), v.Year)
		case Monitor:
			fmt.Println(v.Describe(), v.Brand)
		default:
		}
	}
}

type MyError struct{}

func (e *MyError) Error() string {
	return "Somthig went wrong"
}

func mayFail() error {
	var e *MyError
	return e
}

func main() {
	shapes := []Shape{Circle{4}, Rectangle{2, 4}}
	PrintShapeInfo(shapes[0]) // 50.26548245743669 25.132741228718345
	PrintShapeInfo(shapes[1]) // 8 12

	b1 := Book{Title: "Martind Iden", Author: "Jack London", Year: 1908, Price: 164}
	m1 := Monitor{Brand: "Samsung", Matrix: "IPS", Herz: 60}
	descs := []Descriable{b1, m1}
	fmt.Println(descs[0].Describe()) // Martind Iden by Jack London (1908)
	fmt.Println(descs[1].Describe()) // Samsung with 60 Hz and IPS matrix
	PrintAll(descs)                  // Martind Iden by Jack London (1908) 1908 \n Samsung with 60 Hz and IPS matrix Samsung

	result := mayFail()
	fmt.Println(result == nil) // false
}
