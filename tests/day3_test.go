package adventOfCode_test

import (
	"adventOfCode2022/adventOfCode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrioritySumExample(t *testing.T) {
	// arrange
	filename := "./data/day3Example.txt"
	expected := 157

	// act
	result := adventOfCode.GetPrioritySum(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetPrioritySumActual(t *testing.T) {
	// arrange
	filename := "./data/day3Actual.txt"
	expected := 8298

	// act
	result := adventOfCode.GetPrioritySum(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetBadgePrioritySumExample(t *testing.T) {
	// arrange
	filename := "./data/day3Example.txt"
	expected := 70

	// act
	result := adventOfCode.GetBadgePrioritySum(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetBadgePrioritySumActual(t *testing.T) {
	// arrange
	filename := "./data/day3Actual.txt"
	expected := 2708

	// act
	result := adventOfCode.GetBadgePrioritySum(filename)

	// assert
	assert.Equal(t, expected, result)
}
