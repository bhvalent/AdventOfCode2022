package aoc_test

import (
	"adventOfCode2022/aoc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumOfCharsBeforePacketExample(t *testing.T) {
	// arrange
	filename := "./data/day6Example.txt"
	expected := 7

	// act
	result := aoc.GetNumOfCharsBeforePacket(filename)

	// assert
	assert.Equal(t, expected, result)
}

func TestGetNumOfCharsBeforePacketActual(t *testing.T) {
	// arrange
	filename := "./data/day6Actual.txt"
	expected := 1235

	// act
	result := aoc.GetNumOfCharsBeforePacket(filename)

	// assert
	assert.Equal(t, expected, result)
}
