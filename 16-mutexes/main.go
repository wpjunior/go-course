package main

import (
	"fmt"
	"strconv"
)

type SafeMap struct {
	m map[string]string
}

func (s *SafeMap) Set(key, value string) {
	s.m[key] = value
}

func (s *SafeMap) Get(key string) string {
	return s.m[key]
}

func main() {
	s := &SafeMap{
		m: map[string]string{},
	}

	for i := 0; i < 100; i++ {
		go s.Set("blah", strconv.Itoa(i))
		go fmt.Println(s.Get("blah"))
	}

}
