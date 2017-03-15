package main

import "fmt"

type Animal interface {
	Shout() string
}

type Cow struct{}

func (c *Cow) Shout() string {
	return "Muuu"
}

type Dog struct{}

func (d *Dog) Shout() string {
	return "Rau rau!"
}

func main() {
	var animal Animal

	animal = &Cow{}
	callMyAnimal(animal)

	animal = &Dog{}
	callMyAnimal(animal)
}

func callMyAnimal(a Animal) {
	fmt.Println("My animal says:", a.Shout())
}
