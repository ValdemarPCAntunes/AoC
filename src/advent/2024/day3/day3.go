package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const fileName = "input.txt"

// https://adventofcode.com/2024/day/3


func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	instructions := string(data)

	part1(instructions)
	part2(instructions)
}

func part1(instructions string) {
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)
	validInstructions := re.FindAllString(instructions, -1)
	result := 0

	for _, instruction := range validInstructions {
		numPair := strings.Split(instruction[4:len(instruction)-1], ",")
		num1, _ := strconv.Atoi(numPair[0])
		num2, _ := strconv.Atoi(numPair[1])
		result += num1 * num2
	}
	
	fmt.Printf("Part1: %d\n", result)
}

func part2(instructions string) {
	pattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	
	re := regexp.MustCompile(pattern)

	validInstructions := re.FindAllString(instructions, -1)

	canMul := true

	allowMul := "do()"
	stopMul := "don't()"

	result := 0

	for _, instruction := range validInstructions {
		if(instruction == allowMul) {
			canMul = true
			continue
		} 
		if(instruction == stopMul) {
			canMul = false
			continue
		}
		if(!canMul) {
			continue
		}
		numPair := strings.Split(instruction[4:len(instruction)-1], ",")
		num1, _ := strconv.Atoi(numPair[0])
		num2, _ := strconv.Atoi(numPair[1])
		result += num1 * num2
	}
	
	fmt.Printf("Part2: %d\n", result)
}
