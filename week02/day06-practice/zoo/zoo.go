package zoo

import (
	"fmt"
	"strconv"
)

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Info() string {
	return a.Name + ", age " + strconv.Itoa(a.Age)
}

type Lion struct {
	Animal
	PrideSize   int
	enclosureID string
}

func (l Lion) Speak() string {
	return "Roar!"
}

type Parrot struct {
	Animal
	CanTalk     bool
	enclosureID string
}

func (p Parrot) Speak() string {
	return "Squawk!"
}

type Speaker interface {
	Speak() string
	Info() string
}

func NewLion(name string, age int, prSize int) *Lion {
	nLion := &Lion{PrideSize: prSize, enclosureID: name + strconv.Itoa(age)}
	nLion.Animal.Age = age
	nLion.Animal.Name = name
	return nLion
}

func NewParrot(name string, age int, cTalk bool) *Parrot {
	nParrot := &Parrot{CanTalk: cTalk, enclosureID: name + strconv.Itoa(age)}
	nParrot.Animal.Name = name
	nParrot.Animal.Age = age
	return nParrot
}

func Feed(animals []Speaker) {
	for _, a := range animals {
		fmt.Println(a.Info(), "-", a.Speak())
	}
}

func FilterByAge(animals []Animal, minAge int) []Animal {
	result := []Animal{}
	for _, elem := range animals {
		if elem.Age >= minAge {
			result = append(result, elem)
		}
	}
	return result
}
