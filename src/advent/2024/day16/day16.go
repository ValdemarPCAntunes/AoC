package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2024/day/16

const fileName = "input.txt"


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



func part1(data []string) {
	result := 0

	fmt.Printf("Part1: %d", result)
}


func part2(data []string) {
	result := 0

	fmt.Printf("Part2: %d", result)
}