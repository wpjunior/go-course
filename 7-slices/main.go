package main

import "fmt"

func main() {
	linguagens := make([]string, 10, 12)
	linguagens = append(linguagens, "go")
	fmt.Println("Linguagens da globo.com: ", linguagens)
	fmt.Println("Linguagens da globo.com: ", linguagens[1:3])
}
