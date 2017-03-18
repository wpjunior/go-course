package main

import "fmt"

type MyError struct {
	Code  int
	Cause string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d", e.Code)
}

func doSomething() error {
	return &MyError{
		Code:  100,
		Cause: "Unexpected ...",
	}
}

func main() {
	err := doSomething()

	if myError, ok := err.(*MyError); ok {
		fmt.Println("Code of error:", myError.Code)
		fmt.Println("Cause of error:", myError.Cause)
	}
}
