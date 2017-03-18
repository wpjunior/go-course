package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for {
		fmt.Println("Executing task", name)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("goroutine 1")
	go task("goroutine 2")
	go task("goroutine 3")
	go task("goroutine 4")
	go task("goroutine 5")

	task("main")
}
