package day5

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/5

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
	fresh_ingredients int
}

type Pair struct {
	firstId, secondId int
}

func SolvePart1(data []string) (result Result) {
	i := 0
	ranges := make([]Pair, 0)
	for ; data[i] != ""; i++ {
		dataSplit := strings.Split(data[i], "-")
		n1, _ := strconv.Atoi(dataSplit[0])
		n2, _ := strconv.Atoi(dataSplit[1])
		p := Pair{
			firstId: n1,
			secondId: n2,
		}
		ranges = append(ranges, p)
	}
	
	i++
	for ; i < len(data); i++ {
		id, _ := strconv.Atoi(data[i])
		
		for _, p := range ranges {
			if p.firstId <= id && p.secondId >= id {
				result.fresh_ingredients++
				break
			}
		}
	} 
	return
}


func SolvePart2(data []string) (result Result) {
	ranges := make([]Pair, 0)
	for i := 0; data[i] != ""; i++ {
		dataSplit := strings.Split(data[i], "-")
		n1, _ := strconv.Atoi(dataSplit[0])
		n2, _ := strconv.Atoi(dataSplit[1])
		p := Pair{
			firstId: n1,
			secondId: n2,
		}
		ranges = append(ranges, p)
	}
	slices.SortFunc(ranges, func (p1, p2 Pair) int  {
		if p1.firstId == p2.firstId {
			return p1.secondId - p2.secondId
		}
		return p1.firstId - p2.firstId
	})
	
	mergedRanges := make([]Pair, 0, len(ranges))
	for i := 0; i < len(ranges);{
		j := i + 1
		for ; j < len(ranges) && ranges[j].firstId <= ranges[i].secondId; j++ {
			if ranges[j].secondId >= ranges[i].secondId {
				ranges[i].secondId = ranges[j].secondId
			}
		}
		mergedRanges = append(mergedRanges, ranges[i])
		i = j
	}

	for _, p := range mergedRanges {
		result.fresh_ingredients += p.secondId - p.firstId + 1
	}
	
	return
}