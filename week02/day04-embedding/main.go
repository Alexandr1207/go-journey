package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " makes a sound"
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() string {
	return "Woof! I'm " + d.Name
}

type Cat struct {
	Animal
	Indoor bool
}

type Swimmer struct {
}

func (s Swimmer) Swim() string {
	return "swimming"
}

type Flyer struct {
}

func (f Flyer) Fly() string {
	return "flying"
}

type Duck struct {
	Animal
	Swimmer
	Flyer
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(s string)
}

type ReadWriter interface {
	Reader
	Writer
}

type Student struct {
	ReadWriter
	Name string
}

func (s Student) Read() string {
	return s.Name + " is reading"
}

func (student Student) Write(s string) {
	fmt.Println(student.Name + " is writing " + s)
}

func main() {
	d1 := Dog{}
	d1.Animal.Name = "Chappy"
	c1 := Cat{}
	c1.Animal.Name = "Barsik"
	fmt.Println(d1.Speak()) // Woof! I'm Chappy
	fmt.Println(c1.Speak()) // Barsik makes a sound

	duck1 := Duck{}
	duck1.Animal.Name = "Donald"
	fmt.Println(duck1.Speak(), duck1.Swim(), duck1.Fly()) // Donald makes a sound swimming flying

	var rw ReadWriter = Student{Name: "Alex"}
	fmt.Println(rw.Read()) // Alex is reading
	rw.Write("some text")  // Alex is writing some text

	fmt.Println(d1.Animal.Speak()) // Chappy makes a sound
}
