package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"strconv"
)

var frequency int64 = 0

func main() {

	fmt.Printf("Starting with frequency %d...\n", frequency)


	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		if err := updateFrequency(s.Text()); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("End with frequency %d.\n", frequency)
}

func updateFrequency(s string) error {
	operator := string(s[0])
	operand, err := strconv.ParseInt(s[1:], 10, 64)
	if err != nil {
		return err
	}

	if operator == "+" {
		frequency = frequency + operand
	} else if operator == "-" {
		frequency = frequency - operand
	} else {
		return fmt.Errorf("unknown operator")
	}

	return nil
}