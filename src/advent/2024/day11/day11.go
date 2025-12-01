package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// https://adventofcode.com/2024/day/11

const fileName = "input.txt"


func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	prepData := strings.Split(string(data), " ")

	// part1(prepData)
	// fmt.Println()
	part2(prepData, 250)
	fmt.Println()
	part2(prepData, 750)
	fmt.Println()
	part2v2(prepData, 250)
	fmt.Println()
	part2v2(prepData, 750)
}

type Stone struct {
	val int
	next *Stone
}

func printList(stoneHead *Stone) {
	curr := stoneHead
	for curr != nil {
		fmt.Printf("%d ", curr.val)
		curr = curr.next
	}
	fmt.Println()
}

func part1(data []string) {
	stoneHead := &Stone{}

	curr := stoneHead
	for i, rune := range data {
		num, _ := strconv.Atoi(rune)
		curr.val = num
		if i < len(data)-1 {
			curr.next = &Stone{}
			curr = curr.next
		}
	}
	
	stones := len(data)
	blinks := 25
	pStartTime := time.Now()
	for ; blinks > 0; blinks -=1 {
		curr := stoneHead
		for curr != nil {
			stringVal := fmt.Sprintf("%d", curr.val)
			if curr.val == 0 {
				curr.val = 1
			} else if len(stringVal) & 1 == 0 {
				leftVal,_ := strconv.Atoi(stringVal[:len(stringVal)/2])
				rightVal,_ := strconv.Atoi(stringVal[len(stringVal)/2:])
				newRightStone := &Stone{val: rightVal, next: curr.next}
				curr.val = leftVal
				curr.next = newRightStone
				stones += 1
				curr = curr.next
			} else {
				curr.val *= 2024
			}
			curr = curr.next
		}
	}
	fmt.Printf("Part1: %d stones; %d iterations; Time [%s]\n",stones, 25, time.Since(pStartTime))
}

func part2v2(data []string, blinks int) {
	mapin := map[int]int{}
	mapout := map[int]int{}
	startTime := time.Now()
	for _, rune := range data {
		num,_ := strconv.Atoi(rune)
		mapin[num] += 1
	}
	for blinks > 0 {
		for k, v := range mapin {
			stringVal := fmt.Sprintf("%d", k)
			if k == 0 {
				mapout[1] += v
			} else if len(stringVal) & 1 == 0 {
				leftVal,_ := strconv.Atoi(stringVal[:len(stringVal)/2])
				rightVal,_ := strconv.Atoi(stringVal[len(stringVal)/2:])
				mapout[leftVal] += v
				mapout[rightVal] += v
			} else {
				mapout[k * 2024] += v
			}
			mapin[k] -= v
		}
		tmp := mapin
		mapin = mapout
		mapout = tmp
		blinks -= 1
	}

	stones := 0
	for _, v := range mapin {
		stones += v
	}

	fmt.Printf("Part2v2: %d stones; iterations %d; Time taken [%s]", stones, blinks, time.Since(startTime))
}

func part2(data []string, blinks int) {
	pStartTime := time.Now()
	totalCreated := solve(data, blinks)
	fmt.Printf("Part2: %d stones; iterations %d; Time taken [%s]\n", totalCreated, blinks, time.Since(pStartTime))
}

func solve(data []string, blinksCount int) int {
	cache := map[int]int{}
	
	sum := 0

	for _, xs := range data {
		x,_ := strconv.Atoi(xs)
		sum += blinkRec(x, blinksCount, cache)
	}
	return sum
}

func blinkRec(x, n int, cache map[int]int) (result int) {
	cacheKey := x*100 + n 
	if v, ok := cache[cacheKey]; ok {
		return v
	}
	defer func() {
		cache[cacheKey] = result
	}()
	if n == 0 {
		return 1
	}
	if x == 0 {
		return blinkRec(1, n-1, cache)
	}
	sx := strconv.Itoa(x)
	if len(sx)&1 == 1 {
		return blinkRec(x*2024, n-1, cache)
	}
	half := len(sx) / 2
	left,_ := strconv.Atoi(sx[:half])
	right,_ := strconv.Atoi(sx[half:])
	return blinkRec(left, n-1, cache) + blinkRec(right, n-1, cache)
}

