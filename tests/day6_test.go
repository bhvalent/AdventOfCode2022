package aoc_test

import (
	"adventOfCode2022/aoc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumOfCharsBefore4DistinctExample(t *testing.T) {
	// arrange
	filename := "./data/day6Example.txt"
	n := 4
	expected := 7

	// act
	result := aoc.GetNumOfCharsBeforeNDistinct(filename, n)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfCharsBefore4DistinctActual(t *testing.T) {
	// arrange
	filename := "./data/day6Actual.txt"
	n := 4
	expected := 1235

	// act
	result := aoc.GetNumOfCharsBeforeNDistinct(filename, n)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfCharsBefore14DistinctExample(t *testing.T) {
	// arrange
	filename := "./data/day6Example.txt"
	n := 14
	expected := 19

	// act
	result := aoc.GetNumOfCharsBeforeNDistinct(filename, n)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfCharsBefore14DistinctActual(t *testing.T) {
	// arrange
	filename := "./data/day6Actual.txt"
	n := 14
	expected := 3051

	// act
	result := aoc.GetNumOfCharsBeforeNDistinct(filename, n)

	// assert
	assert.Equal(t, expected, result)
}
