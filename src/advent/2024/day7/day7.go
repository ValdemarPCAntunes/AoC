package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/7

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

func part1(prepData []string) {
	result := 0
	for _, eqstr := range prepData {
		eq := extractIntValues(eqstr)
		if eq.validate() {
			result += eq.testValue
		}
	}
	fmt.Printf("Part1: %d\n", result)
}

func (eq *Equation) validate() bool {
	return eq.dfs(eq.operands[0], 1)
}

func (eq *Equation) dfs(acc int, idx int) bool {
	if idx == len(eq.operands) {
		return eq.testValue == acc
	}
	op := eq.operands[idx];
	return eq.dfs( acc * op, idx + 1) || eq.dfs(acc + op, idx + 1)
}

func extractIntValues(eqstr string) Equation {
	s := strings.Split(eqstr, ": ")
	testValue, _ := strconv.Atoi(s[0])
	operands := []int{}
	for _, opstr := range strings.Split(s[1], " ") {
		op, _ := strconv.Atoi(opstr)
		operands = append(operands, op)
	}
	
	return Equation{testValue: testValue, operands: operands}
}

func part2(prepData []string) {
	result := 0
	for _, eqstr := range prepData {
		eq := extractIntValues(eqstr)
		if eq.validate2() {
			result += eq.testValue
		}
	}
	fmt.Printf("Part2: %d\n", result)
}

func (eq *Equation) validate2() bool {
	return eq.dfs2(eq.operands[0], 1)
}

func (eq *Equation) dfs2(acc int, i int) bool {
	if i == len(eq.operands) {
		return eq.testValue == acc
	}
	op := eq.operands[i];
	return eq.dfs2( acc * op, i + 1) || eq.dfs2(acc + op, i + 1) || eq.dfs2(concat(acc, op), i + 1)
}

func concat(acc int, op int) int {
	newOp, _ := strconv.Atoi(fmt.Sprintf("%d%d", acc, op))
	return newOp
}

type Equation struct {
	testValue int
	operands []int
}