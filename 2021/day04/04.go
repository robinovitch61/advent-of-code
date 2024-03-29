package day04

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	values [5][5]int
}

type Game struct {
	numbers []int
	boards  []Board
}

func parseToInts(input string, sep string) []int {
	var numbers []int
	var allStrings []string
	if sep == " " {
		allStrings = strings.Fields(input)
	} else {
		allStrings = strings.Split(input, sep)
	}
	for _, num := range allStrings {
		intNum, _ := strconv.Atoi(num)
		numbers = append(numbers, intNum)
	}
	return numbers
}

func parseInput(input []string) Game {
	numbersRow := input[0]
	numbers := parseToInts(numbersRow, ",")

	var gameRows = input[1:]
	var boardVals [5][5]int
	var boardRowVals [5]int
	var boardRow int
	var boards []Board
	for _, row := range gameRows {
		if strings.Trim(row, " ") != "" {
			copy(boardRowVals[:], parseToInts(row, " "))
			boardVals[boardRow] = boardRowVals
			boardRow++
			if boardRow == 5 {
				boards = append(boards, Board{boardVals})
				boardRow = 0
			}
		}
	}

	return Game{numbers, boards}
}

func isWinner(board Board, drawn map[int]bool) bool {
	winner := false
	for i := 0; i < 5; i++ {
		rowWinner := true
		colWinner := true
		for _, num := range board.values[i] {
			if !drawn[num] {
				rowWinner = false
			}
		}
		for _, row := range board.values {
			if !drawn[row[i]] {
				colWinner = false
			}
		}
		if rowWinner || colWinner {
			winner = true
		}
	}
	return winner
}

func calcScore(board Board, drawn map[int]bool, lastDrawn int) int {
	sum := 0
	for down := 0; down < 5; down++ {
		for across := 0; across < 5; across++ {
			if val := board.values[down][across]; !drawn[val] {
				sum += val
			}
		}
	}
	return sum * lastDrawn
}

func p1(game Game) int {
	defer common.Time()()
	drawn := make(map[int]bool)
	for _, num := range game.numbers {
		drawn[num] = true
		for _, board := range game.boards {
			if isWinner(board, drawn) {
				return calcScore(board, drawn, num)
			}
		}
	}
	return -1
}

func p2(game Game) int {
	defer common.Time()()
	wonBoards := make(map[Board]bool)
	drawn := make(map[int]bool)
	for _, num := range game.numbers {
		drawn[num] = true
		for _, board := range game.boards {
			if !wonBoards[board] {
				if isWinner(board, drawn) {
					wonBoards[board] = true
					if len(wonBoards) == len(game.boards) {
						return calcScore(board, drawn, num)
					}
				}
			}
		}
	}
	return -1
}

func Run() {
	common.PrintDay(4)
	input := common.ReadFile("04")
	game := parseInput(input)
	fmt.Println(p1(game))
	fmt.Println(p2(game))
}
