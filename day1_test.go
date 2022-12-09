package main

import "testing"

func TestGetElfWithMostCaloriesExample(t *testing.T) {
	// arrange
	filename := "./data/day1Example.txt"
	expected := 24000

	// act
	result := getElfWithMostCalories(filename)

	// assert
	if result != expected {
		t.Fatalf("getElfWithMostCalories expected %d but got %d", expected, result)
	}
}

func TestGetElfWithMostCaloriesActual(t *testing.T) {
	// arrange
	filename := "./data/day1Actual.txt"
	expected := 68442

	// act
	result := getElfWithMostCalories(filename)

	// assert
	if result != expected {
		t.Fatalf("getElfWithMostCalories expected %d but got %d", expected, result)
	}
}
