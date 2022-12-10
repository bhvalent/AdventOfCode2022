package main

import "testing"

func TestGetPrioritySumExample(t *testing.T) {
	// arrange
	filename := "./data/day3Example.txt"
	expected := 157

	// act
	result := getPrioritySum(filename)

	// assert
	if result != expected {
		t.Fatalf("getPrioritySum expected %d but got %d", expected, result)
	}
}

func TestGetPrioritySumActual(t *testing.T) {
	// arrange
	filename := "./data/day3Actual.txt"
	expected := 8298

	// act
	result := getPrioritySum(filename)

	// assert
	if result != expected {
		t.Fatalf("getPrioritySum expected %d but got %d", expected, result)
	}
}
