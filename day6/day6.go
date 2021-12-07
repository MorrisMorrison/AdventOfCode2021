package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPath = "./day6/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Star1() string {
	file, error := os.Open(inputPath)
	check(error)

	inputs := parseInput(file)
	daysCount := 80

	for i := 0; i < daysCount; i++ {
		for index, input := range inputs {
			if input == 0 {
				inputs[index] = 6
				inputs = append(inputs, 8)
			} else {
				inputs[index] -= 1

			}

		}
	}

	count := len(inputs)
	fmt.Println("COUNT LANTERNFISH ", count)
	return ""
}

// Buffer
func Star2() string {
	file, error := os.Open(inputPath)
	check(error)

	inputs := parseInput(file)
	chunks := make([][]int, 1000000)
	for i := 0; i < len(chunks); i++ {
		chunks[i] = make([]int, 50000)
		if i == 0 {
			for index, input := range inputs {
				chunks[0][index] = input
			}
		}
	}
	daysCount := 256
	countElements := len(inputs)
	currentPosition := len(inputs)
	for i := 0; i < daysCount; i++ {
		for index, _ := range chunks {
			for j := 0; j < countElements; j++ {
				if chunks[index][j] == 0 {
					chunks[index][j] = 6
					chunks[index][currentPosition] = 8
					countElements++
					currentPosition++
				} else {
					chunks[index][j] -= 1
					currentPosition++
				}

				if currentPosition == 9999 || j == countElements-1 {
					j = 0
					currentPosition = 0
					index++
				}
			}
		}
	}

	// fmt.Println(inputs)
	count := len(inputs)
	fmt.Println("COUNT LANTERNFISH ", count)
	return ""
}

// chunks

// func Star2() string {
// 	file, error := os.Open(inputPath)
// 	check(error)

// 	inputs := parseInput(file)
// 	chunks := make([][]int, 1000000)
// 	for i := 0; i < len(chunks); i++ {
// 		chunks[i] = make([]int, 50000)
// 		if i == 0 {
// 			for index, input := range inputs {
// 				chunks[0][index] = input
// 			}
// 		}
// 	}
// 	daysCount := 256
// 	countElements := len(inputs)
// 	currentPosition := len(inputs)
// 	for i := 0; i < daysCount; i++ {
// 		for index, _ := range chunks {
// 			for j := 0; j < countElements; j++ {
// 				if chunks[index][j] == 0 {
// 					chunks[index][j] = 6
// 					chunks[index][currentPosition] = 8
// 					countElements++
// 					currentPosition++
// 				} else {
// 					chunks[index][j] -= 1
// 					currentPosition++
// 				}

// 				if currentPosition == 9999 || j == countElements-1 {
// 					j = 0
// 					currentPosition = 0
// 					index++
// 				}
// 			}
// 		}
// 	}

// 	// fmt.Println(inputs)
// 	count := len(inputs)
// 	fmt.Println("COUNT LANTERNFISH ", count)
// 	return ""
// }

func isChunkFinished(chunk []int, pos int) bool {
	part := chunk[pos:]
	for _, p := range part {
		if p != 0 {
			return false
		}
	}

	return true
}

func isChunkFull(chunk []int) bool {
	return chunk[len(chunk)-1] != 0
}

func parseInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	result := make([]int, 0)
	var row int
	for scanner.Scan() {
		inputElement := scanner.Text()
		inputs := strings.Split(inputElement, ",")
		for _, input := range inputs {
			d, err := strconv.Atoi(input)
			check(err)
			result = append(result, d)
		}
		row++
	}

	return result
}
