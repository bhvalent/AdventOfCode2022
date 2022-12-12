// https://adventofcode.com/2022/day/7

package aoc

import (
	"adventOfCode2022/common"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetSumOfDirectorySizes(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	commands := getCommandsFromLines(lines)
	files := getFilesFromCommands(commands)
	calculateDirectorySizes(files)

	total := 0
	for _, f := range files {
		if f.isDirectory && f.size <= 100000 {
			total += f.size
		}
	}

	return total
}

func GetSizeOfSmallestDirectoryToDelete(filename string) int {
	if filename == "" {
		fmt.Print("Invalid Argumnent: filename cannot be empty")
		return 0
	}

	lines, err := common.GetLinesFromFile(filename)

	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	commands := getCommandsFromLines(lines)
	files := getFilesFromCommands(commands)
	calculateDirectorySizes(files)

	sort.Sort(sortedFiles(files))

	largestDirectory := files[len(files)-1]
	spaceAvailable := 70000000 - largestDirectory.size
	spaceNeeded := 30000000 - spaceAvailable

	for _, f := range files {
		if f.isDirectory && f.size >= spaceNeeded {
			return f.size
		}
	}

	return 0
}

type file struct {
	id            uuid.UUID
	directoryPath string
	name          string
	size          int
	isDirectory   bool
}

type sortedFiles []*file

type command struct {
	input  string
	output []string
}

func getCommandsFromLines(lines []string) []command {
	commands := []command{}

	currentCommand := command{}
	for _, line := range lines {
		if line[0] == '$' {
			if currentCommand.input == "" {
				currentCommand.input = line[2:]
				continue
			}
			commands = append(commands, currentCommand)
			currentCommand = command{input: line[2:]}
			continue
		}
		currentCommand.output = append(currentCommand.output, line)
	}
	commands = append(commands, currentCommand)

	return commands
}

func getFilesFromCommands(commands []command) []*file {
	files := []*file{
		{
			id:            uuid.New(),
			directoryPath: "",
			name:          "root",
			isDirectory:   true,
		},
	}
	filepathStart := "root"

	filepath := filepathStart
	instruction := ""
	detail := ""
	for _, com := range commands {
		instruction = com.input[:2]
		detail = ""

		if instruction == "cd" {
			detail = com.input[3:]
			if detail == "/" {
				filepath = filepathStart
			} else if detail == ".." {
				filepath = goUpDirectory(filepath)
			} else {
				filepath = filepath + "/" + detail
			}
		} else if instruction == "ls" {
			fileDetails := []string{}
			fileSize := 0
			for _, item := range com.output {
				if item[:3] == "dir" {
					newFile := file{
						id:            uuid.New(),
						directoryPath: filepath,
						name:          item[4:],
						isDirectory:   true,
					}
					files = append(files, &newFile)
				} else {
					fileDetails = strings.Split(item, " ")
					fileSize, _ = strconv.Atoi(fileDetails[0])

					newFile := file{
						id:            uuid.New(),
						directoryPath: filepath,
						name:          fileDetails[1],
						size:          fileSize,
						isDirectory:   false,
					}
					files = append(files, &newFile)
				}
			}
		}
	}

	return files
}

func goUpDirectory(filepath string) string {
	slashes := common.IndicesOf([]rune(filepath), '/')
	lastSlash := slashes[len(slashes)-1]

	return filepath[:lastSlash]
}

func calculateDirectorySizes(files []*file) int {
	total := 0
	directorySum := 0
	for _, f := range files {
		if f.isDirectory {
			if f.size > 0 {
				continue
			}
			directorySum = calculateDirectorySizes(getDirectoryFiles(files, *f))
			f.size = directorySum
			continue
		}
		total += f.size
	}

	return total
}

func getDirectoryFiles(files []*file, directory file) []*file {
	directoryFiles := []*file{}
	directoryPath := ""
	if directory.directoryPath == "" {
		directoryPath = directory.name
	} else {
		directoryPath = directory.directoryPath + "/" + directory.name
	}

	for _, f := range files {
		if f.id == directory.id {
			continue
		}

		if strings.Contains(f.directoryPath, directoryPath) {
			directoryFiles = append(directoryFiles, f)
		}
	}

	return directoryFiles
}

func (f sortedFiles) Len() int {
	return len(f)
}

func (f sortedFiles) Less(i, j int) bool {
	return f[i].size < f[j].size
}

func (f sortedFiles) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
