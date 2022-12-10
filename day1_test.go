package main

import "testing"

func TestGetTopNElvesTotalCaloriesExampleOne(t *testing.T) {
	// arrange
	filename := "./data/day1Example.txt"
	expected := 24000

	// act
	result := getTopNElvesTotalCalories(filename, 1)

	// assert
	if result != expected {
		t.Fatalf("getTopNElvesTotalCalories expected %d but got %d", expected, result)
	}
}

func TestGetTopNElvesTotalCaloriesOne(t *testing.T) {
	// arrange
	filename := "./data/day1Actual.txt"
	expected := 68442

	// act
	result := getTopNElvesTotalCalories(filename, 1)

	// assert
	if result != expected {
		t.Fatalf("getTopNElvesTotalCalories expected %d but got %d", expected, result)
	}
}

func TestGetTopNElvesTotalCaloriesExampleThree(t *testing.T) {
	// arrange
	filename := "./data/day1Example.txt"
	expected := 45000

	// act
	result := getTopNElvesTotalCalories(filename, 3)

	// assert
	if result != expected {
		t.Fatalf("getTopNElvesTotalCalories expected %d but got %d", expected, result)
	}
}

func TestGetTopNElvesTotalCaloriesThree(t *testing.T) {
	// arrange
	filename := "./data/day1Actual.txt"
	expected := 204837

	// act
	result := getTopNElvesTotalCalories(filename, 3)

	// assert
	if result != expected {
		t.Fatalf("getTopNElvesTotalCalories expected %d but got %d", expected, result)
	}
}
