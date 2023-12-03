package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Schematic struct {
	Coordinates [][]Coordinate
	Numbers     []Number
}

type Coordinate struct {
	Value rune
	X     int
	Y     int
}

func (c Coordinate) IsSymbol() bool {
	return c.Value != '.' && !unicode.IsDigit(c.Value)
}

type Number struct {
	Coordinates []Coordinate
	Value       int
}

func (s Schematic) IsSymbol(x, y int) bool {
	coordinates := s.Coordinates
	return x >= 0 && x < len(coordinates) && y >= 0 && y < len(coordinates) && coordinates[y][x].IsSymbol()
}

func read() (string, error) {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Part1(input string) (string, error) {
	schematic, err := parseSchematic(input)
	if err != nil {
		return "", err
	}

	var sum int
	for _, number := range schematic.Numbers {
		if schematic.IsPartNumber(number) {
			sum += number.Value
		}
	}
	result := strconv.Itoa(sum)
	return result, nil
}

func (s Schematic) IsPartNumber(number Number) bool {
	for _, coordinate := range number.Coordinates {
		x := coordinate.X
		y := coordinate.Y
		if s.IsSymbol(x-1, y-1) ||
			s.IsSymbol(x-1, y) ||
			s.IsSymbol(x-1, y+1) ||
			s.IsSymbol(x, y) ||
			s.IsSymbol(x, y-1) ||
			s.IsSymbol(x, y+1) ||
			s.IsSymbol(x+1, y-1) ||
			s.IsSymbol(x+1, y) ||
			s.IsSymbol(x+1, y+1) {
			return true
		}
	}

	return false
}

func parseSchematic(input string) (*Schematic, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	coordinates := make([][]Coordinate, len(lines))
	numbers := make([]Number, 0)
	for y, line := range lines {
		line := strings.TrimSpace(line)
		coordinates[y] = make([]Coordinate, len(line))

		var currentNumber strings.Builder
		currentNumberCoordinates := make([]Coordinate, 0)
		for x, r := range line {
			coordinate := Coordinate{
				Value: r,
				X:     x,
				Y:     y,
			}

			if unicode.IsDigit(r) {
				currentNumberCoordinates = append(currentNumberCoordinates, coordinate)
				currentNumber.WriteRune(r)
			} else if currentNumber.Len() > 0 {
				numStr := currentNumber.String()
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return nil, err
				}

				numbers = append(numbers, Number{Value: num, Coordinates: currentNumberCoordinates})
				currentNumber.Reset()
				currentNumberCoordinates = make([]Coordinate, 0)
			}

			coordinates[y][x] = coordinate
		}
		if currentNumber.Len() > 0 {
			numStr := currentNumber.String()
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}

			numbers = append(numbers, Number{Value: num, Coordinates: currentNumberCoordinates})
		}
	}

	schematic := Schematic{
		Coordinates: coordinates,
		Numbers:     numbers,
	}

	return &schematic, nil
}

func Part2(input string) (string, error) {
	return "", nil
}

func p1main() {
	out, err := read()
	if err != nil {
		log.Fatalf("Could not read input: %v", err)
	}

	res, err := Part2(out)
	fmt.Println(res)
}
