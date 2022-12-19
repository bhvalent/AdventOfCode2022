package aoc

import (
	"adventOfCode2022/common"
	"fmt"
	"strconv"
	"strings"
)

func GetNumOfDistinctTailLocationsWithTwoKnots(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	instructions := getHeadInstructions(lines)
	tailCoordinates := getTailCoordinates(instructions, 2)

	return len(tailCoordinates)
}

func GetNumOfDistinctTailLocationsWithTenKnots(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	instructions := getHeadInstructions(lines)
	tailCoordinates := getTailCoordinates(instructions, 10)

	return len(tailCoordinates)
}

type headInstruction struct {
	direction knotDirection
	amount    int
}

type knotCoordinate struct {
	x int
	y int
}

type knotDirection string

const (
	undefined knotDirection = ""
	touching  knotDirection = "T"
	right     knotDirection = "R"
	left      knotDirection = "L"
	up        knotDirection = "U"
	down      knotDirection = "D"
)

func getHeadInstructions(lines []string) []headInstruction {
	instructions := []headInstruction{}

	elements := []string{}
	amount := 0
	for _, line := range lines {
		elements = strings.Split(line, " ")
		amount, _ = strconv.Atoi(elements[1])

		instructions = append(
			instructions,
			headInstruction{
				direction: toKnotDirection(elements[0]),
				amount:    amount,
			})
	}

	return instructions
}

func getTailCoordinates(instructions []headInstruction, numOfKnots int) []knotCoordinate {
	tailCoordinates := make([]knotCoordinate, 1)
	knotCoordinates := make([]*knotCoordinate, 0, numOfKnots)
	for i := 0; i < numOfKnots; i++ {
		knotCoordinates = append(knotCoordinates, &knotCoordinate{x: 0, y: 0})
	}

	directions := []knotDirection{}
	for _, instr := range instructions {
		for i := 0; i < instr.amount; i++ {
			for ikc, kc := range knotCoordinates {
				if ikc == 0 {
					kc.moveKnot(instr.direction)
					continue
				}

				directions = getKnotDirections(*knotCoordinates[ikc-1], *kc)
				kc.moveKnot(directions...)
				if ikc == numOfKnots-1 {
					tailCoordinates = append(tailCoordinates, *kc)
				}
			}
		}
	}

	return common.Distinct(tailCoordinates)
}

func (head *knotCoordinate) moveKnot(dir ...knotDirection) {
	for _, d := range dir {
		switch d {
		case right:
			head.x++
		case left:
			head.x--
		case up:
			head.y++
		case down:
			head.y--
		}
	}
}

func getKnotDirections(head knotCoordinate, tail knotCoordinate) []knotDirection {
	x := head.x - tail.x
	y := head.y - tail.y

	if x == 0 && y > 1 {
		return []knotDirection{up}
	} else if x == 0 && y < -1 {
		return []knotDirection{down}
	} else if y == 0 && x > 1 {
		return []knotDirection{right}
	} else if y == 0 && x < -1 {
		return []knotDirection{left}
	} else if (x > 1 && y > 0) || (y > 1 && x > 0) {
		return []knotDirection{up, right}
	} else if (x > 1 && y < 0) || (y < -1 && x > 0) {
		return []knotDirection{down, right}
	} else if (x < -1 && y < 0) || (y < -1 && x < 0) {
		return []knotDirection{down, left}
	} else if (x < -1 && y > 0) || (y > 1 && x < 0) {
		return []knotDirection{up, left}
	}

	return []knotDirection{touching}
}

func toKnotDirection(s string) knotDirection {
	switch s {
	case "L":
		return left
	case "R":
		return right
	case "U":
		return up
	case "D":
		return down
	default:
		return undefined
	}
}
