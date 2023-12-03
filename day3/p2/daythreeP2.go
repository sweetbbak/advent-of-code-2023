package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var filePath = flag.String("file", "input.txt", "File path to the engine schematic file")

type Point struct {
	X int
	Y int
}

type Rect struct {
	// Left is minimum X
	Left  int
	Right int

	// Top is minimum Y
	Top    int
	Bottom int
}

func (r *Rect) containsPoint(point *Point) bool {
	return point.X >= r.Left && point.X <= r.Right && point.Y >= r.Top && point.Y <= r.Bottom
}

type EnginePart struct {
	ID       int
	Collider *Rect
}

func (p *EnginePart) print() {
	log.Printf("Part %d: (%d, %d, %d, %d)", p.ID, p.Collider.Left, p.Collider.Right, p.Collider.Top, p.Collider.Bottom)
}

type Symbol struct {
	Char     rune
	Location *Point
}

func (p *Symbol) print() {
	log.Printf("Symbol %q: (%d, %d)", p.Char, p.Location.X, p.Location.Y)
}

type Schematic struct {
	Parts   []*EnginePart
	Symbols []*Symbol
}

func computeRatioSum(contents string) (int, error) {
	schematic, err := parseEngine(contents)
	if err != nil {
		return 0, err
	}

	var total int
	for _, sym := range schematic.Symbols {
		// Gears are "*" symbols that have exactly 2 nearby parts
		if sym.Char != '*' {
			continue
		}

		var nearby []*EnginePart
		for _, part := range schematic.Parts {
			if part.Collider.containsPoint(sym.Location) {
				nearby = append(nearby, part)
			}
		}
		if len(nearby) != 2 {
			continue
		}

		ratio := nearby[0].ID * nearby[1].ID

		total += ratio
		log.Printf("Adding gear %d*%d = %d", nearby[0].ID, nearby[1].ID, ratio)
	}
	return total, nil
}

func parseEngine(contents string) (*Schematic, error) {
	lines := strings.Split(contents, "\n")
	if len(lines) > 1 {
		// All lines must be the same length
		want := len(strings.TrimSpace(lines[0]))
		for i, line := range lines {
			l := strings.TrimSpace(line)
			if len(l) == 0 {
				continue
			}
			if len(l) != want {
				return nil, fmt.Errorf("Line %d has mismatching length. wanted %d, got %d", i+1, want, len(l))
			}
		}
	}

	schematic := &Schematic{}

	for y, line := range lines {
		start := -1
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		for x, char := range line {
			if unicode.IsDigit(char) {
				if start == -1 {
					start = x
				}
			} else {
				if start != -1 {
					id, err := strconv.ParseInt(line[start:x], 10, 32)
					if err != nil {
						return nil, fmt.Errorf("Cannot parse ID from %q[%d:%d] as int: %w", line, start, x, err)
					}
					schematic.Parts = append(schematic.Parts, &EnginePart{
						ID: int(id),
						// Collider is 1 larger than the number since the symbol can be nearby
						Collider: &Rect{Left: start - 1, Right: x, Top: y - 1, Bottom: y + 1},
					})
				}
				start = -1

				if char != '.' {
					schematic.Symbols = append(schematic.Symbols, &Symbol{
						Char:     char,
						Location: &Point{X: x, Y: y},
					})
				}
			}
		}

		if start != -1 {
			id, err := strconv.ParseInt(line[start:], 10, 32)
			if err != nil {
				return nil, fmt.Errorf("Cannot parse ID from %q[%d:] as int: %w", line, start, err)
			}
			schematic.Parts = append(schematic.Parts, &EnginePart{
				ID: int(id),
				// Collider is 1 larger than the number since the symbol can be nearby
				Collider: &Rect{Left: start - 1, Right: len(line), Top: y - 1, Bottom: y + 1},
			})
		}
	}

	return schematic, nil
}

func main() {
	flag.Parse()

	if *filePath == "" {
		log.Fatalf("Must specify the game results file!")
	}

	contents, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	score, err := computeRatioSum(string(contents))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total: %d", score)
}
