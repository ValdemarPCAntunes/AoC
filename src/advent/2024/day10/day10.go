package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

// https://adventofcode.com/2024/day/10

const fileName = "input.txt"

const (
	Score = 9
	Head = 0
)

func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	prepData := strings.Split(string(data), "\r\n\r\n")
	p1ans := []int{2,2,4,36,6,10,4,2,2,582}
	p2ans := []int{6,2,13,81,6,10,4,2,227,1302}
	for i, test := range prepData {
		fmt.Printf("[TEST] %d\n", i+1)
		testData := toIntDoubleArray(strings.Split(test, "\r\n"))
		answer1 := part1(testData)
		fmt.Printf("Part1 E: %d\tA: %d\n", p1ans[i], answer1)
		answer2 := part2(testData)
		fmt.Printf("Part2 E: %d\tA: %d\n", p2ans[i], answer2)
		fmt.Println()
	}
}

func toIntDoubleArray(data [] string) [][]int {
	matrix := make([][]int, len(data))
	for i, row := range data {
		matrix[i] = make([]int, len(row))
		for j, r := range row {
			matrix[i][j] = int(r - '0')
		}
	}
	return matrix
}

func part1(data [][]int) int {
	result := 0
	for i := 0; i < len(data); i += 1 {
		for j := 0; j < len(data[i]); j += 1 {
			if data[i][j] == Head {
				trailMem := make([][]int, len(data))
				for i, line := range data {
					trailMem[i] = make([]int, len(line))
				}
				result += tryCountTrail(data, i, j, Head, trailMem)
			}
		}
	}
	return result
}

func tryCountTrail(data [][]int, i, j, curr int, trailMem [][]int) int {
	top := countTrails(data, i - 1, j, curr, trailMem)
	left := countTrails(data, i, j - 1, curr, trailMem)
	right := countTrails(data, i, j + 1, curr, trailMem)
	bottom := countTrails(data, i + 1, j, curr, trailMem)
	trailMem[i][j] = top + left + right + bottom
	return top + left + right + bottom
}

func countTrails(data [][]int, i, j, prev int, trailMem [][]int) int {
	if isOutOfBounds(data, i, j) {
		return 0
	}
	curr := data[i][j]
	if curr - prev != 1 || trailMem[i][j] != 0{
		return 0
	}
	if curr == Score { 
		trailMem[i][j] = 1
		return 1
	}
	return tryCountTrail(data, i, j, curr, trailMem)
}

func tryFindTrail(data [][]int, i, j, curr int, trailMem [][]int) int {
	top := findTrail(data, i - 1, j, curr, trailMem)
	left := findTrail(data, i, j - 1, curr, trailMem)
	right := findTrail(data, i, j + 1, curr, trailMem)
	bottom := findTrail(data, i + 1, j, curr, trailMem)
	trails := top + left + right + bottom
	trailMem[i][j] = trails
	return trails
}

func findTrail(data [][]int, i, j, prev int, trailMem [][]int) int {
	if isOutOfBounds(data, i, j) {
		return 0
	}
	curr := data[i][j]
	if curr - prev != 1 {
		return 0
	}
	if curr == Score {
		return 1
	}
	if nTrails, ok := alreadyFoundTrail(i, j, trailMem); ok {
		trailMem[i][j] = nTrails
		return nTrails
	}
	return tryFindTrail(data, i, j, curr, trailMem)
}

func alreadyFoundTrail(i, j int, trailMem [][]int) (int, bool) {
	nTrails := trailMem[i][j]
	return nTrails, nTrails != -1
}

func isOutOfBounds(data [][]int, i, j int) bool {
	return i < 0 || j < 0 || i >= len(data) || j >= len(data[i])
}


func part2(data [][]int) int {
	result := 0
	trailMem := make([][]int, len(data))
	for i, line := range data {
		trailMem[i] = slices.Repeat([]int{-1}, len(line))
	}
	for i := 0; i < len(data); i += 1 {
		for j := 0; j < len(data[i]); j += 1 {
			if data[i][j] == Head {
				trails := tryFindTrail(data, i, j, Head, trailMem)
				result += trails
			}
		}
	}
	return result
}

