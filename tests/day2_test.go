package aoc_test

import (
	"adventOfCode2022/aoc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPartialPuzzleWinningScoreExample(t *testing.T) {
	// arrange
	filename := "./data/day2Example.txt"
	expected := 15

	// act
	result := aoc.GetPartialPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetPartialPuzzleWinningScoreActual(t *testing.T) {
	// arrange
	filename := "./data/day2Actual.txt"
	expected := 14827

	// act
	result := aoc.GetPartialPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetPartialPuzzleWinningScoreWinners(t *testing.T) {
	// arrange
	filename := "./data/day2Winners.txt"
	expected := 24

	// act
	result := aoc.GetPartialPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetFullPuzzleWinningScoreExample(t *testing.T) {
	// arrange
	filename := "./data/day2Example.txt"
	expected := 12

	// act
	result := aoc.GetFullPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetFullPuzzleWinningScoreActual(t *testing.T) {
	// arrange
	filename := "./data/day2Actual.txt"
	expected := 13889

	// act
	result := aoc.GetFullPuzzleWinningScore(filename)

	// assert
	assert.Equal(t, expected, result)
}
