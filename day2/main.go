package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
)

var input []string

func init() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		input = append(input, s.Text())
	}
}

func main() {
	fmt.Printf("Starting with input of size %d.\n", len(input))

	twos := countTwos(input)
	threes := countThrees(input)

	fmt.Printf("twos is %d\n", twos)
	fmt.Printf("threes is %d\n", threes)

	fmt.Printf("Ending with checksum: %d.\n", twos * threes)
}

func countTwos(ids []string) int {
	twos := 0

	for _, id := range ids {
		for _, v := range countChars(id) {
			if v == 2 {
				twos = twos + 1
				break
			}
		}
	}

	return twos
}

func countThrees(ids []string) int {
	threes := 0

	for _, id := range ids {
		for _, v := range countChars(id) {
			if v == 3 {
				threes = threes + 1
				break
			}
		}
	}

	return threes
}

func countChars(c string) map[string]int {
	counts := map[string]int{}
	for _, char := range c {
		counts[string(char)] = counts[string(char)] + 1
	}
	return counts
}