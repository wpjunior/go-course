package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var globalMap = &sync.Map{}

func worker(id int) {
	key := fmt.Sprintf("blah-%d", id)

	globalMap.Store(key, strconv.Itoa(id*id))
	value, found := globalMap.Load(key)

	fmt.Println(key, value, found)
}

func main() {
	for i := 0; i < 100; i++ {
		go worker(i)
	}
	time.Sleep(time.Second)
}
