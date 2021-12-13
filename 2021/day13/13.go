package day13

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type Dot struct {
	right int
	down  int
}

type FoldInstruction struct {
	location     int
	isHorizontal bool // true when y
}

type PuzzleInput struct {
	dots             []Dot
	foldInstructions []FoldInstruction
}

func parseInput(input []string) PuzzleInput {
	var start []string
	var end []string
	doEnd := false
	for _, row := range input {
		if strings.TrimSpace(row) == "" {
			doEnd = true
		} else {
			if !doEnd {
				start = append(start, row)
			} else {
				end = append(end, row)
			}
		}
	}

	var points []Dot
	for _, point := range start {
		split := strings.Split(point, ",")
		right, _ := strconv.Atoi(split[0])
		down, _ := strconv.Atoi(split[1])
		points = append(points, Dot{right, down})
	}

	var folds []FoldInstruction
	for _, fold := range end {
		split := strings.Split(fold, "fold along ")
		split2 := strings.Split(split[1], "=")
		isHorizontal := true
		if split2[0] == "x" {
			isHorizontal = false
		}
		location, _ := strconv.Atoi(split2[1])
		folds = append(folds, FoldInstruction{location, isHorizontal})
	}
	return PuzzleInput{points, folds}
}

func foldNum(num int, loc int) int {
	dist := num - loc
	return num - 2*dist
}

func foldPaper(fold FoldInstruction, dots []Dot) {
	if fold.isHorizontal {
		// only affects dots with y-vals > fold.location
		for idx, dot := range dots {
			if dot.down > fold.location {
				dots[idx] = Dot{dot.right, foldNum(dot.down, fold.location)}
			}
		}
	} else {
		// only affects dots with y-vals > fold.location
		for idx, dot := range dots {
			if dot.right > fold.location {
				dots[idx] = Dot{foldNum(dot.right, fold.location), dot.down}
			}
		}
	}
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()
	fold := puzzleInput.foldInstructions[0]
	dots := puzzleInput.dots
	foldPaper(fold, dots)

	visibleDots := make(map[Dot]bool)
	countVisibleDots := 0
	for _, dot := range dots {
		if _, exists := visibleDots[dot]; !exists {
			visibleDots[dot] = true
			countVisibleDots++
		}
	}
	return countVisibleDots
}

func plotGrid(visibleDots map[Dot]bool) {
	const maxDown = 40
	const maxRight = 25
	var grid [maxDown][maxRight]string
	for down := 0; down < maxDown; down++ {
		for right := 0; right < maxRight; right++ {
			if visibleDots[Dot{down, right}] {
				grid[down][right] = "#"
			} else {
				grid[down][right] = "."
			}
		}
	}
	for _, row := range grid {
		fmt.Println(row)
	}
}

func p2(puzzleInput PuzzleInput) {
	defer common.Time()()
	dots := puzzleInput.dots
	for _, fold := range puzzleInput.foldInstructions {
		foldPaper(fold, dots)
	}

	visibleDots := make(map[Dot]bool)
	for _, dot := range dots {
		if _, exists := visibleDots[dot]; !exists {
			visibleDots[dot] = true
		}
	}
	plotGrid(visibleDots)
}

func Run() {
	common.PrintDay(13)
	input := common.ReadFile("13")
	puzzleInput := parseInput(input)
	fmt.Println(p1(puzzleInput))
	p2(puzzleInput)
}
