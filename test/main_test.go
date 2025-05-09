package main

import "testing"

func TestAdd(t *testing.T) {
	totalExpected := 10
	valueA := 5
	valueB := 5

	total := Add(valueA, valueB)
	
	if total != totalExpected {
		t.Errorf("Sum was incorrect, got %d, expected %d", total, totalExpected)
	}
}
