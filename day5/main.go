package main

import (
	"log"
	"fmt"
	"io/ioutil"
	"strings"
)

var polymer string

func init() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	polymer = strings.TrimSpace(string(b))
}

func main() {
	fmt.Printf("Start with polymer: %s\n", polymer)
	fmt.Printf("Start with size: %d\n", len(polymer))

	result := iter(polymer)

	fmt.Printf("End with polymer: %s\n", result)
	fmt.Printf("End with size: %d\n", len(result))

}

func iter(s string) string {
	x, y, found := findFirstReaction(s)
	if !found {
		return s
	}

	if y == len(s) - 1 {
		return iter(s[:x])
	}

	return iter(s[:x] + s[y+1:])
}

func findFirstReaction(s string) (int, int, bool) {
	for i:=0; i<len(s)-1; i++ {
		if s[i] == s[i+1] {
			continue
		}

		first, second := string(s[i]), string(s[i+1])
		if first == strings.ToUpper(second) || first == strings.ToLower(second) {
			return i, i + 1, true
		}
	}
		return -1, -1, false

}