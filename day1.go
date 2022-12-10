/* Day 1 Description:
The Elves take turns writing down the number of Calories
contained by the various meals, snacks, rations, etc.
that they've brought with them, one item per line.
Each Elf separates their own inventory from the previous Elf's
inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items'
Calories and end up with the following list:
	1000
	2000
	3000

	4000

	5000
	6000

	7000
	8000
	9000

	10000
This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.

In case the Elves get hungry and need extra snacks,
they need to know which Elf to ask: they'd like to know how many Calories are
being carried by the Elf carrying the most Calories.
In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
*/

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
