package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var inputPath = "./day7/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type resultEntry struct {
	Index int
	Count int
}

func Star1() string {
	file, error := os.Open(inputPath)
	check(error)

	inputs := parseInput(file)
	result := make([]resultEntry, 0)
	for i := 0; i < len(inputs); i++ {
		fuelAmountTotal := 0
		for _, input := range inputs {
			fuelAmount := int(math.Abs(float64(input - i)))
			fuelAmountTotal += fuelAmount
		}

		result = append(result, resultEntry{Index: i, Count: fuelAmountTotal})
	}

	leastResult := findSmallestResult(result)

	fmt.Println("LEAST FUEL")
	fmt.Println("POSITION: ", leastResult.Index)
	fmt.Println("AMOUNT: ", leastResult.Count)

	return ""
}
func Star2() string {
	file, error := os.Open(inputPath)
	check(error)

	inputs := parseInput(file)
	result := make([]resultEntry, 0)

	for i := 0; i < len(inputs); i++ {
		fuelAmountTotal := 0
		for _, input := range inputs {
			c := int(math.Abs(float64(input - i)))
			fuelAmount := 0

			for x := 0; x < c; x++ {
				fuelAmount += x + 1
			}

			fuelAmountTotal += fuelAmount
		}

		result = append(result, resultEntry{Index: i, Count: fuelAmountTotal})
	}

	leastResult := findSmallestResult(result)

	fmt.Println("LEAST FUEL")
	fmt.Println("POSITION: ", leastResult.Index)
	fmt.Println("AMOUNT: ", leastResult.Count)

	return ""
}

func findSmallestResult(result []resultEntry) resultEntry {
	smallest := result[0]

	for _, res := range result {
		if res.Count < smallest.Count {
			smallest = res
		}
	}

	return smallest

}

// func Star2() string {
// 	file, error := os.Open(inputPath)
// 	check(error)

// 	input := parseInput(file)

// 	return ""
// }

func parseInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	input := make([]int, 0)
	for scanner.Scan() {
		inputElement := scanner.Text()
		inputText := strings.Split(inputElement, ",")
		for _, t := range inputText {
			n, err := strconv.Atoi(t)
			check(err)
			input = append(input, n)
		}
	}
	return input
}
