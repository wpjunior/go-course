package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade    int
	Ativo bool
}

func main() {
	p := Pessoa{
		Nome:   "Wilson Júnior",
		Idade:    24,
		Ativo: true,
	}
	fmt.Println("Nome", p.Nome)
	fmt.Println("Idade", p.Idade)
	fmt.Println("Ativo", p.Ativo)
}
