package main

// https://adventofcode.com/2024/day/2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "input.txt"
const maxDiff = 3
const minDiff = 1
const tolerance = 1

func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	part1(data)
	fmt.Println()
	part2(data)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(data []byte) {
	reports := strings.Split(string(data), "\r\n")
	result := 0
	for _, report := range reports {
		if isReportSafe(strings.Fields(report), 0) {
			result += 1
		}
	}
	fmt.Printf("Part1: %d", result)
}

func isReportSafe(report [] string, allowedFailures int ) bool {
	val1, _ := strconv.Atoi(report[0])
	val2, _ := strconv.Atoi(report[1])
	isIncreasing :=  val2 > val1
	for i := 1; i < len(report); i++ {
		val1, _ = strconv.Atoi(report[i-1])
		val2, _ = strconv.Atoi(report[i])
		if(!isSafeDistance(val1, val2, isIncreasing)) {
			if(allowedFailures > 0) {
				allowedFailures -= 1
				continue
			}
			return false
		}
	}
	return true
}

func isSafeDistance(val1 int, val2 int, isIncreasing bool) bool {
	diff := val2 - val1
	if abs(diff) > maxDiff || abs(diff) < minDiff {
		return false
	}
	if isIncreasing && diff < 0 {
		return false
	} 
	if !isIncreasing && diff > 0 {
		return false
	}
	return true
}

func part2(data []byte) {
	reports := strings.Split(string(data), "\r\n")
	result := 0
	for _, report := range reports {
		
		if isReportSafe(strings.Fields(report), tolerance) {
			result += 1
		}
	}

	fmt.Printf("Part2: %d", result)
}