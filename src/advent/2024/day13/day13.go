package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2024/day/13

const fileName = "input.txt"


func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	prepData := strings.Split(string(data), "\r\nNEXT TEST\r\n")
	btnA := 0
	btnB := 1
	goal := 2
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

type Position struct {
	x,y int
}

func part1(data []string) (result int) {

	return
}


func part2(data []string) (result int) {

	return
}
