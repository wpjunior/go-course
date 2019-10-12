package main

import "fmt"

func main() {
	pessoas := map[string]int{
		"wilson": 25,
		"andre":  26,
	}

	pessoas["junior"] = 10

	for nome, idade := range pessoas {
		fmt.Printf("Chave %s do mapa valor: %d\n", nome, idade)
	}

	delete(pessoas, "wilson")
	fmt.Printf("Posição não alocada: %d\n", pessoas["wilson"])
}
