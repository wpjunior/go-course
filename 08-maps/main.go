package main

import "fmt"

func main() {
	people := map[string]int{
		"wilson": 25,
		"andre":  26,
	}

	people["junior"] = 10

	for name, age := range people {
		fmt.Printf("Chave %s do mapa valor: %d\n", name, age)
	}

	delete(people, "wilson")
	fmt.Printf("Posição não alocada: %d\n", people["wilson"])
}
