package day7

import (
	"bufio"
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
	result := 0
	for i := 0; i < len(inputs); i++ {
		fuelAmountTotal := 0
		for _, input := range inputs {
			fuelAmountTotal += int(math.Abs(float64(input - i)))
		}

		if fuelAmountTotal < result || i == 0 {
			result = fuelAmountTotal
		}

	}

	return strconv.Itoa(result)
}

func Star2() string {
	file, error := os.Open(inputPath)
	check(error)

	inputs := parseInput(file)
	result := 0

	for i := 0; i < len(inputs); i++ {
		fuelAmountTotal := 0
		for _, input := range inputs {
			distance := int(math.Abs(float64(input - i)))

			for x := 0; x < distance; x++ {
				fuelAmountTotal += x + 1
			}

		}

		if fuelAmountTotal < result || i == 0 {
			result = fuelAmountTotal
		}

	}

	return strconv.Itoa(result)
}

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
