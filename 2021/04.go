package main

import (
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

func day4P1(game Game) int {
	return 1
}

//func day4P2(diagnostics []string) int {
//}

func day4() {
	printDay(4)
	input := readFile("./04_input.txt")
	game := parseInput(input)
	fmt.Println(game)
	fmt.Println(day4P1(game))
	//fmt.Println(day4P2(diagnostics))
}
