package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inputPath = "./day9/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Star1() string {
	file, error := os.Open(inputPath)
	check(error)

	result := make([]int, 0)
	input := parseInput(file)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			toCheck := input[i][j]
			isLowest := true
			for k := j - 1; k <= j+1; k++ {
				if k < 0 || k > len(input[i])-1 || k == j {
					continue
				}

				if input[i][k] <= toCheck {
					isLowest = false
				}

				if !isLowest {
					break
				}
			}
			for l := i - 1; l <= i+1; l++ {
				if l < 0 || l > len(input)-1 || l == i {
					continue
				}

				if input[l][j] <= toCheck {
					isLowest = false
				}

				if !isLowest {
					break
				}
			}
			if isLowest {
				result = append(result, toCheck)
			}
		}

	}
	fmt.Println(result)
	sum := 0
	for _, value := range result {
		sum += value + 1
	}

	fmt.Println("SUM", sum)

	return ""
}

func Star2() string {
	file, error := os.Open(inputPath)
	check(error)

	parseInput(file)

	return ""
}

func parseInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	input := make([][]int, 100)
	var row int
	for scanner.Scan() {
		inputElement := scanner.Text()
		input[row] = make([]int, len(inputElement))

		for col, char := range inputElement {
			digit, err := strconv.Atoi(string(char))
			check(err)
			input[row][col] = digit
		}
		row++
	}

	return input
}
