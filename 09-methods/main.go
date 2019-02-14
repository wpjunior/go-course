package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	if p.Age < 18 {
		return fmt.Sprintf("Olá jovem %s, seja bem vindo", p.Name)
	} else if p.Age > 50 {
		return fmt.Sprintf("Olá senhor %s, seja bem vindo", p.Name)
	}

	return fmt.Sprintf("Olá %s, seja bem vindo", p.Name)
}

func (p *Person) privateMethod() int {
	return p.Age * 4
}

func main() {
	person := Person{Name: "júnior", Age: 17}
	fmt.Println(person.String())

	person = Person{Name: "Wilson", Age: 27}
	fmt.Println(person.String())

	person = Person{Name: "Antônio", Age: 60}
	fmt.Println(person.String())

}
