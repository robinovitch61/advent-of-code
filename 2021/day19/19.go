package day19

import (
	"aoc/common"
	"fmt"
)

type PuzzleInput struct{}

func parseInput(input []string) PuzzleInput {
	return PuzzleInput{}
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return -1
}

func p2(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(19)
	input := common.ReadFile("19")
	puzzleInput := parseInput(input)
	fmt.Println(p1(puzzleInput))
	fmt.Println(p2(puzzleInput))
}
