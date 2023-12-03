package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var filePath = flag.String("file", "part1test.txt", "File path to the cube game results file")
var defaultMaxPossibles = Set{
	Red:   12,
	Green: 13,
	Blue:  14,
}

type Game struct {
	ID  int
	Max Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func computePossible(contents string, maxPossible Set) (int, error) {
	games := strings.Split(contents, "\n")

	var total int
	for _, g := range games {
		game, err := parseGame(g)
		if err != nil {
			return 0, err
		}
		if game == nil {
			continue
		}

		log.Printf("Game %d has (%d, %d, %d)", game.ID, game.Max.Red, game.Max.Green, game.Max.Blue)
		if isGamePossible(game, maxPossible) {
			total += game.ID
		}
	}
	return total, nil
}

func isGamePossible(game *Game, maxPossible Set) bool {
	return game.Max.Red <= maxPossible.Red && game.Max.Green <= maxPossible.Green && game.Max.Blue <= maxPossible.Blue
}

func parseGame(contents string) (*Game, error) {
	// The contents must contain at minumum "Game #: "
	if len(contents) < 8 {
		return nil, nil
	}

	i := strings.Index(contents, ":")

	// Skip "Game " and up to the ':' is the ID.
	id, err := strconv.ParseInt(contents[5:i], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse ID from %q as int: %w", contents[5:i], err)
	}

	game := &Game{ID: int(id)}

	sets := strings.Split(contents[i+1:], ";")
	for _, set := range sets {
		// Just in case a set could possibly have a color listed twice,
		// I will sum up the colors first.
		var r, g, b int

		parts := strings.Split(set, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			i = strings.Index(part, " ")
			x, err := strconv.ParseInt(part[:i], 10, 32)
			if err != nil {
				return nil, fmt.Errorf("Cannot parse value from %q (int %q) as int: %w", part[:i], part, err)
			}
			v := int(x)
			t := part[i+1:]

			if t == "red" {
				r += v
			} else if t == "green" {
				g += v
			} else if t == "blue" {
				b += v
			}
		}

		if r > game.Max.Red {
			game.Max.Red = r
		}
		if g > game.Max.Green {
			game.Max.Green = g
		}
		if b > game.Max.Blue {
			game.Max.Blue = b
		}
	}
	return game, nil
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

	score, err := computePossible(string(contents), defaultMaxPossibles)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total: %d", score)
}

