package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumberOfEnvelopingPairsExample(t *testing.T) {
	// arrange
	filename := "./data/day4Example.txt"
	expected := 2

	// act
	result := getNumberOfEnvelopingPairs(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumberOfEnvelopingPairsActual(t *testing.T) {
	// arrange
	filename := "./data/day4Actual.txt"
	expected := 487

	// act
	result := getNumberOfEnvelopingPairs(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumberOfOverlappingPairsExample(t *testing.T) {
	// arrange
	filename := "./data/day4Example.txt"
	expected := 4

	// act
	result := getNumberOfOverlappingPairs(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumberOfOverlappingPairsActual(t *testing.T) {
	// arrange
	filename := "./data/day4Actual.txt"
	expected := 849

	// act
	result := getNumberOfOverlappingPairs(filename)

	// assert
	assert.Equal(t, expected, result)
}
