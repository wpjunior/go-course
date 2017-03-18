package main

import (
	"fmt"
	"time"
)

func task(name string, channel chan int) {
	for {
		value := <-channel
		fmt.Println("Executing task", value, "by", name)
	}
}

func main() {
	channel := make(chan int)

	go task("goroutine 1", channel)
	go task("goroutine 2", channel)
	go task("goroutine 3", channel)
	go task("goroutine 4", channel)
	go task("goroutine 5", channel)

	for i := 0; i < 100; i++ {
		channel <- i
	}
	// wait all gorotines to finish
	time.Sleep(time.Second)
}
