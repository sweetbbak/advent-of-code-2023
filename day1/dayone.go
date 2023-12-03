package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func linevalue(line string) int {
	r := regexp.MustCompile("[^0-9]")
	res := r.ReplaceAllString(line, "")
	return int(res[0]-'0')*10 + int(res[len(res)-1]-'0')
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		sum += linevalue(line)
	}
	return sum
}

func part2(input []string) int {
	replacements := [][2]string{
		{"one", "o1e"},
		{"two", "t2o"},
		{"three", "t3e"},
		{"four", "4"},
		{"five", "5e"},
		{"six", "6"},
		{"seven", "7n"},
		{"eight", "e8t"},
		{"nine", "n9e"},
	}

	sum := 0
	for _, line := range input {
		for _, replacement := range replacements {
			line = strings.ReplaceAll(line, replacement[0], replacement[1])
		}
		sum += linevalue(line)
	}

	return sum
}

// ReadFileAsLines reads a file and returns its contents as a slice of strings
// and removes the last line if it is empty
func ReadFileAsLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

// SliceMemberOrEmptyString returns the member of a slice at the given index,
// or an empty string if the index is out of bounds
func SliceMemberOrEmptyString(slice []string, index int) string {
	if index < len(slice) {
		return slice[index]
	}
	return ""
}

// Atoi converts a string to an int, ignoring errors (return zero instead)
func Atoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func main() {
	fmt.Println("Part 1: ",
		part1(ReadFileAsLines("input.txt")))

	fmt.Println("Part 2: ",
		part2(ReadFileAsLines("input.txt")))
}
