package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade  int
}

func (p *Pessoa) String() string {
	if p.Idade < 18 {
		return fmt.Sprintf("Olá jovem %s, seja bem vindo", p.Nome)
	} else if p.Idade > 50 {
		return fmt.Sprintf("Olá senhor %s, seja bem vindo", p.Nome)
	}

	return fmt.Sprintf("Olá %s, seja bem vindo", p.Nome)
}

func (p *Pessoa) privateMethod() int {
	return p.Idade * 4
}

func main() {
	pessoa := Pessoa{Nome: "júnior", Idade: 17}
	fmt.Println(pessoa.String())

	pessoa = Pessoa{Nome: "Wilson", Idade: 27}
	fmt.Println(person.String())

	pessoa = Pessoa{Nome: "Antônio", Idade: 60}
	fmt.Println(person.String())

}
