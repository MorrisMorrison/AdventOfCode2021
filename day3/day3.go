package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inputPath = "./day3/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Star1() string {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)

	gammaRate := make([]int, 0)
	epsilonRate := make([]int, 0)

	var row int
	var col int
	for col = 0; col < len(input[0]); col++ {
		var result [2]int

		for row = 0; row < len(input); row++ {
			if input[row][col] == 0 {
				result[0] += 1
			}
			if input[row][col] == 1 {
				result[1] += 1
			}
		}

		if result[0] > result[1] {
			gammaRate = append(gammaRate, 0)
			epsilonRate = append(epsilonRate, 1)
		} else {
			gammaRate = append(gammaRate, 1)
			epsilonRate = append(epsilonRate, 0)
		}
	}

	// fmt.Println(gammaRate)
	// fmt.Println(epsilonRate)

	return ""
}

func Star2() string {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)

	// oxygenGeneratorRating := make([]int, 0)
	// scrubberRating := make([]int, 0)

	// find oxygenGeneratorRating
	findOxygenGeneratorRating(input)
	scrubberRatingRating(input)
	return ""
}

func scrubberRatingRating(input [][]int) []int {
	result := make([]int, 0)
	possibleResults := make([][]int, 0)

	var col int

	for col = 0; col < len(input[0]); col++ {
		if len(possibleResults) == 1 {
			break
		}
		var leastCommonBit int
		var rowsWithMostCommonBit [][]int
		if col == 0 {
			leastCommonBit = findLeastCommonBit(getColumn(input, col))
			rowsWithMostCommonBit = findRowsWithMostCommonBit(input, leastCommonBit, col)
			possibleResults = rowsWithMostCommonBit

		} else {
			leastCommonBit = findLeastCommonBit(getColumn(possibleResults, col))
			rowsWithMostCommonBit = findRowsWithMostCommonBit(possibleResults, leastCommonBit, col)
			possibleResults = rowsWithMostCommonBit

		}
		for i := 0; i < len(possibleResults); i++ {
			if len(possibleResults) == 1 {
				break
			}
			var found bool

			for _, row := range rowsWithMostCommonBit {
				if compareSlices(possibleResults[i], row) {
					found = true
				}
			}

			if !found {
				possibleResults = removeRow(possibleResults, i)
			}
		}
	}

	fmt.Println(possibleResults)

	return result
}

func findLeastCommonBit(input []int) int {
	var result [2]int
	for _, value := range input {
		if value == 0 {
			result[0] += 1
		}
		if value == 1 {
			result[1] += 1
		}
	}

	if result[0] <= result[1] {
		return 0
	} else {
		return 1
	}
}

func findOxygenGeneratorRating(input [][]int) []int {
	result := make([]int, 0)
	possibleResults := make([][]int, 0)

	var col int

	for col = 0; col < len(input[0]); col++ {
		if len(possibleResults) == 1 {
			break
		}
		var mostCommonBit int
		var rowsWithMostCommonBit [][]int
		if col == 0 {
			mostCommonBit = findMostCommonBit(getColumn(input, col))
			rowsWithMostCommonBit = findRowsWithMostCommonBit(input, mostCommonBit, col)
			possibleResults = rowsWithMostCommonBit

		} else {
			mostCommonBit = findMostCommonBit(getColumn(possibleResults, col))
			rowsWithMostCommonBit = findRowsWithMostCommonBit(possibleResults, mostCommonBit, col)
			possibleResults = rowsWithMostCommonBit

		}
		for i := 0; i < len(possibleResults); i++ {
			if len(possibleResults) == 1 {
				break
			}
			var found bool

			for _, row := range rowsWithMostCommonBit {
				if compareSlices(possibleResults[i], row) {
					found = true
				}
			}

			if !found {
				possibleResults = removeRow(possibleResults, i)
			}
		}
	}

	fmt.Println(possibleResults)

	return result
}

func compareSlices(slice1 []int, slice2 []int) bool {
	for index, item := range slice1 {
		if item != slice2[index] {
			return false
		}
	}
	return true
}

func removeRow(slice [][]int, s int) [][]int {
	return append([][]int{}, append(slice[:s], slice[s+1:]...)...)
}

func findRowsWithMostCommonBit(input [][]int, mostCommonBit, col int) [][]int {
	result := make([][]int, 0)
	for row := 0; row < len(input); row++ {
		if input[row][col] == mostCommonBit {
			result = append(result, input[row])
		}
	}

	return result
}

func findMostCommonBit(input []int) int {
	var result [2]int
	for _, value := range input {
		if value == 0 {
			result[0] += 1
		}
		if value == 1 {
			result[1] += 1
		}
	}

	if result[1] >= result[0] {
		return 1
	} else {
		return 0
	}
}

func getColumn(input [][]int, colNum int) []int {
	result := make([]int, 0)
	for i := 0; i < len(input); i++ {
		result = append(result, input[i][colNum])
	}

	return result
}

func getRow(input [][]int, rowNum int) []int {
	result := make([]int, 0)
	for i := 0; i < len(input[0]); i++ {
		result = append(result, input[rowNum][i])
	}

	return result
}

func parseInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	input := make([][]int, 1000)
	var row int
	for scanner.Scan() {
		inputElement := scanner.Text()
		input[row] = make([]int, 12)

		for col, char := range inputElement {
			digit, err := strconv.Atoi(string(char))
			check(err)
			input[row][col] = digit
		}

		row++
	}
	return input
}
