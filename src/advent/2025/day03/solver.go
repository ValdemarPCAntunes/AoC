package day03

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2025/day/03

const fileName = "input.txt"

func Solve(part1, part2 bool) (r1, r2 Result) {
	data, err := os.ReadFile(fileName)
	elemSep := "\r\n"
	if err != nil {
		log.Fatal(err.Error())
	}
	preparedData := strings.Split(string(data), elemSep)
	if part1 {
		r1 = SolvePart1(preparedData)
		fmt.Printf("Part 1 Answer: %d\n", r1.jolts)
	}
	if part2 {
		r2 = SolvePart2(preparedData)
		fmt.Printf("Part 2 Answer: %d\n", r2.jolts)
	}
	return
}

type Result struct {
	jolts int
}

func SolvePart1(data []string) (result Result) {
	for _, batteryBanks := range data {
		nl := 0
		nr := 0
		nmax, imax := int(batteryBanks[0]-'0'), 0
		for i, c := range batteryBanks {
			n := int(c - '0')
			if n > nmax {
				nl = nmax
				nmax = n
				imax = i
				nr = 0
			} else if i != imax && n > nr {
				nr = n
			}
		}
		if nr == 0 {
			nr = nmax
		} else {
			nl = nmax
		}
		result.jolts += nl * 10 + nr
	}
	return
}


func SolvePart2(data []string) (result Result) {

	return
}