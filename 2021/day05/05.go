package day05

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func parsePoint(input string, sep string) Point {
	split := strings.Split(input, sep)
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return Point{x, y}
}

func parseInput(input []string) []Line {
	var lines []Line
	for _, row := range input {
		split := strings.Split(row, " -> ")
		firstPoint := parsePoint(split[0], ",")
		secondPoint := parsePoint(split[1], ",")
		lines = append(lines, Line{firstPoint, secondPoint})
	}
	return lines
}

func pointsOnLine(p1 Point, p2 Point) []Point {
	// bresenham's from wikipedia
	dx := p2.x - p1.x
	if dx < 0 {
		dx = -dx
	}
	dy := p2.y - p1.y
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if p1.x < p2.x {
		sx = 1
	} else {
		sx = -1
	}
	if p1.y < p2.y {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	var points []Point
	x0 := p1.x
	y0 := p1.y
	for {
		points = append(points, Point{x0, y0})
		if x0 == p2.x && y0 == p2.y {
			return points
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func solve(lines []Line, cond func(Line) bool) int {
	pointCounter := make(map[Point]int)
	for _, line := range lines {
		if cond(line) {
			points := pointsOnLine(line.start, line.end)
			for _, point := range points {
				if pointCounter[point] == 0 {
					pointCounter[point] = 1
				} else {
					pointCounter[point] += 1
				}
			}
		}
	}
	moreThanTwo := 0
	for _, count := range pointCounter {
		if count > 1 {
			moreThanTwo++
		}
	}
	return moreThanTwo
}

func p1(lines []Line) int {
	cond := func(line Line) bool {
		return (line.start.x == line.end.x) || (line.start.y == line.end.y)
	}
	return solve(lines, cond)
}

func p2(lines []Line) int {
	return solve(lines, func(line Line) bool { return true })
}

func Run() {
	common.PrintDay(5)
	input := common.ReadFile("05")
	lines := parseInput(input)
	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}
