package aoc_test

import (
	"adventOfCode2022/aoc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopCratesOldCraneExample(t *testing.T) {
	// arrange
	filename := "./data/day5Example.txt"
	expected := "CMZ"
	newCrane := false

	// act
	result := aoc.GetTopCrates(filename, newCrane)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetTopCratesOldCraneActual(t *testing.T) {
	// arrange
	filename := "./data/day5Actual.txt"
	expected := "ZWHVFWQWW"
	newCrane := false

	// act
	result := aoc.GetTopCrates(filename, newCrane)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetTopCratesNewCraneExample(t *testing.T) {
	// arrange
	filename := "./data/day5Example.txt"
	expected := "MCD"
	newCrane := true

	// act
	result := aoc.GetTopCrates(filename, newCrane)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetTopCratesNewCraneActual(t *testing.T) {
	// arrange
	filename := "./data/day5Actual.txt"
	expected := "HZFZCCWWV"
	newCrane := true

	// act
	result := aoc.GetTopCrates(filename, newCrane)

	// assert
	assert.Equal(t, expected, result)
}
