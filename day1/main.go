package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	f := int64(0)
	seenFrequencies := map[int64]bool{}
	seenFrequencies[f] = true
	fmt.Printf("Starting with frequency %d...\n", f)

	i := 0
	var firstDoubleFreq *int64 = nil
	for firstDoubleFreq == nil {
		i = i + 1
		//fmt.Printf("Start loop %d with frequency %d\n", i, f)

		freq, d, err := loop(f, seenFrequencies)
		if err != nil {
			log.Fatal(err)
		}

		f = freq
		//fmt.Printf("End loop %d with frequency %d.\n", i, f)

		if d != nil {
			firstDoubleFreq = d
		}
	}

	fmt.Printf("First frequency reached twice: %d (in %d loops).\n", *firstDoubleFreq, i)
}


func loop(freq int64, seenFreqs map[int64]bool) (int64, *int64, error) {
	found := false
	var firstDoubleFreq int64 = 0

	f, err := os.Open("input.txt")
	if err != nil {
		return 0, nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		freq, err = updateFrequency(freq, s.Text())
		if err != nil {
			return 0, nil, err
		}

		if found {
			continue
		}

		if seenFreqs[freq] {
			firstDoubleFreq = freq
			found = true
		} else {
			seenFreqs[freq] = true
		}
	}

	if !found {
		return freq, nil, nil
	}

	return freq, &firstDoubleFreq, nil
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