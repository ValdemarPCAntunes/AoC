package day01

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/1

const fileName = "input.txt"

func Solve(part1, part2 bool) (r1, r2 Result) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	if part1 {
		r1 = SolvePart1(strings.Split(string(data), "\r\n"))
		fmt.Printf("Part 1 Answer: %d\n", r1)
	}
	if part2 {
		r2 = SolvePart2(strings.Split(string(data), "\r\n"))
		fmt.Printf("Part 2 Answer: %d\n", r2)
	}
	return
}

type Result struct {
	dial int
	pwd int
}

const initialPosition = 50
const fullRotation = 100
const leftTurn = 'L'
//const rightTurn = 'R'

func SolvePart1(data []string) (result Result) {
	result.dial = initialPosition
	result.pwd = 0
	for _, line := range data {
		op := rune(line[0])	
		rotation, _ := strconv.Atoi(line[1:])
		if op == leftTurn {
			result.dial = (result.dial - rotation + fullRotation) % fullRotation
		} else {
			result.dial = (result.dial + rotation) % fullRotation
		}
		if result.dial == 0 {
			result.pwd += 1
		}
	}

	return
}


func SolvePart2(data []string) (result Result) {
	result.dial = initialPosition
	result.pwd = 0
	for _, line := range data {
		op := rune(line[0])	
		rotation, _ := strconv.Atoi(line[1:])
		tmp := 0
		cmp := rotation % fullRotation
		if op == leftTurn {
			tmp := result.dial - cmp
			//went below 0
			if result.dial != 0 && tmp <= 0 {
				result.pwd += 1 
			}
			result.dial = (tmp + fullRotation) % fullRotation
		} else {
			tmp = result.dial + cmp
			//went over 100
			if tmp >= fullRotation {
				result.pwd += 1 
			}
			result.dial = tmp % fullRotation
		}
		// count how many revolutions happened
		result.pwd += rotation / fullRotation
	}

	return
}