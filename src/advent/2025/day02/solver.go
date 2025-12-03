package day02

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/02

const fileName = "input.txt"

func Solve(part1, part2 bool) (r1, r2 Result) {
	data, err := os.ReadFile(fileName)
	elemSep := ","
	if err != nil {
		log.Fatal(err.Error())
	}
	if part1 {
		r1 = SolvePart1(strings.Split(string(data), elemSep))
		fmt.Printf("Part 1 Answer: %d\n", r1.invalids)
	}
	if part2 {
		r2 = SolvePart2(strings.Split(string(data), elemSep))
		fmt.Printf("Part 2 Answer: %d\n", r2.invalids)
	}
	return
}

type Result struct {
	invalids int
	n   int
	skipped []int
}

func SolvePart1(data []string) (result Result) {
	result.invalids = 0
	for _, line := range data {
		pair := strings.Split(line, "-")
		firstId, secondId := pair[0], pair[1]
		if len(firstId) == len(secondId) && len(firstId) & 1 == 1 {
			continue
		}
		upperhalf1, _ := strconv.Atoi(firstId[:len(firstId)/2])
		upperhalf2, _ := strconv.Atoi(secondId[:(len(secondId)+1)/2])
		firstLimit, _ := strconv.Atoi(firstId)
		secondLimit, _ := strconv.Atoi(secondId)

		if upperhalf2 == upperhalf1 {
			bottomhalf1, _ := strconv.Atoi(firstId[len(firstId)/2:])
			bottomhalf2, _ := strconv.Atoi(secondId[len(secondId)/2:])
			if upperhalf1 >= bottomhalf1 && upperhalf1 <= bottomhalf2 {
				result.invalids += upperhalf1 * int(math.Pow10(len(strconv.Itoa(upperhalf1)))) + upperhalf1
			}
			continue
		}

		for i := range(upperhalf2 - upperhalf1 + 1) {
			dupe := upperhalf1 + i
			invalidCandidate := dupe * int(math.Pow10(len(strconv.Itoa(dupe)))) + dupe
			if invalidCandidate < firstLimit || invalidCandidate > secondLimit {
				continue
			}
			result.invalids += invalidCandidate
		}
	}
	return
}


func SolvePart2(data []string) (result Result) {
	result.invalids = 0
	for _, line := range data {
		pair := strings.Split(line, "-")
		firstId, secondId := pair[0], pair[1]
		
		upperhalf1, _ := strconv.Atoi(firstId[:len(firstId)/2])
		upperhalf2, _ := strconv.Atoi(secondId[:(len(secondId)+1)/2])
		firstLimit, _ := strconv.Atoi(firstId)
		secondLimit, _ := strconv.Atoi(secondId)

		mem := make(map[string]bool)
		for i := range(upperhalf2 - upperhalf1 + 1) {
			start := strconv.Itoa(upperhalf1 + i)
			for j := 1; j < len(start) + 1; j++ {
				seq := start[:j]
				//check first limit
				if len(firstId) % len(seq) == 0  {
					seq1 := seq + strings.Repeat(seq, len(firstId) / len(seq) - 1)
					n1, _ := strconv.Atoi(seq1)
					if _, ok := mem[seq1]; !ok && n1 >= firstLimit && n1 <= secondLimit {
						result.invalids += n1
						mem[seq1] = true
					}
				}
				if len(firstId) == len(secondId) {
					continue
				}
				//check second limit if lengths are different
				if len(secondId) % len(seq) == 0  {
					seq2 := seq + strings.Repeat(seq, len(secondId) / len(seq) - 1)
					n2, _ := strconv.Atoi(seq2)
					if _, ok := mem[seq2]; !ok && n2 >= firstLimit && n2 <= secondLimit {
						result.invalids += n2
						mem[seq2] = true
					}
				}
			}
		}
	}
	return
}