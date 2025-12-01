package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2024/day/6

const fileName = "input.txt"

var labMap = [][]rune{}

type Position struct {
	x int
	y int
	dir Direction
}

type Direction struct {
	x int
	y int
	face string
}


var (
	Top			= Direction {x: -1, y: 0, face: "⮝"}
	Left		= Direction {x: 0,  y: -1, face: "⮜"}
	Right		= Direction {x: 0,  y: 1, face: "⮞"}
	Bottom		= Direction {x: 1,  y: 0, face: "⮟"}
)

const (
	OBSTACLE = '#'
	PATH = '.'
	VISITED = 'X'
	GUARD = '^'
)

func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	prepData := strings.Split(string(data), "\r\n")

	part1(prepData)
	fmt.Println()
	part2(prepData)
}

func dfs(l int, c int, newSteps int, guardDir Direction) int {
	increment := 0
	if IsOutOfBounds(l, c) {
		return newSteps
	}
	if isObstacle(l, c) {
		return dfs(l - guardDir.x, c - guardDir.y, newSteps, TurnRight(guardDir))
	}
	if labMap[l][c] != VISITED {
		labMap[l][c] = VISITED
		increment = 1
	}
	return dfs(l + guardDir.x, c + guardDir.y, newSteps + increment, guardDir)
}


func IsOutOfBounds(l int, c int) bool {
	return l >= len(labMap) || l < 0 || c >= len(labMap[l]) || c < 0
}

func isObstacle(l int, c int) bool {
	return labMap[l][c] == OBSTACLE
}

func TurnRight(dir Direction) Direction {
	switch dir {
		case Top:
			return Right
		case Right:
			return Bottom
		case Bottom: 
			return Left
		case Left:
			return Top
		default:
			panic("idk where you going")
	}
}

func part1(prepData []string) {
	gx, gy := findInitialGuardPos(prepData)

	result := dfs(gx, gy, 0, Top)

	// printMap()

	fmt.Printf("Part1: %d\n", result)
}

func printMap(gx int, gy int, dir Direction) {
	for l, line := range labMap {
		for c, char := range line {
			if l == gx && c == gy {
				fmt.Printf(dir.face)
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func printMapLoop(gx int, gy int, nx int, ny int, dir Direction) {
	
	for l, line := range labMap {
		for c, char := range line {
			if l == nx && c == ny && l == gx && c == gy {
				fmt.Print("X")
			} else if l == gx && c == gy {
				fmt.Print(dir.face)
			} else if l == nx && c == ny {
				fmt.Print("O")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func part2(prepData []string) {
	gx, gy := findInitialGuardPos(prepData)

	result3 := findLoopObstacles(gx, gy, Top)
	
	fmt.Printf("Part2.3: %d\n", result3)
}

func findLoopObstacles(l, c int, d Direction) int {

	nLoops := 0
	memWalked := make(map[Position]bool)
	currPos := Position{x: l, y: c, dir: d}
	isWithinBounds := true
	for isWithinBounds {

		obsPos := Position{x: currPos.x + currPos.dir.x, y: currPos.y + currPos.dir.y}
		_, hasWalked := memWalked[newPos(obsPos.x, obsPos.y)]

		if !hasWalked && !isOutOfBounds2(obsPos) && !isObstacle2(obsPos) && isLoop3(currPos, obsPos) {
			nLoops += 1
		}
		memWalked[newPos(currPos.x, currPos.y)] = true
		currPos, isWithinBounds = move(currPos)
	}

	return nLoops
}

func move(pos Position) (Position, bool) {
	return moveOnce(pos, pos)
}
func moveOnce(pos, obs Position) (Position, bool) {
	peek := newPosDir(pos.x + pos.dir.x, pos.y + pos.dir.y, pos.dir)
	if isObstacle2(peek) || obs.equalsXY(peek) {
		pos = newPosDir(pos.x, pos.y, TurnRight(pos.dir))
	} else {
		pos = peek
	}
	return pos, !isOutOfBounds2(pos)
}

func isObstacle2(pos Position) bool {
	return !isOutOfBounds2(pos) && labMap[pos.x][pos.y] == OBSTACLE
}

func isOutOfBounds2(pos Position) bool {
	return pos.x < 0 || pos.x >= len(labMap) || pos.y < 0 || pos.y >= len(labMap[pos.x])
}

func isLoop3(pos, obsPos Position) bool {
	walked := make(map[Position]bool)

	currPos := newPosDir(pos.x, pos.y, pos.dir)
	for  {
		if _, ok := walked[currPos]; ok {
			return true
		}
		walked[currPos] = true
		
		nPos, isWithinBounds := moveOnce(currPos, obsPos)
		
		if !isWithinBounds {
			break
		} 
	
		currPos = nPos
	}

	return false
}

func newPosDir(x,y int, d Direction) Position {
	return Position{x: x, y: y, dir: d}
}

func newPos(x,y int) Position {
	return Position{x: x, y: y}
}

func findInitialGuardPos(prepData []string) (int, int) {
	gx, gy := -1 , -1
	labMap = make([][]rune, 0)
	for l, line := range prepData {
		labMap = append(labMap, []rune(line))
		if gx != -1 {
			continue
		}
		for c, char := range line {
			if char == GUARD {
				gx = l
				gy = c
				labMap[l][c] = PATH
				break
			}
		}
	}
	return gx, gy
}

func (p *Position) equalsXY(o Position) bool {
	return p.x == o.x && p.y == o.y
}