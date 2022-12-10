package main

import "testing"

func TestGetWinningScoreExample(t *testing.T) {
	// arrange
	filename := "./data/day2Example.txt"
	expected := 15

	// act
	result := getWinningScore(filename)

	// assert
	if result != expected {
		t.Fatalf("getWinningScore expected %d but got %d", expected, result)
	}
}

func TestGetWinningScoreActual(t *testing.T) {
	// arrange
	filename := "./data/day2Actual.txt"
	expected := 14827

	// act
	result := getWinningScore(filename)

	// assert
	if result != expected {
		t.Fatalf("getWinningScore expected %d but got %d", expected, result)
	}
}

func TestGetWinningScoreWinners(t *testing.T) {
	// arrange
	filename := "./data/day2Winners.txt"
	expected := 24

	// act
	result := getWinningScore(filename)

	// assert
	if result != expected {
		t.Fatalf("getWinningScore expected %d but got %d", expected, result)
	}
}
