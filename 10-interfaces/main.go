package main

import "fmt"

type Animal interface {
	Gritar() string
}

type Vaca struct{}

func (c *Vaca) Gritar() string {
	return "Muuu"
}

type Cachorro struct{}

func (d *Cachorro) Gritar() string {
	return "Rau rau!"
}

func main() {
	var animal Animal

	animal = &Vaca{}
	chamarMeuAnimal(animal)

	animal = &Cachorro{}
	chamarMeuAnimal(animal)
}

func chamarMeuAnimal(a Animal) {
	fmt.Println("Meu animal diz:", a.Gritar())
}
