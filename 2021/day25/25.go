package day25

import (
	"aoc/common"
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

type Locations map[Point]bool

type Board struct {
	easts, souths, empties Locations
	width, height          int
}

func (b Board) Print() {
	for x := 0; x < b.height; x++ {
		row := ""
		for y := 0; y < b.width; y++ {
			p := Point{x, y}
			if b.easts[p] {
				row += ">"
			} else if b.souths[p] {
				row += "v"
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
	fmt.Println("")
}

func parseInput(input []string) Board {
	easts, souths, empties := make(Locations), make(Locations), make(Locations)
	width := len(input[0])
	height := len(input)
	for i, row := range input {
		for j, char := range strings.Split(row, "") {
			switch char {
			case ".":
				empties[Point{i, j}] = true
			case ">":
				easts[Point{i, j}] = true
			case "v":
				souths[Point{i, j}] = true
			}
		}
	}
	return Board{easts, souths, empties, width, height}
}

func southPoint(point Point, height int) Point {
	if point.x+1 >= height {
		return Point{0, point.y}
	}
	return Point{point.x + 1, point.y}
}

func eastPoint(point Point, width int) Point {
	if point.y+1 >= width {
		return Point{point.x, 0}
	}
	return Point{point.x, point.y + 1}
}

func doStep(board Board) (Board, bool) {
	moveEast := make(Locations)
	moveSouth := make(Locations)
	for p, v := range board.easts {
		if v && board.empties[eastPoint(p, board.width)] {
			moveEast[p] = true
		}
	}
	for p := range moveEast {
		board.easts[p] = false
		board.empties[p] = true
		board.easts[eastPoint(p, board.width)] = true
		board.empties[eastPoint(p, board.width)] = false
	}

	for p, v := range board.souths {
		if v && board.empties[southPoint(p, board.height)] {
			moveSouth[p] = true
		}
	}
	for p := range moveSouth {
		board.souths[p] = false
		board.empties[p] = true
		board.souths[southPoint(p, board.height)] = true
		board.empties[southPoint(p, board.height)] = false
	}
	return board, len(moveEast) == 0 && len(moveSouth) == 0
}

func p1(puzzleInput Board) int {
	defer common.Time()()
	noneMoved := false
	board := puzzleInput
	for step := 0; ; step++ {
		if noneMoved {
			return step
		}
		board, noneMoved = doStep(board)
	}
}

func p2(puzzleInput Board) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(25)
	input := common.ReadFile("25")
	puzzleInput := parseInput(input)
	fmt.Println(p1(puzzleInput))
	//fmt.Println(p2(puzzleInput))
}
