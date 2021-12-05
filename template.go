package main

import (
	"bufio"
	"os"
	"strconv"
)

var inputPath = "./template/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Star1() string {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)

	return ""
}

func Star2() string {
	file, error := os.Open(inputPath)
	check(error)

	input := parseInput(file)

	return ""
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
