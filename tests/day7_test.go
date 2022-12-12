package aoc_test

import (
	"adventOfCode2022/aoc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSumOfDirectorySizesExample(t *testing.T) {
	// arrange
	filename := "./data/day7Example.txt"
	expected := 95437

	// act
	result := aoc.GetSumOfDirectorySizes(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetSumOfDirectorySizesActual(t *testing.T) {
	// arrange
	filename := "./data/day7Actual.txt"
	expected := 2031851

	// act
	result := aoc.GetSumOfDirectorySizes(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetSizeOfSmallestDirectoryToDeleteExample(t *testing.T) {
	// arrange
	filename := "./data/day7Example.txt"
	expected := 24933642

	// act
	result := aoc.GetSizeOfSmallestDirectoryToDelete(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetSizeOfSmallestDirectoryToDeleteActual(t *testing.T) {
	// arrange
	filename := "./data/day7Actual.txt"
	expected := 2568781

	// act
	result := aoc.GetSizeOfSmallestDirectoryToDelete(filename)

	// assert
	assert.Equal(t, expected, result)
}
