// https://adventofcode.com/2024/day/1

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const fileName = "input.txt"

func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	

	lines := strings.Split(string(data), "\r\n") 
	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))
	for i, line := range lines {
		pair := strings.Split(line, "   ")
		val1, _ := strconv.Atoi(pair[0])
		val2, _ := strconv.Atoi(pair[1])
		list1[i] = val1
		list2[i] = val2
	}
	sort.Ints(list1)
	sort.Ints(list2)

	part1(list1, list2)
	fmt.Println()
	part2(list1, list2)
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(list1 []int, list2 []int) {
	result := 0
	for i := 0; i < len(list1); i++ {
		result += abs(list1[i] - list2[i])
	}
	fmt.Printf("Part1: %d", result)
}

func part2(list1 []int, list2 []int) {
	result := 0
	for i1, i2 := 0, 0; i1 < len(list1) && i2 < len(list2);  {
		v1, v2 := list1[i1], list2[i2]
		if v1 == v2{
			result += v1
			i2 += 1
		} else if v1 > v2 {
			i2 += 1
		} else {
			i1 += 1
		}
	}
	fmt.Printf("Part2: %d", result)
}