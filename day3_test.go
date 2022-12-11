package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrioritySumExample(t *testing.T) {
	// arrange
	filename := "./data/day3Example.txt"
	expected := 157

	// act
	result := getPrioritySum(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetPrioritySumActual(t *testing.T) {
	// arrange
	filename := "./data/day3Actual.txt"
	expected := 8298

	// act
	result := getPrioritySum(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetBadgePrioritySumExample(t *testing.T) {
	// arrange
	filename := "./data/day3Example.txt"
	expected := 70

	// act
	result := getBadgePrioritySum(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetBadgePrioritySumActual(t *testing.T) {
	// arrange
	filename := "./data/day3Actual.txt"
	expected := 2708

	// act
	result := getBadgePrioritySum(filename)

	// assert
	assert.Equal(t, expected, result)
}
