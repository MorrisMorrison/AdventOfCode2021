package day6

import (
	"bufio"
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
		for j, input := range inputs {
			if input == 0 {
				inputs[j] = 6
				inputs = append(inputs, 8)
			} else {
				inputs[j] -= 1

			}

		}
	}

	return strconv.Itoa(len(inputs))
}

func Star2() string {
	file, error := os.Open(inputPath)
	check(error)

	inputs := parseInput(file)
	days := 256
	result := make(map[int]int)
	for _, input := range inputs {
		result[input] += 1
	}
	for i := 0; i < days; i++ {
		current := result[0]
		for j := 1; j < 10; j++ {
			result[j-1] = result[j]
		}
		result[6] += current
		result[8] = current
	}

	sum := 0
	for _, v := range result {
		sum += v
	}

	return strconv.Itoa(sum)
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
