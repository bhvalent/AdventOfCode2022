package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPartialPuzzleWinningScoreExample(t *testing.T) {
	// arrange
	filename := "./data/day2Example.txt"
	expected := 15

	// act
	result := getPartialPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetPartialPuzzleWinningScoreActual(t *testing.T) {
	// arrange
	filename := "./data/day2Actual.txt"
	expected := 14827

	// act
	result := getPartialPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetPartialPuzzleWinningScoreWinners(t *testing.T) {
	// arrange
	filename := "./data/day2Winners.txt"
	expected := 24

	// act
	result := getPartialPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetFullPuzzleWinningScoreExample(t *testing.T) {
	// arrange
	filename := "./data/day2Example.txt"
	expected := 12

	// act
	result := getFullPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetFullPuzzleWinningScoreActual(t *testing.T) {
	// arrange
	filename := "./data/day2Actual.txt"
	expected := 13889

	// act
	result := getFullPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}
