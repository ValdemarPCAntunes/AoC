package day6

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/6

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
		fmt.Printf("Part 1 Answer: %d\n", r1.result)
	}
	if part2 {
		r2 = SolvePart2(preparedData)
		fmt.Printf("Part 2 Answer: %d\n", r2.result)
	}
	return
}

type Result struct {
	result int
}

const (
	MUL = '*'
	ADD = '+'
	EMPTY = ' '
)


func SolvePart1(data []string) (result Result) {
	ops := make([]int, len(data) - 1)
	for i := 0; i < len(data[len(data)-1]); {
		largestNum := 0
		opidx := 0
		for j := range len(data)-1 {
			startIdx := i
			for ; startIdx < len(data[j]); startIdx++ {
				if data[j][startIdx] == EMPTY {
					continue
				} else {
					break
				}
			}
			k := startIdx
			for  ; k < len(data[j]) ; k++ {
				if data[j][k] == EMPTY {
					size := k - startIdx
					ops[opidx], _ = strconv.Atoi(data[j][startIdx:k])
					opidx++
					if largestNum < size {
						largestNum = size
					}
					break
				}
			}
			//last possible iteration
			if k == len(data[j]) {
				size := k - startIdx
					ops[opidx], _ = strconv.Atoi(data[j][startIdx:k])
					opidx++
					if largestNum < size {
						largestNum = size
					}
			}
		}
		
		subResult := 0
		switch data[len(data)-1][i] {
		case MUL:
			subResult = 1
			for _, o := range ops {
				subResult *= o
			}
		case ADD:
			for _, o := range ops {
				subResult += o
			}
		}
		
		i += largestNum + 1
		
		result.result += subResult
	}
	return
}


func SolvePart2(data []string) (result Result) {

	return
}