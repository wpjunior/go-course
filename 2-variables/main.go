package main

import "fmt"

func main() {
	// int
	// declarando uma variável utilizando o modelo pré declarativo
	var x int // ou var x int = 10
	x = 10
	fmt.Printf("O valor de 'x' é %d\n", x)

	// declarando uma variável utilizando o modelo auto-declarativo
	a := 20
	fmt.Printf("O valor de 'a' era %d\n", a)
	a = 30
	fmt.Printf("O valor de 'a' agora é %d\n", a)

	// string
	name := "wilson"
	fmt.Println("Meu nome é", name)

	// float64
	pi := 3.14
	fmt.Printf("O valor de PI é %0.6f\n", pi)

	// bool
	verdade := true
	mentira := false

	fmt.Println("Verdade é igual a mentira:", verdade == mentira)
}
