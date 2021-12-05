package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPath = "./day4/input.txt"
var _draws []int

type bingoBoard [][]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Star1() string {
	file, error := os.Open(inputPath)
	check(error)
	isBingo := false
	input := parseInput(file)
	var finalDraw int
	var sum int

	for _, draw := range _draws {

		for index, board := range input {
			isBingo = checkBingo(board)
			if isBingo {
				fmt.Println("BINGO !!!")
				fmt.Println(board)
				sum = calculateBoardSum(board)
				break
			}

			input[index] = markNumber(board, draw)
		}

		if isBingo {
			break
		}
		finalDraw = draw
	}
	fmt.Printf("FINAL DRAW %v", finalDraw)
	fmt.Println()
	fmt.Printf("SUM %v", sum)

	return ""
}

func calculateBoardSum(board bingoBoard) int {
	sum := 0

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[1]); col++ {
			if board[row][col] != -1 {
				sum += board[row][col]
			}
		}
	}

	return sum
}

func markNumber(board bingoBoard, draw int) bingoBoard {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[1]); col++ {
			if board[row][col] == draw {
				board[row][col] = -1
			}
		}
	}

	return board
}

func checkBingo(board [][]int) bool {
	for i := 0; i < 5; i++ {
		row := getRow(board, i)
		col := getColumn(board, i)

		if everyElementMatches(row, -1) {
			return true
		}

		if everyElementMatches(col, -1) {
			return true
		}
	}

	return false
}

func everyElementMatches(input []int, matchingElement int) bool {
	for _, element := range input {
		if element != matchingElement {
			return false
		}
	}

	return true
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
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func Star2() string {
	file, error := os.Open(inputPath)
	check(error)
	isBingo := false
	input := parseInput(file)
	var finalDraw int
	var sum int
	totalBoardCount := len(input)
	finishedBoardIndizes := make([]int, 0)
	finishedBoardCount := 0

	for _, draw := range _draws {

		for index, board := range input {
			input[index] = markNumber(board, draw)
			isBingo = checkBingo(board)
			if isBingo {
				if !contains(finishedBoardIndizes, index) {
					finishedBoardIndizes = append(finishedBoardIndizes, index)
					finishedBoardCount++
				}
			}

			if finishedBoardCount == totalBoardCount {
				fmt.Println(board)
				sum = calculateBoardSum(board)
				break
			}

		}
		finalDraw = draw

		if finishedBoardCount == totalBoardCount {
			break
		}
	}
	fmt.Printf("FINAL DRAW %v", finalDraw)
	fmt.Println()
	fmt.Printf("SUM %v", sum)

	return ""
}

func parseInput(file *os.File) []bingoBoard {
	scanner := bufio.NewScanner(file)

	boards := make([]bingoBoard, 0)
	board := make(bingoBoard, 5)
	for i := 0; i < 5; i++ {
		board[i] = make([]int, 5)
	}
	var rowNum int
	var boardRowNum int
	var draws []int
	for scanner.Scan() {
		row := scanner.Text()
		if rowNum == 0 {
			draws = make([]int, 0)
			drawsString := strings.Split(row, ",")
			for _, s := range drawsString {
				num, err := strconv.Atoi(s)
				check(err)
				draws = append(draws, num)
			}
			rowNum++
		} else {
			if row != "" {
				fields := strings.Fields(row)
				for index, field := range fields {
					num, err := strconv.Atoi(field)
					check(err)
					board[boardRowNum][index] = num
				}

				if (boardRowNum+1)%5 == 0 {
					boards = append(boards, board)
					board = make(bingoBoard, 5)
					for i := 0; i < 5; i++ {
						board[i] = make([]int, 5)
					}
					boardRowNum = 0
				} else {
					boardRowNum++
				}

				rowNum++
			}
		}
	}

	_draws = draws
	fmt.Println("DRAWS")
	fmt.Println(_draws)
	return boards
}
