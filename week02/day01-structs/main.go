package main

import "fmt"

// Task 1. Base struct and variants to create
type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

// Task 2. Embedding struct
type Engine struct {
	HorsePower int
	Type       string
}

type Car struct {
	Brand  string
	Model  string
	Engine Engine
}

func main() {
	// Task 1
	zBook := Book{}
	nBook := Book{Title: "myBook", Author: "Me", Year: 2026, Price: 123.4}
	pBook := Book{"HelloBook", "Frank Sinatra", 1963, 4314.9}
	fmt.Println(zBook, nBook, pBook) // {  0 0} {myBook Me 2026 123.4} {HelloBook Frank Sinatra 1963 4314.9}

	// Task 2
	e1 := Engine{HorsePower: 102, Type: "Gasoline"}
	e2 := Engine{HorsePower: 90, Type: "Gasoline"}

	octavia := Car{Brand: "Skoda", Model: "Octavia", Engine: e1}
	granta := Car{Brand: "Lada", Model: "Granta", Engine: e2}

	fmt.Println(octavia, granta)                          // {Skoda Octavia {102 Gasoline}} {Lada Granta {90 Gasoline}}
	fmt.Println(octavia.Brand, octavia.Engine.HorsePower) // Skoda 102

	// Task 3
	firstBook := Book{Title: "myBook", Author: "Me", Year: 2002, Price: 123.4}
	secondBook := Book{Title: "myBook", Author: "Me", Year: 2002, Price: 123.4}
	fmt.Println(firstBook == secondBook) // true
	secondBook.Author = "Not me"
	fmt.Println(firstBook == secondBook) // false

	// Task 4
	books := []Book{zBook, nBook, pBook, secondBook}
	fmt.Println(FilterByYear(books, 2000)) // [{myBook Me 2026 123.4} {myBook Not me 2002 123.4}]
	fmt.Println(TotalPrice(books))         // 4561.699999999999

	// Task 5
	ApplyDiscount(&nBook, 50)
	fmt.Println(nBook.Price) // 61.7
}

// Task 4. Struct slice
func FilterByYear(books []Book, year int) []Book {
	result := []Book{}
	for _, b := range books {
		if b.Year >= year {
			result = append(result, b)
		}
	}
	return result
}

func TotalPrice(books []Book) float64 {
	result := 0.0
	for _, b := range books {
		result += b.Price
	}
	return result
}

// Task 5. Struct pointer
func ApplyDiscount(book *Book, percent float64) {
	p := (book.Price / 100) * percent
	book.Price -= p
}
