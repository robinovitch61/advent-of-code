package day15

import (
	"aoc/common"
	"aoc/priority_queue"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type PuzzleInput map[common.Point]int

func parseInput(input []string) PuzzleInput {
	posToRisk := make(map[common.Point]int)
	for i, row := range input {
		for j, riskStr := range strings.Split(row, "") {
			risk, _ := strconv.Atoi(riskStr)
			posToRisk[common.Point{i, j}] = risk
		}
	}
	return posToRisk
}

func getUnvisitedNeighbors(current common.Point, visited map[common.Point]bool, sideLength int) []common.Point {
	var unvisited []common.Point
	up := common.Point{current[0] - 1, current[1]}
	if _, v := visited[up]; !v && up[0] >= 0 {
		unvisited = append(unvisited, up)
	}
	right := common.Point{current[0], current[1] + 1}
	if _, v := visited[right]; !v && right[1] < sideLength {
		unvisited = append(unvisited, right)
	}
	down := common.Point{current[0] + 1, current[1]}
	if _, v := visited[down]; !v && down[0] < sideLength {
		unvisited = append(unvisited, down)
	}
	left := common.Point{current[0], current[1] - 1}
	if _, v := visited[left]; !v && left[1] >= 0 {
		unvisited = append(unvisited, left)
	}
	return unvisited
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()
	sideLength := int(math.Sqrt(float64(len(puzzleInput))))
	endPos := common.Point{sideLength - 1, sideLength - 1}
	visited := make(map[common.Point]bool)
	distances := make(map[common.Point]int)
	pq := priority_queue.New()
	for pos, _ := range puzzleInput {
		distances[pos] = math.MaxInt64
	}
	current := common.Point{0, 0}
	distances[current] = 0
	for {
		for _, neighbor := range getUnvisitedNeighbors(current, visited, sideLength) {
			distance := distances[current] + puzzleInput[neighbor]
			prevDistance := distances[neighbor]
			if distance < prevDistance {
				distances[neighbor] = distance
				pq.Enqueue(&priority_queue.Elem{Score: distance, Data: neighbor})
			}
		}
		visited[current] = true
		if current == endPos {
			break
		} else {
			deq, _ := pq.Dequeue()
			current = deq.Data
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