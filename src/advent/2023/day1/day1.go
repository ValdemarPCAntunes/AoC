package main

// --- Day 1: Trebuchet?! ---
// Something is wrong with global snow production, and you've been selected to take a look.
// The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.
// You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.
// Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar;
// the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!
// You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky")
// and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from")
// when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").
// As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf
// who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

// The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover.
// On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

// For example:

// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

// Consider your entire calibration document. What is the sum of all of the calibration values?

// --- Part Two ---
// Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

// Equipped with this new information, you now need to find the real first and last digit on each line. For example:

// two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen
// In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

// What is the sum of all of the calibration values?

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const fileName = "input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	readFile, err := os.Open(fileName)

	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalSum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var num1 int
		var num2 int
		length := len(line)
		for i := 0; i < length; i++ {
			c := line[i] - '0'
			if c <= 9 {
				num1 = int(c)
				break
			}
		}
		for i := length - 1; i >= 0; i-- {
			c := line[i] - '0'
			if c <= 9 {
				num2 = int(c)
				break
			}
		}
		// fmt.Printf("%#v : scanned %d%d\n", line, num1, num2)
		totalSum += num1*10 + num2
	}

	readFile.Close()
	fmt.Printf("\nPart1 Total sum = %d", totalSum)
}

func part2() {
	type pair struct {
		key   string
		value int
	}

	valuesMap := make(map[string]pair)

	valuesMap["one"] = pair{key: "one", value: 1}
	valuesMap["two"] = pair{key: "two", value: 2}
	valuesMap["thr"] = pair{key: "three", value: 3}
	valuesMap["ree"] = valuesMap["thr"]
	valuesMap["fou"] = pair{key: "four", value: 4}
	valuesMap["our"] = valuesMap["fou"]
	valuesMap["fiv"] = pair{key: "five", value: 5}
	valuesMap["ive"] = valuesMap["fiv"]
	valuesMap["six"] = pair{key: "six", value: 6}
	valuesMap["sev"] = pair{key: "seven", value: 7}
	valuesMap["ven"] = valuesMap["sev"]
	valuesMap["eig"] = pair{key: "eight", value: 8}
	valuesMap["ght"] = valuesMap["eig"]
	valuesMap["nin"] = pair{key: "nine", value: 9}
	valuesMap["ine"] = valuesMap["nin"]

	readFile, err := os.Open(fileName)

	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalSum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var num1 int
		var num2 int
		length := len(line)
		for i := 0; i < length; i++ {
			c := line[i] - '0'
			if c <= 9 {
				num1 = int(c)
				break
			}

			if i >= length-3 {
				continue
			}

			val, ok := valuesMap[line[i:i+3]] 

			if ok && val.key == line[i:i+len(val.key)] {
				
				num1 = val.value
				break
			}

		}
		for i := length - 1; i >= 0; i-- {
			c := line[i] - '0'
			if c <= 9 {
				num2 = int(c)
				break
			}
			if i <= 3 {
				continue
			}
			val, ok := valuesMap[line[i-2:i+1]] 

			if ok && val.key == line[i - len(val.key) + 1:i+1]{
				num2 = val.value
				break
			}
		}
		//fmt.Printf("%#v : scanned %d%d\n", line, num1, num2)
		totalSum += num1*10 + num2
	}

	readFile.Close()
	fmt.Printf("\nPart2 Total sum = %d", totalSum)
}

// TRIED WITH PASSING FUNCTIONS BUT ENDED UP WORSE IN VERBOSE TERMS WELP

// fgetValueAsc := func(line string) func(int) pairSuccess {
// 	return func(i int) pairSuccess {
// 		if i+3 == len(line)+1 {
// 			fmt.Println("here")
// 		}
// 		val, ok := valuesMap[line[i:i+3]]
// 		return pairSuccess{
// 			success: ok && (len(line)-i) >= len(val.key) && val.key == line[i:i+len(val.key)],
// 			value:   val.value,
// 		}
// 	}
// }

// fgetValueDesc := func(line string) func(int) pairSuccess {
// 	return func(i int) pairSuccess {
// 		val, ok := valuesMap[line[i-2:i+1]]
// 		return pairSuccess{
// 			success: ok && i >= len(val.key) && val.key == line[i-len(val.key)+1:i+1],
// 			value:   val.value,
// 		}
// 	}
// }

// for fileScanner.Scan() {
// 	line := fileScanner.Text()
// 	var num1 int
// 	var num2 int
// 	length := len(line)
// 	for i, j := 0, length-1; (num1 * num2) == 0; {
// 		if(num1 == 0) {
// 			result1 := getNumberFromString(i, line[i]-'0', fgetValueAsc(line), func(k int) bool { return k <= length-3 })
// 			if result1.success {
// 				num1 = result1.value
// 			} else {
// 				i = i + 1
// 			}
// 		}
// 		if(num2 == 0) {
// 			result2 := getNumberFromString(j, line[j]-'0', fgetValueDesc(line), func(k int) bool { return k >= 3 })
// 			if result2.success {
// 				num2 = result2.value
// 			} else {
// 				j = j - 1
// 			}
// 		}
// 	}

// type pairSuccess struct {
// 	value   int
// 	success bool
// }

// func getNumberFromString(
// 	i int,
// 	firstChar byte,
// 	fgetValue func(idx int) pairSuccess,
// 	fcanGetValue func(idx int) bool,
// ) pairSuccess {
// 	if firstChar <= 9 {
// 		return pairSuccess{value: int(firstChar), success: true}
// 	}

// 	if !fcanGetValue(i) {
// 		return pairSuccess{}
// 	}
// 	result := fgetValue(i)
// 	return pairSuccess{success: result.success, value: result.value}
// }
