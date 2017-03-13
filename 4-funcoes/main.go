package main

import "fmt"

// exemplo de função privada
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("O resultado de 10 + 20 é", add(10, 20))
}
