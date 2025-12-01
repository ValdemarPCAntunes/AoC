package main

// https://adventofcode.com/2024/day/4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const fileName = "input.txt"

const (
    X = 'X'
    M = 'M'
    A = 'A'
    S = 'S'
)

var xmasSequence = []rune{X , M, A, S}

var puzzle = [][]rune{}

type Direction struct {
	x int
	y int
}

var (
	TopLeft		= Direction {x: -1, y: -1}
	Top			= Direction {x: -1, y: 0}
	TopRight	= Direction {x: -1, y: 1}
	Left		= Direction {x: 0,  y: -1}
	Right		= Direction {x: 0,  y: 1}
	BottomLeft	= Direction {x: 1,  y: -1}
	Bottom		= Direction {x: 1,  y: 0}
	BottomRight	= Direction {x: 1,  y: 1}
)

func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	prepData := strings.Split(string(data), "\r\n")


	for _, line := range prepData {
		puzzle = append(puzzle, []rune(line))
	}

	part1()
	fmt.Println()
	part2()
}

func part1() {
	result := 0
	for l := 0; l < len(puzzle); l += 1 {
		for c := 0; c < len(puzzle[l]); c += 1 {
			result += tryMatchXmas(l, c)
		}
	}

	fmt.Printf("Part1: %d", result)
}

func tryMatchXmas(l int, c int) int {
	firstMatch := xmasSequence[0]
	startsWithX := puzzle[l][c] == firstMatch
	if(!startsWithX) {
		return 0
	}
	
	return 	traverse(l, c, TopLeft) + traverse(l, c, Top) + traverse(l, c, TopRight) +
			traverse(l, c, Left) + traverse(l, c, Right) +
			traverse(l, c, BottomLeft) + traverse(l, c, Bottom) + traverse(l, c, BottomRight)
}

func traverse(l int, c int, dir Direction) int {
	return traverseDfs(l, c, dir, 1)
}

func traverseDfs(l int, c int, dir Direction, counter int) int {
	l = l + dir.x
	c = c + dir.y
	if (!isWithinBounds(l,c)) {
		return 0
	}
	isXmas := xmasSequence[counter] == puzzle[l][c]
	counter += 1
	if (isXmas && ( counter == len(xmasSequence) || traverseDfs(l, c, dir, counter) == 1 )) {
		return 1
	} 
	return 0
}

func isWithinBounds(l int, c int) bool {
	return l >= 0 && len(puzzle) > l && c >= 0 && len(puzzle[l]) > c
}

func part2() {
	result := 0
	for l := 1; l < len(puzzle) - 1; l += 1 {
		for c := 1; c < len(puzzle[l]) - 1 ; c += 1 {
			if(puzzle[l][c] != A) {
				continue
			}
			if (isXShapeMAS(l, c)) {
				result += 1
			}
		}
	}

	fmt.Printf("Part2: %d", result)
}

func isXShapeMAS(l int, c int) bool {
	return isBackDiagonalMAS(l, c) && isForwardDiagonalMAS(l, c)
}

func isBackDiagonalMAS(l int, c int) bool {
	return  puzzle[l-1][c-1] == M && puzzle[l+1][c+1] == S  ||
			puzzle[l-1][c-1] == S && puzzle[l+1][c+1] == M 
}

func isForwardDiagonalMAS(l int, c int) bool {
	return  puzzle[l+1][c-1] == M && puzzle[l-1][c+1] == S ||
			puzzle[l+1][c-1] == S && puzzle[l-1][c+1] == M
}