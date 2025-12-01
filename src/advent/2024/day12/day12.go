package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2024/day/12

const fileName = "input.txt"


func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	prepData := strings.Split(string(data), "\r\n\r\n")
	for i, test := range prepData {
		fmt.Printf("[TEST] %d\n", i+1)
		testData :=	strings.Split(test, "\r\n")
		answer1 := part1(testData[2:])
		fmt.Printf("Part1 E: %s\tA: %d\n", testData[0], answer1)
		answer2 := part2(testData[2:])
		fmt.Printf("Part2 E: %s\tA: %d\n", testData[1], answer2)
		fmt.Println()
	}
}


func part1(data []string) int {
	mem := make([][]bool, len(data))
	for i, line := range data {
		mem[i] = make([]bool, len(line)) 		
	}
	

	var isOutOfBounds = func(x, y int) bool {
		return x < 0 || y < 0 || x >= len(data) || y >= len(data[x]) 
	}

	var dfs func(x, y int, c rune) (area, perimeter int)
	dfs = func(x, y int, c rune) (area, perimeter int) {
		if isOutOfBounds(x, y) || c != rune(data[x][y]) {
			return 0, 1
		}
		if mem[x][y] {
			return 0, 0
		}
		mem[x][y] = true
		ta, tp := dfs(x - 1, y, c)
		la, lp := dfs(x, y - 1, c)
		ra, rp := dfs(x, y + 1, c)
		ba, bp := dfs(x + 1, y, c)
		return 1 + ta + la + ra + ba, tp + lp + rp + bp
	}

	price := 0
	for x, line := range data {
		for y, c := range line {
			if mem[x][y] {
				continue
			}
			a, p := dfs(x, y, c)
			price += a * p
		}
	}
	return price
}


type Pos struct {
	x, y int
}

func part2(data []string) int {
	mem := make([][]bool, len(data))
	for i, line := range data {
		mem[i] = make([]bool, len(line)) 		
	}
	

	var isOutOfBounds = func(x, y int) bool {
		return x < 0 || y < 0 || x >= len(data) || y >= len(data[x]) 
	}

	var dfs func(x, y int, c rune, region map[Pos]bool) (area int)
	dfs = func(x, y int, c rune, region map[Pos]bool) (area int) {
		if isOutOfBounds(x, y) || c != rune(data[x][y]) {
			return 0
		}
		if mem[x][y] {
			return 0
		}
		mem[x][y] = true
		region[Pos{x: x, y: y}] = true
		ta := dfs(x - 1, y, c, region)
		la := dfs(x, y - 1, c, region)
		ra := dfs(x, y + 1, c, region)
		ba := dfs(x + 1, y, c, region)
		return 1 + ta + la + ra + ba
	}

	price := 0
	for x, line := range data {
		for y, c := range line {
			if mem[x][y] {
				continue
			}
			region := make(map[Pos]bool)
			area := dfs(x, y, c, region)
			corners := getCorners(region)
			price += area * corners
		}
	}
	return price
}

func getCorners(region map[Pos]bool) (corners int) {
	for k := range region {
		_, t := region[Pos{x: k.x - 1, y: k.y}]
		_, r := region[Pos{x: k.x, y: k.y + 1}]
		_, l := region[Pos{x: k.x, y: k.y - 1}]
		_, b := region[Pos{x: k.x + 1, y: k.y}]
		_, dtl := region[Pos{x: k.x - 1, y: k.y - 1}]
		_, dtr := region[Pos{x: k.x - 1, y: k.y + 1}]
		_, dbl := region[Pos{x: k.x + 1, y: k.y - 1}]
		_, dbr := region[Pos{x: k.x + 1, y: k.y + 1}]
        //normal corners
        corners += btoi(!t && !l) + btoi(!t && !r) + btoi(!b && !l) + btoi(!b && !r)
        //diag corners
        corners += btoi(b && r && !dbr) +  btoi(b && l && !dbl) + btoi(t && r && !dtr) + btoi(t && l && !dtl)
        }
        return 
}

func btoi(b bool) int {
        if b {
                return 1
        }
        return 0
}

// ---------------------------- AREA OF SHAME ---------------------------
// func part1(data []string) int {
// 	mem := map[rune]Perimeter{}
// 	var isOutOfBoundsOrDifferent = func(x, y int, c rune) int {
// 		if x < 0 || y < 0 || x >= len(data) || y >= len(data[x]) || rune(data[x][y]) != c {
// 			return 1
// 		}
// 		return 0
// 	}
// 	for i, row := range data {
// 		for j, c := range row {
// 			if _, ok := mem[c]; !ok {
// 				mem[c] = Perimeter{}
// 			}
// 			p := mem[c]
// 			top := isOutOfBoundsOrDifferent(i - 1, j, c)
// 			left := isOutOfBoundsOrDifferent(i, j - 1, c)
// 			right := isOutOfBoundsOrDifferent(i, j + 1, c)
// 			bottom := isOutOfBoundsOrDifferent(i + 1, j, c)
// 			p.perimeter += top + left + right + bottom
// 			p.plants += 1
// 			mem[c] = p
// 		} 
// 	}
// 	result := 0
// 	for _, v := range mem {
// 		result += v.perimeter * v.plants
// 	}
// 	return result
// }