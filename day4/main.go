package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

var input []string

func init() {
	f, err := os.Open("input_sorted.txt")
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
	fmt.Printf("Starting with %d inputs.\n", len(input))

	// create guard overview
	guards := map[string][]int{}
	lastId := ""
	asleep := 0
	for _, r := range input {
		if strings.Contains(r, "begins shift") {
			id := strings.Split(r, " ")[3]
			lastId = id
			if len(guards[id]) == 0 {
				guards[id] = make([]int, 60)
			}
			continue
		}

		if strings.Contains(r, "falls asleep") {
			asleep = parseMinutes(r)
			continue
		}

		if strings.Contains(r, "wakes up") {
			awake := parseMinutes(r)

			for i := asleep; i < awake; i++ {
				guards[lastId][i] = guards[lastId][i] + 1
			}
		}
	}

	fmt.Printf("Done: %+v\n", guards)

	max := 0
	mostAsleepId := ""
	for id, time := range guards {
		asleep = calcAsleep(time)
		if max < asleep {
			max = asleep
			mostAsleepId = id
		}
	}

	max, minute := 0, 0
	for id, time := range guards {
		if id != mostAsleepId {
			continue
		}

		for i:=0; i<len(time); i++ {
			if max < time[i] {
				max = time[i]
				minute = i
			}
		}
	}

	fmt.Printf("Strategy 1: Guard %q is asleep most at minute: %d\n", mostAsleepId, minute)

	max, maxMinute := 0, 0
	maxId := ""
	for id, time := range guards {
		for i:=0; i<len(time); i++ {
			if max < time[i] {
				max = time[i]
				maxId = id
				maxMinute = i
			}
		}
	}

	fmt.Printf("Strategy 2: Guard %q is asleep most at minute: %d\n", maxId, maxMinute)
}

func calcAsleep(t []int) int {
	sum := 0
	for i:=0; i<len(t); i++ {
		sum = sum + t[i]
	}
	return sum
}

func parseMinutes(s string) int {
	time := strings.Split(s, " ")[1]
	m := strings.Split(time, ":")[1]
	minutes := strings.Replace(m, "]", "", -1)

	n, err := strconv.ParseInt(minutes, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	return int(n)
}