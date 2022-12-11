// https://adventofcode.com/2022/day/4

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getNumberOfEnvelopingPairs(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := getLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	sum := 0
	for _, line := range lines {
		sr1, sr2 := getPairSectionRanges(line)
		isSectionFullyContained := (sr1.low <= sr2.low && sr1.high >= sr2.high) || (sr2.low <= sr1.low && sr2.high >= sr1.high)
		if isSectionFullyContained {
			sum++
		}
	}

	return sum
}

func getNumberOfOverlappingPairs(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := getLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	sum := 0
	for _, line := range lines {
		sr1, sr2 := getPairSectionRanges(line)
		isSectionFullyContained := isSectionInRange(sr1.low, sr2) ||
			isSectionInRange(sr1.high, sr2) ||
			isSectionInRange(sr2.low, sr1) ||
			isSectionInRange(sr2.high, sr1)
		if isSectionFullyContained {
			sum++
		}
	}

	return sum
}

type sectionRange struct {
	low  int
	high int
}

func getPairSectionRanges(line string) (sectionRange, sectionRange) {
	sectionRanges := strings.Split(line, ",")
	sr1 := getSectionRange(sectionRanges[0])
	sr2 := getSectionRange(sectionRanges[1])

	return sr1, sr2
}

func getSectionRange(sections string) sectionRange {
	sectionBoundaries := strings.Split(sections, "-")

	var sr sectionRange
	sr.low, _ = strconv.Atoi(sectionBoundaries[0])
	sr.high, _ = strconv.Atoi(sectionBoundaries[1])

	return sr
}

func isSectionInRange(section int, secRange sectionRange) bool {
	return section >= secRange.low && section <= secRange.high
}
