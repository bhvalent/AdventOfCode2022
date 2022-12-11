package aoc

import (
	"adventOfCode2022/common"
	"fmt"
	"strconv"
	"strings"
)

func GetTopCrates(filename string, use9001Crane bool) string {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return ""
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return ""
	}

	lineBreakIndex := common.IndexOf(lines, "")
	stackLines := lines[:lineBreakIndex]
	instructionLines := lines[lineBreakIndex+1:]

	stacks := processStackLines(stackLines)
	instructions := processInstructionLines(instructionLines)

	if use9001Crane {
		stacks = moveCratesWith9001(stacks, instructions)
	} else {
		stacks = moveCrates(stacks, instructions)
	}

	var topCrates string
	for _, stack := range stacks {
		topCrates += stack[len(stack)-1]
	}

	return topCrates
}

func processStackLines(lines []string) [][]string {
	indicesLine := lines[len(lines)-1]
	numOfStacksRune := []rune(indicesLine)[len(indicesLine)-2]
	numOfStacks, _ := strconv.Atoi(string(numOfStacksRune))

	var crate string
	stacks := [][]string{}
	for i := len(lines) - 2; i >= 0; i-- {
		for j := 0; j < numOfStacks; j++ {
			crate = processCrate(lines[i][4*j : 4*j+3])
			if crate == " " {
				continue
			}
			if i == len(lines)-2 {
				stacks = append(stacks, []string{crate})
				continue
			}
			stacks[j] = append(stacks[j], crate)
		}
	}

	return stacks
}

func processCrate(crate string) string {
	return string(crate[1])
}

type instruction struct {
	count       int
	start       int
	destination int
}

func processInstructionLines(lines []string) []instruction {
	instructions := []instruction{}
	var count int
	var start int
	var destination int
	for _, line := range lines {
		words := strings.Split(line, " ")
		count, _ = strconv.Atoi(words[1])
		start, _ = strconv.Atoi(words[3])
		destination, _ = strconv.Atoi(words[5])
		instructions = append(
			instructions,
			instruction{
				count:       count,
				start:       start - 1,
				destination: destination - 1,
			})
	}
	return instructions
}

func moveCrates(stacks [][]string, instructions []instruction) [][]string {
	var crate string
	var topCrateIndex int
	for _, instr := range instructions {
		for i := 0; i < instr.count; i++ {
			topCrateIndex = len(stacks[instr.start]) - 1
			crate = stacks[instr.start][topCrateIndex]
			stacks[instr.destination] = append(stacks[instr.destination], crate)
			stacks[instr.start] = stacks[instr.start][:len(stacks[instr.start])-1]
		}
	}

	return stacks
}

func moveCratesWith9001(stacks [][]string, instructions []instruction) [][]string {
	var cratesToMove []string
	var bottomCrateIndex int
	for _, instr := range instructions {
		bottomCrateIndex = len(stacks[instr.start]) - instr.count
		cratesToMove = stacks[instr.start][bottomCrateIndex:]
		stacks[instr.destination] = append(stacks[instr.destination], cratesToMove...)
		stacks[instr.start] = stacks[instr.start][:bottomCrateIndex]
	}

	return stacks
}
