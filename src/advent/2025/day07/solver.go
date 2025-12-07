package day7

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2025/day/7

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
	splits int
	timelines int
}

type Pair struct {
	x, y int
}

const (
	SPLITTER = '^'
	START = 'S'
	EMPTY = '.'
)

func SolvePart1(data []string) (result Result) {
	line := make([]bool, len(data))
	for i, c := range data[0] {
		if c == START {
			line[i] = true
			break
		}
	}
	
	for _, str := range data[1:] {
		for i, c :=  range str {
			if c == EMPTY {
				continue
			}
			if line[i] && c == SPLITTER {
				result.splits++
				line[i] = false
				line[i-1] = true
				line[i+1] = true
			}
		}
	}
	return
}


func SolvePart2(data []string) (result Result) {
	y := 0
	for i, c := range data[0] {
		if c == START {
			y = i
			break
		}
	}
	mem := make(map[Pair]int, 0)

	result.timelines = dfs(1, y, data, mem)
	return
}

func dfs(x,y int, data []string, mem map[Pair]int) (timelines int) {
	if x == len(data) {
		timelines += 1
		return 
	}
	p := Pair{x: x, y: y}
	if _, ok := mem[p]; ok {
		timelines += mem[p]
		return
	} 
	if data[x][y] == SPLITTER {
		timelines += dfs(x+1, y-1, data, mem) + dfs(x+1, y+1, data, mem)
	} else {
		timelines += dfs(x+1, y, data, mem)
	}
	mem[p] = timelines
	return
}
