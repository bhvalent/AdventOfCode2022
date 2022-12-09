package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func getLinesFromFile(str string) ([]string, error) {
	lines := []string{}

	if str == "" {
		return lines, errors.New("Filename cannot be empty!")
	}

	file, err := os.Open(str)

	if err != nil {
		errMsg := fmt.Sprintf("Error opening %s: %s", str, err)
		return lines, errors.New(errMsg)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
