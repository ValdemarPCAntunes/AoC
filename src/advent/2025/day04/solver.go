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

const paper = '@'
const empty_space = '.'
const forklift_maximum_size = 4

func SolvePart1(data []string) (result Result) {
	isPaper := func(i, j int)int {
		if  i >= 0 && i < len(data) && j >= 0 && j < len(data[i]) && data[i][j] == paper {
			return 1
		}
		return 0
	}
	for i := range data {
		for j := range data[i] {
			if data[i][j] != paper {
				continue
			}
			pc := isPaper(i - 1, j - 1) 	// Top Left
			pc += isPaper(i - 1, j) 		// Top
			pc += isPaper(i - 1, j + 1) 	// Top Right
			pc += isPaper(i, j - 1) 		// Left
			pc += isPaper(i, j + 1) 		// Right
			pc += isPaper(i + 1, j - 1) 	// Bottom Left
			pc += isPaper(i + 1, j) 		// Bot
			pc += isPaper(i + 1, j + 1) 	// Bottom Right
			if pc < forklift_maximum_size {
				result.paper_rolls++
			}
		}
	}
	return
}


func SolvePart2(data []string) (result Result) {
	runeData := make([][]rune, len(data))
	for i := range runeData {
		runeData[i] = []rune(data[i])
	}
	isPaper := func(i, j int)int {
		if  i >= 0 && i < len(runeData) && j >= 0 && j < len(runeData[i]) && runeData[i][j] == paper {
			return 1
		}
		return 0
	}
	done := false
	for !done {
		done = true
		for i := range runeData {
			for j := range runeData[i] {
				if runeData[i][j] != paper {
					continue
				}
				pc := isPaper(i - 1, j - 1) 	// Top Left
				pc += isPaper(i - 1, j) 		// Top
				pc += isPaper(i - 1, j + 1) 	// Top Right
				pc += isPaper(i, j - 1) 		// Left
				pc += isPaper(i, j + 1) 		// Right
				pc += isPaper(i + 1, j - 1) 	// Bottom Left
				pc += isPaper(i + 1, j) 		// Bot
				pc += isPaper(i + 1, j + 1) 	// Bottom Right
				if pc < forklift_maximum_size {
					result.paper_rolls++
					runeData[i][j] = empty_space
					done = false
				}
			}
		}
	}
	return
}