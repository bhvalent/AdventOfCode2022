package adventOfCode_test

import (
	"adventOfCode2022/adventOfCode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopNElvesTotalCaloriesExampleOne(t *testing.T) {
	// arrange
	filename := "./data/day1Example.txt"
	expected := 24000

	// act
	result := adventOfCode.GetTopNElvesTotalCalories(filename, 1)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetTopNElvesTotalCaloriesOne(t *testing.T) {
	// arrange
	filename := "./data/day1Actual.txt"
	expected := 68442

	// act
	result := adventOfCode.GetTopNElvesTotalCalories(filename, 1)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetTopNElvesTotalCaloriesExampleThree(t *testing.T) {
	// arrange
	filename := "./data/day1Example.txt"
	expected := 45000

	// act
	result := adventOfCode.GetTopNElvesTotalCalories(filename, 3)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetTopNElvesTotalCaloriesThree(t *testing.T) {
	// arrange
	filename := "./data/day1Actual.txt"
	expected := 204837

	// act
	result := adventOfCode.GetTopNElvesTotalCalories(filename, 3)

	// assert
	assert.Equal(t, expected, result)
}