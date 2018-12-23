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
	id string
	x, y, width, height int64
}

var input []*claim

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
		input = append(input, claim)
	}
}

func parse(s string) (*claim, error) {
	spl := strings.Split(s, "@")
	measures := strings.Split(spl[1], ":")
	coords := strings.Split(measures[0], ",")
	dimensions := strings.Split(measures[1], "x")

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

	return &claim{
		id: strings.TrimSpace(spl[0]),
		x : x,
		y : y,
		width: width,
		height: height,
	}, nil
}

func main() {
	fmt.Printf("Starting with input of size %d.\n", len(input))
	fmt.Printf("Here: %+v\n", input[1372])
}

