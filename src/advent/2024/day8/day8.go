package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2024/day/8

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
	fMap := FMap{
		freqs: make(map[rune][]Position),
		lineLen: len(data),
		colLen: len(data[0]),
	}

	fMap.prepareData(data)

	result := fMap.calcAntiFreq()

	fmt.Printf("Part1: %d", result)
}

type FMap struct {
	freqs map[rune][]Position
	lineLen, colLen int
}

type Position struct {
	x,y int
} 

func (fmap *FMap) prepareData(data []string) {
	for l, line := range data {
		for c, char := range line {
			if char == '.' {
				continue
			}
			fmap.freqs[char] = append(fmap.freqs[char], Position{x: l, y: c})
		}
	}
}

func (fmap *FMap) calcAntiFreq() int {
	antiFreqs := make(map[Position]bool)
	for _, frequencies := range fmap.freqs {
		if len(frequencies) < 2 {
			continue
		}
		maxLen := len(frequencies)/2 + 1
		for i, node := range frequencies[:maxLen] {
			for _, nextNode := range frequencies[i+1:] {
				vector := findVector(node, nextNode)
				
				if antiNode1, ok1 := fmap.tryPlaceNewNode(node, vector, -1); ok1 {
					antiFreqs[antiNode1] = true
				} 
				if antiNode2, ok2 := fmap.tryPlaceNewNode(nextNode, vector, 1); ok2 {
					antiFreqs[antiNode2] = true
				}
			}
		}
	}
	return len(antiFreqs)
}


func findVector(a, b Position) Position {
	return Position{
		x: b.x - a.x,
		y: b.y - a.y,
	}
}

func (m *FMap) tryPlaceNewNode(a, v Position, d int) (Position, bool) {
	pos := Position{
		x: a.x + v.x * d,
		y: a.y + v.y * d,
	}

	return pos, !m.isOutOfBounds(pos)
}

func (m *FMap) isOutOfBounds(p Position) bool {
	return p.x < 0 || p.y < 0 || p.x >= m.lineLen ||  p.y >= m.colLen
}


func part2(data []string) {
	fMap := FMap{
		freqs: make(map[rune][]Position),
		lineLen: len(data),
		colLen: len(data[0]),
	}

	fMap.prepareData(data)

	result := fMap.calcAntiFreqLooped()

	fmt.Printf("Part2: %d", result)
}

func (fmap *FMap) calcAntiFreqLooped() int {
	antiFreqs := make(map[Position]bool)
	for _, frequencies := range fmap.freqs {
		if len(frequencies) < 2 {
			continue
		}
		maxLen := len(frequencies)/2 + 1
		for i, node := range frequencies[:maxLen] {
			for _, nextNode := range frequencies[i+1:] {
				vector := findVector(node, nextNode)
				antiNode, ok := fmap.tryPlaceNewNode(node, vector, 1)
				for ok {
					antiFreqs[antiNode] = true
					antiNode, ok = fmap.tryPlaceNewNode(antiNode, vector, 1)
				}
				antiNode, ok = fmap.tryPlaceNewNode(nextNode, vector, -1)
				for ok {
					antiFreqs[antiNode] = true
					antiNode, ok = fmap.tryPlaceNewNode(antiNode, vector, -1)
				}
			}
		}
	}
	return len(antiFreqs)
}
