package main

import (
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2024/day/9

const fileName = "input.txt"


func main() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	part1([]rune(string(data)))
	fmt.Println()
	part2([]rune(string(data)))
	fmt.Println()
	part2v2([]rune(string(data)))
}



func part1(data []rune) {
	result := 0
	counter := 0
	for start, end := 0, len(data) - 1; start <= end; {
		cs := 0
		if start & 1 == 0 {
			//start is a file index
			cs, counter, _ = getCheckSumForFile(data, start, counter, start)
			start += 1
		}else if end & 1 == 0 {
			//start is space index and end is file index
			hasSpaceForFileId := true
			cs, counter, hasSpaceForFileId = getCheckSumForFile(data, end, counter, start)
			if(hasSpaceForFileId) {
				end -= 2
			} else {
				start += 1
			}
		} else {
			//both indexes are pointing at spaces, move endIdx
			end -= 1
		}
		result += cs
	}
	fmt.Printf("Part1: %d", result)
}


func getCheckSumForFile(data []rune, fileIdx, counter, spaceIdx int) (int, int, bool) {
	fileId := fileIdx / 2
	result := 0
	nSpaces := toInt(data[spaceIdx])
	availableSpace := nSpaces - toInt(data[fileIdx])
	for ; nSpaces > 0 && nSpaces > availableSpace; nSpaces -= 1 {
		result += fileId * counter
		counter += 1
	}
	if availableSpace < 0 {
		data[fileIdx] = '0' + rune(-availableSpace)
	} else {
		data[fileIdx] = '0'
		data[spaceIdx] = '0' + rune(availableSpace)
	}
	return result, counter, availableSpace > 0
}

func toInt(r rune) int {
	return int(r - '0')
}

func part2(data []rune) {
	sys := transformData(data)
	sorted := sort(sys)
	checkSum := checkSum(sorted)

	fmt.Printf("Part2: %d", checkSum)
}

func printFileSpace(sys []FileSpace) {
	ids := []int{}
	counter := 0
	for _, fs := range sys {
		if fs.indexType == SPACE && len(fs.childs) != 0 {
			for _, c := range fs.childs {
				for i := 0; i < c.len; i +=1 {
					ids = append(ids, c.id)
					counter += 1
				}
			}
		} 
		for i := 0; i < fs.len; i +=1 {
			ids = append(ids, fs.id)
			counter += 1
		}
	}
	fmt.Printf("%v\n", ids)
}

func transformData(data []rune) []FileSpace {
	sys := []FileSpace{}
	// fileIds := 0
	for i, r := range data {
		fs := FileSpace {}
		if i & 1 == 0 {
			fs.indexType = FILE
			fs.id = i / 2
			fs.len = toInt(r)
		} else {
			fs.indexType = SPACE
			fs.id = 0
			fs.len = toInt(r)
		}
		fs.len = toInt(r)
		sys = append(sys, fs)
	}
	return sys
}

func sort(sys []FileSpace) []FileSpace {
	for s, e := 1, len(sys) - 1; e > 0; {
		if s > e {
			s = 1
			e -= 2
		}
		fs := sys[e]
		if fs.len == 0 {
			e -= 2
		}
		space := sys[s]
		if fs.len <= space.len {
			space.childs = append(space.childs, FileSpaceChild{id: fs.id, len: fs.len})
			space.len -= fs.len
			sys[e] = FileSpace{indexType: SPACE, id: 0, len: fs.len}
			sys[s] = space
			e -= 2
			s = 1
		} else {
			s += 2
		}
	}
	return sys
}

func checkSum(sys []FileSpace) int {
	checkSum := 0
	counter := 0
	for _, fs := range sys {
		if fs.indexType == SPACE && len(fs.childs) != 0 {
			for _, child := range fs.childs {
				checkSum += calcCheckSum(child.id, counter, child.len)
				counter += child.len
			}
		} 
		checkSum += calcCheckSum(fs.id, counter, fs.len)
		counter += fs.len
	}
	return checkSum
}

const (
	FILE = "F"
	SPACE = "S"
)

type FileSpace struct {
	indexType string
	len, id int
	childs []FileSpaceChild
}

type FileSpaceChild struct {
	id, len int
}

func part2v2(data []rune) {
	checkSum := 0

	virginData := make([]rune, len(data))
	copy(virginData, data)
	
	for e := len(data) - 1; e >= 0; e -=2 {
		counter := 0
		fileSize := toInt(data[e])
		fid := e / 2
		swapped := false
		for s:= 0; s < e; s += 1 {
			size := toInt(data[s])
			virginSize := toInt(virginData[s]) 
			if isSpace(s) && size >= fileSize {
				if size != virginSize {
					counter += 	virginSize - size
				}
				checkSum += calcCheckSum(fid, counter, fileSize)
				data[s] = '0' + rune(size - fileSize)
				swapped = true
				break
			}
			counter += virginSize
		}
		if !swapped {
			checkSum += calcCheckSum(fid, counter, fileSize)
		}
	}
	fmt.Printf("Part2v2: %d", checkSum)
}

func calcCheckSum(fid, counter, size int) int {
	return fid * (2*counter + size - 1) * size / 2
}

func isSpace(i int) bool {
	return i & 1 == 1
}