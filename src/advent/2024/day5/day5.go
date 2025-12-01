package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/5


const fileName = "input.txt"


type RuleSet struct {
	rules map[int][]int
}

var ruleSet = RuleSet{}

func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	prepData := strings.Split(string(data), "\r\n\r\n")
	rules := strings.Split(prepData[0], "\r\n")
	pages := strings.Split(prepData[1], "\r\n")
	ruleSet.rules = make(map[int][]int)
	for _, rule := range rules {
		numPair := strings.Split(rule, "|")
		prevNum, _ := strconv.Atoi(numPair[0])
		afterNum, _ := strconv.Atoi(numPair[1])
		ruleSet.addRule(prevNum, afterNum)
	}
	
	part1(pages)
	fmt.Println()
	part2(pages)
}

func (r *RuleSet) addRule(key, value int) {
	ruleSet.rules[key] = append(ruleSet.rules[key], value)
}

func (r *RuleSet) has(key, value int) bool {
	if rule, ok := r.rules[key]; ok {
		for _, num := range rule {
			if num == value {
				return true
			}
		}
	}
	return false
}

func isPageValid(page []int) bool {
	for i := len(page) - 1; i > 0 ; i -= 1 {
		for j := i-1; j >= 0; j -= 1 {
			if ruleSet.has(page[i], page[j]) {
				return false
			}
		}
	}
	return true
}

func part1(pages []string) {
	result := 0
	for _, page := range pages {
		updates := toInt(strings.Split(page, ","))

		if isPageValid(updates) {
			result += updates[len(updates)/2]
		}
	}

	fmt.Printf("Part1: %d", result)
}

func toInt(data []string) []int {
	ints := []int{}
	for _, update := range data {
		num, _ := strconv.Atoi(update)
		ints = append(ints, num)
	}
	return ints
}

func sortPage(page []int) ([]int, bool) {
	hadToSort := false
	for i := len(page) - 1; i > 0 ; i -= 1 {
		for j := i-1; j >= 0; j -= 1 {
			if ruleSet.has(page[i], page[j]) {
				tmp := page[i]
				page[i] = page[j]  
				page[j] = tmp
				hadToSort = true
			}
		}
	}
	return page, hadToSort
}

func part2(pages []string) {
	result := 0
	for _, page := range pages {
		updates := toInt(strings.Split(page, ","))

		sortedPage, hadToSort := sortPage(updates)
		if hadToSort {
			result += sortedPage[len(sortedPage)/2]
		}
		
	}

	fmt.Printf("Part2: %d", result)
}











