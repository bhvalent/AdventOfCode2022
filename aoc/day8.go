// https://adventofcode.com/2022/day/8

package aoc

import (
	"adventOfCode2022/common"
	"fmt"
	"sort"
	"strconv"
)

func GetNumOfVisibleTrees(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	forest, inverseForest := processForestData(lines)

	numOfInteriorVisibleTrees := getInteriorVisibleTrees(forest, inverseForest)

	return numOfInteriorVisibleTrees + (len(forest) * 2) + (len(inverseForest) * 2) - 4
}

func GetMaxScenicTreeScore(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	forest, inverseForest := processForestData(lines)

	maxScenicScore := 0
	score := 0
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			score = getTreeScenicScore(i, j, forest, inverseForest)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}

	return maxScenicScore
}

func processForestData(data []string) ([][]int, [][]int) {
	forest := [][]int{}
	inverseForest := [][]int{}

	forestRow := []int{}
	tree := 0
	for _, iVal := range data {
		forestRow = []int{}
		for j, jVal := range iVal {
			tree, _ = strconv.Atoi(string(jVal))
			forestRow = append(forestRow, tree)
			if len(inverseForest) <= j {
				inverseForest = append(inverseForest, []int{})
			}
			inverseForest[j] = append(inverseForest[j], tree)
		}
		forest = append(forest, forestRow)
	}

	return forest, inverseForest
}

func getInteriorVisibleTrees(forest, inverseForest [][]int) int {
	visibleTreeCount := 0

	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			if isTreeVisibleInRow(j, forest[i]) {
				visibleTreeCount++
				continue
			}

			if isTreeVisibleInRow(i, inverseForest[j]) {
				visibleTreeCount++
				continue
			}
		}
	}

	return visibleTreeCount
}

func isTreeVisibleInRow(treeIndex int, row []int) bool {
	rowCopy := make([]int, len(row))
	copy(rowCopy, row)

	tree := rowCopy[treeIndex]

	sortedLeft := rowCopy[:treeIndex]
	sort.Ints(sortedLeft)
	sortedRight := rowCopy[treeIndex+1:]
	sort.Ints(sortedRight)

	return sortedLeft[len(sortedLeft)-1] < tree || sortedRight[len(sortedRight)-1] < tree
}

func getTreeScenicScore(i, j int, forest, inverseForest [][]int) int {
	row := make([]int, len(forest[i]))
	copy(row, forest[i])
	col := make([]int, len(inverseForest[j]))
	copy(col, inverseForest[j])

	scenicValues := []int{}
	left := row[:j]
	right := row[j+1:]
	up := col[:i]
	down := col[i+1:]

	tree := forest[i][j]
	scenicValues = append(scenicValues, getScenicScoreForDirection(common.Reverse(left), tree))
	scenicValues = append(scenicValues, getScenicScoreForDirection(right, tree))
	scenicValues = append(scenicValues, getScenicScoreForDirection(common.Reverse(up), tree))
	scenicValues = append(scenicValues, getScenicScoreForDirection(down, tree))

	return common.Product(scenicValues)
}

func getScenicScoreForDirection(dir []int, tree int) int {
	if len(dir) == 0 {
		return 0
	}

	score := 0
	for _, d := range dir {
		if d < tree {
			score++
			continue
		}
		score++
		break
	}

	return score
}
