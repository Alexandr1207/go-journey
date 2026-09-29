package main

import (
	"fmt"
	"myfirstgo/week02/day06-practice/tasks"
	"myfirstgo/week02/day06-practice/zoo"
)

func main() {
	lion := zoo.NewLion("Alex", 21, 4)
	parrot := zoo.NewParrot("Poppy", 5, true)
	fmt.Println(lion.Info())   // Alex, age 21
	fmt.Println(parrot.Info()) // Poppy, age 5

	fmt.Println(lion.Speak())   // Roar!
	fmt.Println(parrot.Speak()) // Squawk!

	speakers := []zoo.Speaker{lion, parrot}
	zoo.Feed(speakers) // Alex, age 21 - Roar!
	// Poppy, age 5 - Squawk!

	animals := []zoo.Animal{parrot.Animal, lion.Animal}
	filteredAnimals := zoo.FilterByAge(animals, 14)
	fmt.Println(filteredAnimals) // [{Alex 21}]

	arr1 := [][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}
	fmt.Println(tasks.Number(arr1))

	fmt.Println(tasks.NbDig(550, 5))

	fmt.Println(tasks.BandNameGenerator("step-uncle"))
}
