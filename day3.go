// https://adventofcode.com/2022/day/3

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func getPrioritySum(filename string) int {
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
		comp1, comp2 := getCompartments(line)
		commonItem := getCommonItem(comp1, comp2)
		sum += getItemPriority(commonItem)
	}

	return sum
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
