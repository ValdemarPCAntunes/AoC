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
		n1, i1 := 0, -1
		for i, c := range batteryBanks {
			n := int(c - '0')
			if n > n1 {
				n1 = n
				i1 = i
			}
		}
		n2, i2 := 0, -1
		l, r := i1 + 1, len(batteryBanks)
		if l == len(batteryBanks) {
			r = i1
			l = 0
		} 
		for i := l; i < r; i++ {
			n := int(batteryBanks[i] - '0')
			if n > n2 {
				n2 = n
				i2 = i
			}
		}
		if i2 < i1 {
			n1, n2 = n2, n1
		}
		result.jolts += n1 * 10 + n2
	}
	return
}


func SolvePart2(data []string) (result Result) {

	return
}