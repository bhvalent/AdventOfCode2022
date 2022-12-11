// https://adventofcode.com/2022/day/6

package aoc

import (
	"adventOfCode2022/common"
	"fmt"
)

func GetNumOfCharsBeforeNDistinct(filename string, n int) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	dataStream := lines[0]
	marker := ""
	for i, _ := range dataStream {
		if i < n-1 {
			continue
		}
		marker = dataStream[i-(n-1) : i+1]
		if isStartOfPacket(marker) {
			return i + 1
		}
	}

	return 0
}

func isStartOfPacket(marker string) bool {
	startingLength := len(marker)
	return startingLength == len(common.Distinct([]rune(marker)))
}
