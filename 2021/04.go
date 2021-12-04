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
	for _, num := range strings.Split(input, sep) {
		intNum, _ := strconv.Atoi(num)
		numbers = append(numbers, intNum)
	}
	return numbers
}

func parseInput(input []string) Game {
	numbersRow := input[0]
	numbers := parseToInts(numbersRow, ",")

	var gameRows []string = input[1:]
	for _, row := gameRows {
		var boardVals [5][5]int
		if strings.Trim(row) != "" {
		}
	}

	return Game{numbers: numbers, boards: []Board{}}
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
