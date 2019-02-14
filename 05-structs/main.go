package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	p := Person{
		Name:   "Wilson Júnior",
		Age:    24,
		Active: true,
	}
	fmt.Println("Nome", p.Name)
	fmt.Println("Age", p.Age)
	fmt.Println("Active", p.Active)
}
