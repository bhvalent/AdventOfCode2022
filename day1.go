// https://adventofcode.com/2022/day/1

package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

func getTopNElvesTotalCalories(filename string, n int) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	} else if n < 1 {
		fmt.Print("Invalid Argumnent: n must be greater than 0")
		return 0
	}

	lines, err := getLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	elfCalories, err := getElfCalories(lines)
	if err != nil {
		fmt.Print(err.Error())
		return 0
	}
	maxElfCals := elfCalories[len(elfCalories)-n:]

	return sum(maxElfCals)
}

func getElfCalories(lines []string) ([]int, error) {
	elfCalories := []int{}
	elf := []int{}
	for i, line := range lines {
		if line == "" {
			total := sum(elf)
			elfCalories = append(elfCalories, total)
			elf = []int{}
			continue
		} else if i == len(lines)-1 {
			cal, err := strconv.Atoi(line)
			if err != nil {
				msg := fmt.Sprintf("Error converting elf calorie: %s", line)
				return []int{}, errors.New(msg)
			}
			elf = append(elf, cal)
			total := sum(elf)
			elfCalories = append(elfCalories, total)
			continue
		}
		cal, err := strconv.Atoi(line)
		if err != nil {
			msg := fmt.Sprintf("Error converting elf calorie: %s", line)
			return []int{}, errors.New(msg)
		}
		elf = append(elf, cal)
	}

	sort.Ints(elfCalories)

	return elfCalories, nil
}

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
