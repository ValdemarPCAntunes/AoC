package day9

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/9

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
	largestArea int
}

type Point struct {
	x, y int
}

func (a Point) CalcArea(b Point) int {
	c, l := math.Abs(float64(a.x - b.x + 1)), math.Abs(float64(a.y - b.y + 1))
	return int(c * l)
}

func SolvePart1(data []string) (result Result) {
	dlength := len(data)
	points := make([]Point, dlength)

	for i, l := range data {
		tuple := strings.Split(l, ",")
		p := Point{}
		p.x, _ = strconv.Atoi(tuple[0])
		p.y, _ = strconv.Atoi(tuple[1])
		points[i] = p
	}

	for i := range dlength {
		for j := i + 1; j  < dlength; j++ {
			area := points[i].CalcArea(points[j])
			if area > result.largestArea {
				result.largestArea = area
			}
		}
	}

	return
}


func SolvePart2(data []string) (result Result) {

	return
}