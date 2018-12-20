package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"strconv"
)


func main() {
	freq  := int64(0)
	fmt.Printf("Starting with frequency %d...\n", freq)

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		freq, err = updateFrequency(freq, s.Text())
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("End with frequency %d.\n", freq)
}

func updateFrequency(f int64, s string) (int64, error) {
	operator := string(s[0])
	operand, err := strconv.ParseInt(s[1:], 10, 64)
	if err != nil {
		return 0, err
	}

	if operator == "+" {
		return f + operand, nil
	} else if operator == "-" {
		return f - operand, nil
	} else {
		return 0, fmt.Errorf("unknown operator")
	}
}