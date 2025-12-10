package day8

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/8

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
	circuits int
	lastJunctionXmul int
}

type Point struct {
	x, y, z float64
}

type Junction struct {
	//lowest record distance
	dist float64
	//pair
	a, b int
}

func (a Point) CalcEuclideanDistTo(b Point) float64 {
	x, y, z := (a.x-b.x), (a.y-b.y), (a.z-b.z)
	return math.Sqrt( x*x + y*y + z*z )
}

func find_set(set []int, x int) int {
	if set[x] != x {
		set[x] = find_set(set, set[x])
	}
	return set[x]
}

func union_set(set []int, x, y int) {
	x1,y2 := find_set(set, x), find_set(set, y)
	if x1 == y2 {
		return
	}
	set[x1] = y2
}

func is_same_set(set []int) bool {
	cmp := find_set(set, 0)
	for i := range set {
		if find_set(set, i) != cmp {
			return false
		}
	}
	return true
}

func SolvePart1(data []string) (result Result) {
	length := len(data)
	boxes := make([]Point, length)
	
	for i := range data {
		points := strings.Split(data[i], ",")
		p := Point{}
		p.x, _ = strconv.ParseFloat(points[0], 32)
		p.y, _ = strconv.ParseFloat(points[1], 32)
		p.z, _ = strconv.ParseFloat(points[2], 32)
		boxes[i] = p
	}
	
	dists := make([]Junction, 0, length)
	for i := range boxes {
		for j := i + 1; j < length; j++ {
			dist := boxes[i].CalcEuclideanDistTo(boxes[j]) 
			dists = append(dists, Junction{
				a: i,
				b: j,
				dist: dist,
			})
		}
	}
	
	slices.SortFunc(dists, func (i, j Junction) int {
		return int(math.Floor(i.dist - j.dist))
	})
	
	//explicit disjoint union set init
	circuits := make([]int, length)
	for i := range length {
		circuits[i] = i
	}

	for i := range length {
		p := dists[i]
		union_set(circuits, p.a, p.b)
	}

	counts := make([]int, length)

	for i := range circuits {
		counts[find_set(circuits, i)]++
	}

	maxmax, max2, max3 := 0, 0, 0
    for _, count := range counts {
        if count > maxmax {
            max3, max2, maxmax = max2, maxmax, count
        } else if count > max2 {
            max3, max2 = max2, count
        } else if count > max3 {
            max3 = count
        }
    }
	result.circuits = maxmax * max2 * max3
	
	return
}


func SolvePart2(data []string) (result Result) {
	length := len(data)
	boxes := make([]Point, length)
	
	for i := range data {
		points := strings.Split(data[i], ",")
		p := Point{}
		p.x, _ = strconv.ParseFloat(points[0], 32)
		p.y, _ = strconv.ParseFloat(points[1], 32)
		p.z, _ = strconv.ParseFloat(points[2], 32)
		boxes[i] = p
	}
	
	dists := make([]Junction, 0, length)
	for i := range boxes {
		for j := i + 1; j < length; j++ {
			dist := boxes[i].CalcEuclideanDistTo(boxes[j]) 
			dists = append(dists, Junction{
				a: i,
				b: j,
				dist: dist,
			})
		}
	}
	
	slices.SortFunc(dists, func (i, j Junction) int {
		return int(math.Floor(i.dist - j.dist))
	})
	
	//explicit disjoint union set init
	circuits := make([]int, length)
	for i := range length {
		circuits[i] = i
	}
	i := 0
	for ; i < len(dists); i++ {
		p := dists[i]
		union_set(circuits, p.a, p.b)
		if is_same_set(circuits) {
			break
		}
	}

	result.lastJunctionXmul = int(boxes[dists[i].a].x) * int(boxes[dists[i].b].x)
	
	return
}

