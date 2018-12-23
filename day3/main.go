package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type claim struct {
	id                  string
	x, y, width, height int64
}

var claims []*claim

func init() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		claim, err := parse(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		claims = append(claims, claim)
	}
}

func parse(s string) (*claim, error) {
	spl := strings.Split(s, "@")
	measures := strings.Split(spl[1], ":")
	coords := strings.Split(measures[0], ",")
	dimensions := strings.Split(measures[1], "x")

	id := strings.TrimSpace(spl[0])

	x, err := strconv.ParseInt(strings.TrimSpace(coords[0]), 10, 0)
	if err != nil {
		return nil, err
	}

	y, err := strconv.ParseInt(strings.TrimSpace(coords[1]), 10, 0)
	if err != nil {
		return nil, err
	}

	width, err := strconv.ParseInt(strings.TrimSpace(dimensions[0]), 10, 0)
	if err != nil {
		return nil, err
	}

	height, err := strconv.ParseInt(strings.TrimSpace(dimensions[1]), 10, 0)
	if err != nil {
		return nil, err
	}

	return &claim{id, x, y, width, height}, nil
}

func main() {
	fmt.Printf("Starting with %d claims.\n", len(claims))

	// create grid
	grid := make([][][]string, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([][]string, 1000)
	}

	// fill grid with claims
	for _, claim := range claims {
		for x := claim.x; x < claim.x + claim.width; x++ {
			for y := claim.y; y < claim.y + claim.height; y++ {
				grid[x][y] = append(grid[x][y], claim.id)
			}
		}
	}

	// count overlap: cells with more than 1 claim id
	overlap := 0
	for i:=0; i<1000; i++ {
		for j:=0; j<1000; j++ {
			if len(grid[i][j]) > 1 {
				overlap = overlap + 1
			}
		}
	}

	// source ids
	ids := make(map[string]bool, len(claims))
	for i:=0; i<len(claims); i++ {
		ids[claims[i].id] = true
	}

	for i:=0; i<1000; i++ {
		for j:=0; j<1000; j++ {
			cell := grid[i][j]
			if len(cell) > 1 {
				for k := 0; k < len(cell); k++ {
					ids[cell[k]] = false
				}
			}
		}
	}

	found := ""
	for k, v := range ids {
		if v {
			found = k
		}
	}

	fmt.Printf("Done, overlap: %d.\n", overlap)
	fmt.Printf("Done, id: %+v.\n", found)
}
