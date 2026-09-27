package library

import "strconv"

type Book struct {
	Title      string
	Author     string
	Year       int
	Price      float64
	internalID string
}

func (b Book) Describe() string {
	res := b.Title + " by " + b.Author + " (" + strconv.Itoa(b.Year) + ")" + ", ID: " + b.internalID
	return res
}

func (book *Book) ApplyDiscount(percent float64) {
	p := (book.Price / 100) * percent
	book.Price -= p
}

func NewBook(title string, author string, year int, price float64) *Book {
	return &Book{Title: title, Author: author, Year: year, Price: price, internalID: title + author}
}
