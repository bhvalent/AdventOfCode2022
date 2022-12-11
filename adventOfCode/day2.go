// https://adventofcode.com/2022/day/2

package adventOfCode

import (
	"adventOfCode2022/common"
	"fmt"
	"strings"
)

func GetPartialPuzzleWinningScore(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	total := 0
	for _, line := range lines {
		total += getPartialPuzzleGameScore(getGameFromLine(line))
	}

	return total
}

func GetFullPuzzleWinningScore(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	total := 0
	for _, line := range lines {
		total += getFullPuzzleGameScore(getGameFromLine(line))
	}

	return total
}

func getFullPuzzleGameScore(a string, x string) int {
	score := 0
	if x == "Z" {
		score += 6
	} else if x == "Y" {
		score += 3
	}

	return score + getShapeScoreFromOutcome(a, x)
}

func getPartialPuzzleGameScore(a string, x string) int {
	score := 0
	switch {
	case (a == "B" && x == "Z") || (a == "C" && x == "X") || (a == "A" && x == "Y"):
		score += 6
	case (a == "A" && x == "X") || (a == "B" && x == "Y") || (a == "C" && x == "Z"):
		score += 3
	}
	return score + getShapeScore(x)
}

func getGameFromLine(line string) (string, string) {
	game := strings.Split(line, " ")
	return game[0], game[1]
}

func getShapeScoreFromOutcome(a string, x string) int {
	switch {
	case (x == "Z" && a == "A") || (x == "Y" && a == "B") || (x == "X" && a == "C"):
		return 2
	case (x == "Z" && a == "B") || (x == "Y" && a == "C") || (x == "X" && a == "A"):
		return 3
	case (x == "Z" && a == "C") || (x == "Y" && a == "A") || (x == "X" && a == "B"):
		return 1
	default:
		return 0
	}
}

func getShapeScore(shape string) int {
	switch shape {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}
