package day1

import (
	"bufio"
	"os"
	"strconv"
)

var inputPath = "./day1/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Star1() int {
	file, error := os.Open(inputPath)
	check(error)
	scanner := bufio.NewScanner(file)

	var count int
	var predecessor int

	for scanner.Scan() {
		i, error := strconv.Atoi(scanner.Text())
		check(error)

		if i > predecessor && predecessor != 0 {
			count++
		}

		predecessor = i
	}

	return count
}

func Star2() int {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)

	var count int
	for i := 0; i < len(input)-3; i++ {
		sum1 := input[i] + input[i+1] + input[i+2]
		sum2 := input[i+1] + input[i+2] + input[i+3]
		if sum2 > sum1 {
			count++
		}
	}

	return count
}

func parseInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	input := make([]int, 0)
	for scanner.Scan() {
		inputElement, error := strconv.Atoi(scanner.Text())
		input = append(input, inputElement)
		check(error)
	}
	return input
}
