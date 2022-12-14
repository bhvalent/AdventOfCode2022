// https://adventofcode.com/2022/day/3

package aoc

import (
	"adventOfCode2022/common"
	"fmt"
	"strings"
	"unicode"
)

func GetPrioritySum(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	sum := 0
	for _, line := range lines {
		comp1, comp2 := getCompartments(line)
		commonItem := getCommonItem(comp1, comp2)
		sum += getItemPriority(commonItem)
	}

	return sum
}

func GetBadgePrioritySum(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	elfGroup := []string{}
	badgePriorities := []int{}
	for _, line := range lines {
		elfGroup = append(elfGroup, line)
		if len(elfGroup) < 3 {
			continue
		}
		priority := getElfGroupBadgePriority(elfGroup)
		badgePriorities = append(badgePriorities, priority)
		elfGroup = []string{}
	}

	return common.Sum(badgePriorities)
}

func getCompartments(rucksack string) (string, string) {
	length := len(rucksack)
	return rucksack[:length/2], rucksack[length/2:]
}

func getCommonItem(comp1, comp2 string) rune {
	for _, letter := range comp1 {
		if strings.ContainsRune(comp2, letter) {
			return letter
		}
	}
	return '-'
}

func getItemPriority(item rune) int {
	if item == '-' {
		return 0
	} else if unicode.IsLower(item) {
		return int(item) - 96
	} else if unicode.IsUpper(item) {
		return int(item) - 38
	}

	return 0
}

func getElfGroupBadgePriority(elfGroup []string) int {
	for _, letter := range elfGroup[0] {
		if strings.ContainsRune(elfGroup[1], letter) && strings.ContainsRune(elfGroup[2], letter) {
			return getItemPriority(letter)
		}
	}

	return 0
}
