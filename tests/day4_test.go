package adventOfCode_test

import (
	"adventOfCode2022/adventOfCode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumberOfEnvelopingPairsExample(t *testing.T) {
	// arrange
	filename := "./data/day4Example.txt"
	expected := 2

	// act
	result := adventOfCode.GetNumberOfEnvelopingPairs(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumberOfEnvelopingPairsActual(t *testing.T) {
	// arrange
	filename := "./data/day4Actual.txt"
	expected := 487

	// act
	result := adventOfCode.GetNumberOfEnvelopingPairs(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumberOfOverlappingPairsExample(t *testing.T) {
	// arrange
	filename := "./data/day4Example.txt"
	expected := 4

	// act
	result := adventOfCode.GetNumberOfOverlappingPairs(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumberOfOverlappingPairsActual(t *testing.T) {
	// arrange
	filename := "./data/day4Actual.txt"
	expected := 849

	// act
	result := adventOfCode.GetNumberOfOverlappingPairs(filename)

	// assert
	assert.Equal(t, expected, result)
}
