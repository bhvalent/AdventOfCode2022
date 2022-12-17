package aoc_test

import (
	"adventOfCode2022/aoc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumOfVisibleTreesExample(t *testing.T) {
	// arrange
	filename := "./data/day8Example.txt"
	expected := 21

	// act
	result := aoc.GetNumOfVisibleTrees(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfVisibleTreesActual(t *testing.T) {
	// arrange
	filename := "./data/day8Actual.txt"
	expected := 1543

	// act
	result := aoc.GetNumOfVisibleTrees(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetMaxScenicTreeScoreExample(t *testing.T) {
	// arrange
	filename := "./data/day8Example.txt"
	expected := 8

	// act
	result := aoc.GetMaxScenicTreeScore(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetMaxScenicTreeScoreActual(t *testing.T) {
	// arrange
	filename := "./data/day8Actual.txt"
	expected := 595080

	// act
	result := aoc.GetMaxScenicTreeScore(filename)

	// assert
	assert.Equal(t, expected, result)
}
