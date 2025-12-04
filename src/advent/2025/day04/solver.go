package day4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2025/day/4

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
		fmt.Printf("Part 1 Answer: %d\n", r1)
	}
	if part2 {
		r2 = SolvePart2(preparedData)
		fmt.Printf("Part 2 Answer: %d\n", r2)
	}
	return
}

type Result struct {
	paper_rolls int
}

func SolvePart1(data []string) (result Result) {

	return
}


func SolvePart2(data []string) (result Result) {

	return
}