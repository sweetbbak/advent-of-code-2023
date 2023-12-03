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

type Game struct {
	ID  int
	Max Set
}

func (game *Game) Power() int {
	return game.Max.Red * game.Max.Green * game.Max.Blue
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func computePowerSum(contents string) (int, error) {
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
		power := game.Power()

		log.Printf("Game %d has (%d, %d, %d) - %d", game.ID, game.Max.Red, game.Max.Green, game.Max.Blue, power)
		total += power
	}
	return total, nil
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

	score, err := computePowerSum(string(contents))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total: %d", score)
}
