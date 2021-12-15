package day15

import (
	"aoc/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point [2]int

type PuzzleInput map[Point]int

func parseInput(input []string) PuzzleInput {
	posToRisk := make(map[Point]int)
	for i, row := range input {
		for j, riskStr := range strings.Split(row, "") {
			risk, _ := strconv.Atoi(riskStr)
			posToRisk[Point{i, j}] = risk
		}
	}
	return posToRisk
}

func getUnvisitedNeighbors(current Point, visited map[Point]bool, sideLength int) []Point {
	var unvisited []Point
	up := Point{current[0] - 1, current[1]}
	if _, v := visited[up]; !v && up[0] >= 0 {
		unvisited = append(unvisited, up)
	}
	right := Point{current[0], current[1] + 1}
	if _, v := visited[right]; !v && right[1] < sideLength {
		unvisited = append(unvisited, right)
	}
	down := Point{current[0] + 1, current[1]}
	if _, v := visited[down]; !v && down[0] < sideLength {
		unvisited = append(unvisited, down)
	}
	left := Point{current[0], current[1] - 1}
	if _, v := visited[left]; !v && left[1] >= 0 {
		unvisited = append(unvisited, left)
	}
	return unvisited
}

func smallestDistanceUnvisitedPoint(distances map[Point]int, visited map[Point]bool) Point {
	minDist := math.MaxInt64
	var minPoint Point
	for point, dist := range distances {
		if _, v := visited[point]; dist < minDist && !v {
			minDist = dist
			minPoint = point
		}
	}
	return minPoint
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()
	sideLength := int(math.Sqrt(float64(len(puzzleInput))))
	endPos := Point{sideLength - 1, sideLength - 1}
	visited := make(map[Point]bool)
	distances := make(map[Point]int)
	for pos, _ := range puzzleInput {
		distances[pos] = math.MaxInt64
	}
	current := Point{0, 0}
	distances[current] = 0
	for {
		for _, neighbor := range getUnvisitedNeighbors(current, visited, sideLength) {
			distance := distances[current] + puzzleInput[neighbor]
			prevDistance := distances[neighbor]
			if distance < prevDistance {
				distances[neighbor] = distance
			}
		}
		visited[current] = true
		if current == endPos {
			break
		} else {
			current = smallestDistanceUnvisitedPoint(distances, visited)
		}
	}
	return distances[endPos]
}

func p2(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(15)
	input := common.ReadFile("15")
	puzzleInput := parseInput(input)
	fmt.Println(p1(puzzleInput))
	fmt.Println(p2(puzzleInput))
}
