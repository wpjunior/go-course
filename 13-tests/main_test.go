package main

import (
	"testing"
)

func TestAddSuccess(t *testing.T) {
	calculator := &Calculator{}
	result := calculator.Add(42, 10)
	if result != 52 {
		t.Error("Result of 42 + 10 must be 52")
	}
}

func TestAddFail(t *testing.T) {
	calculator := &Calculator{}
	result := calculator.Add(42, 10)
	if result != 50 {
		t.Error("Result of 42 + 10 must be 52")
	}
}
