package aoc_test

import (
	"adventOfCode2022/aoc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumOfDistinctTailLocationsWithTwoKnotsExample(t *testing.T) {
	// arrange
	filename := "./data/day9Example.txt"
	expected := 13

	// act
	result := aoc.GetNumOfDistinctTailLocationsWithTwoKnots(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfDistinctTailLocationsWithTwoKnotsActual(t *testing.T) {
	// arrange
	filename := "./data/day9Actual.txt"
	expected := 5710

	// act
	result := aoc.GetNumOfDistinctTailLocationsWithTwoKnots(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfDistinctTailLocationsWithTenKnotsExample(t *testing.T) {
	// arrange
	filename := "./data/day9Example.txt"
	expected := 1

	// act
	result := aoc.GetNumOfDistinctTailLocationsWithTenKnots(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfDistinctTailLocationsWithTenKnotsLargeExample(t *testing.T) {
	// arrange
	filename := "./data/day9LargeExample.txt"
	expected := 36

	// act
	result := aoc.GetNumOfDistinctTailLocationsWithTenKnots(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfDistinctTailLocationsWithTenKnotsActual(t *testing.T) {
	// arrange
	filename := "./data/day9Actual.txt"
	expected := 2259

	// act
	result := aoc.GetNumOfDistinctTailLocationsWithTenKnots(filename)

	// assert
	assert.Equal(t, expected, result)
}
