package main

import "testing"

func TestGetPartialPuzzleWinningScoreExample(t *testing.T) {
	// arrange
	filename := "./data/day2Example.txt"
	expected := 15

	// act
	result := getPartialPuzzleWinningScore(filename)

	// assert
	if result != expected {
		t.Fatalf("getPartialPuzzleWinningScore expected %d but got %d", expected, result)
	}
}

func TestGetPartialPuzzleWinningScoreActual(t *testing.T) {
	// arrange
	filename := "./data/day2Actual.txt"
	expected := 14827

	// act
	result := getPartialPuzzleWinningScore(filename)

	// assert
	if result != expected {
		t.Fatalf("getPartialPuzzleWinningScore expected %d but got %d", expected, result)
	}
}

func TestGetPartialPuzzleWinningScoreWinners(t *testing.T) {
	// arrange
	filename := "./data/day2Winners.txt"
	expected := 24

	// act
	result := getPartialPuzzleWinningScore(filename)

	// assert
	if result != expected {
		t.Fatalf("getPartialPuzzleWinningScore expected %d but got %d", expected, result)
	}
}

func TestGetFullPuzzleWinningScoreExample(t *testing.T) {
	// arrange
	filename := "./data/day2Example.txt"
	expected := 12

	// act
	result := getFullPuzzleWinningScore(filename)

	// assert
	if result != expected {
		t.Fatalf("getFullPuzzleWinningScore expected %d but got %d", expected, result)
	}
}

func TestGetFullPuzzleWinningScoreActual(t *testing.T) {
	// arrange
	filename := "./data/day2Actual.txt"
	expected := 13889

	// act
	result := getFullPuzzleWinningScore(filename)

	// assert
	if result != expected {
		t.Fatalf("getFullPuzzleWinningScore expected %d but got %d", expected, result)
	}
}
