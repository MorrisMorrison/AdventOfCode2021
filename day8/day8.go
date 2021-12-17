package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var inputPath = "./day8/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 1 -> 2
// 7 -> 3
// 4 -> 4
// 8 -> 7

func Star1() string {
	valuesToCheck := []int{2, 3, 4, 7}

	file, error := os.Open(inputPath)
	check(error)

	result := make([]int, 9)
	inputs := parseInput(file)
	for _, input := range inputs {
		fields := strings.Fields(input)
		for _, field := range fields {
			result[utf8.RuneCountInString(field)-1] += 1
		}
	}

	sum := 0
	for _, valueToCheck := range valuesToCheck {
		sum += result[valueToCheck-1]
	}

	return strconv.Itoa(sum)
}

func Star2() string {

	patterns := map[int]string{
		0: "deagbc",
		1: "ab",
		2: "dafgc",
		3: "dafbc",
		4: "efab",
		5: "defbc",
		6: "defbcg",
		7: "dab",
		8: "dafegcb",
		9: "dafecb",
	}

	file, error := os.Open(inputPath)
	check(error)

	inputs := parseInput(file)
	results := make([]string, 0)

	for _, input := range inputs {
		result := ""

		fields := strings.Fields(input)
		for _, field := range fields {
			digit := decode(patterns, field)
			result += strconv.Itoa(digit)
		}

		results = append(results, result)
		fmt.Print(input, " ")
		fmt.Println(result)
	}

	sum := 0
	for _, res := range results {
		i, err := strconv.Atoi(res)
		check(err)
		sum += i
	}

	fmt.Println("Sum")
	fmt.Println(sum)

	return ""
}

func decode(patterns map[int]string, field string) int {
	var result int

	for k, v := range patterns {
		matchesPattern := true
		if len(v) == len(field) {
			for _, c := range v {
				if !contains(field, c) {
					matchesPattern = false
					break
				}
			}
		} else {
			matchesPattern = false
		}
		if matchesPattern {
			result = k
			break
		}
	}

	return result
}

func contains(s string, r rune) bool {
	for _, rs := range s {
		if r == rs {
			return true
		}
	}

	return false
}

func parseInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	var row int
	for scanner.Scan() {
		line := scanner.Text()
		lineElements := strings.Split(line, "|")
		input = append(input, lineElements[1])
		row++
	}
	return input
}
