package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

type Engine struct {
	HorsePower int
	Type       string
}

type Car struct {
	Brand  string
	Model  string
	Engine Engine
}

type Monitor struct {
	Brand  string
	Matrix string
	Herz   int
}

// Task 1. Describe string
func (b Book) Describe() string {
	res := b.Title + " by " + b.Author + " (" + strconv.Itoa(b.Year) + ")"
	return res
}

// Task 2. Rewrite to pointer receiver
func (book *Book) ApplyDiscount(percent float64) {
	p := (book.Price / 100) * percent
	book.Price -= p
}

// Task 3.
func (c *Car) AddHorsepower(hp int) {
	c.Engine.HorsePower += hp
}

// Task 4. Return new struct(pattern "builder")
func (b Book) WithDiscount(percent float64) Book {
	p := (b.Price / 100) * percent
	b.Price -= p
	return b
}

// Task 5. My iwn struct
func (m Monitor) DescribeMonitor() string {
	res := m.Brand + " with " + strconv.Itoa(m.Herz) + " Hz" + " and " + m.Matrix + " matrix"
	return res
}

func (m *Monitor) AddHerz(h int) {
	m.Herz += h
}

func main() {
	b1 := Book{Title: "Martind Iden", Author: "Jack London", Year: 1908, Price: 164}
	fmt.Println(b1.Describe()) // Martind Iden by Jack London (1908)

	b1.ApplyDiscount(50)
	fmt.Println(b1.Price)

	e1 := Engine{HorsePower: 102, Type: "Gasoline"}
	octavia := Car{Brand: "Skoda", Model: "Octavia", Engine: e1}

	octavia.AddHorsepower(35)
	fmt.Println(octavia.Engine.HorsePower) // 102 if no pointer and 137 if pointer

	b2 := Book{Title: "The witcher", Author: "Andjey", Year: 1993, Price: 144}
	newBook := b2.WithDiscount(50)
	fmt.Println(newBook.Price) // pointer when we need to change struct, builder when we need to create new struct with changes

	m1 := Monitor{Brand: "Samsung", Matrix: "IPS", Herz: 60}
	fmt.Println(m1.DescribeMonitor()) // Samsung with 60 Hz and IPS matrix
	m1.AddHerz(10)
	fmt.Println(m1.DescribeMonitor()) // Samsung with 70 Hz and IPS matrix
}
