package day09

import (
	"aoc/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input []string) [][]int {
	var heights [][]int
	for _, row := range input {
		split := strings.Split(row, "")
		var ints []int
		for _, item := range split {
			entry, _ := strconv.Atoi(item)
			ints = append(ints, entry)
		}
		heights = append(heights, ints)
	}
	return heights
}

func getIsLowPoint(heights [][]int, i int, j int) (int, bool) {
	candidate := heights[i][j]
	if i > 0 && heights[i-1][j] <= candidate {
		return candidate, false
	} else if j > 0 && heights[i][j-1] <= candidate {
		return candidate, false
	} else if i < len(heights)-1 && heights[i+1][j] <= candidate {
		return candidate, false
	} else if j < len(heights[0])-1 && heights[i][j+1] <= candidate {
		return candidate, false
	}
	return candidate, true
}

func p1(heights [][]int) int {
	var lowPoints []int
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if lowPoint, isLowPoint := getIsLowPoint(heights, i, j); isLowPoint {
				lowPoints = append(lowPoints, lowPoint)
			}
		}
	}

	riskLevelSum := 0
	for _, lowPoint := range lowPoints {
		riskLevelSum += lowPoint + 1
	}
	return riskLevelSum
}

func getBasinSize(heights [][]int, i int, j int) int {
	if heights[i][j] == 9 {
		panic("wtf")
	}
	heights[i][j] = 9
	basinSize := 1
	if i > 0 && heights[i-1][j] != 9 {
		basinSize += getBasinSize(heights, i-1, j)
	}
	if j > 0 && heights[i][j-1] != 9 {
		basinSize += getBasinSize(heights, i, j-1)
	}
	if i < len(heights)-1 && heights[i+1][j] != 9 {
		basinSize += getBasinSize(heights, i+1, j)
	}
	if j < len(heights[0])-1 && heights[i][j+1] != 9 {
		basinSize += getBasinSize(heights, i, j+1)
	}
	return basinSize
}

func p2(heights [][]int) int {
	var basinSizes []int
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if _, isLowPoint := getIsLowPoint(heights, i, j); isLowPoint {
				basinSizes = append(basinSizes, getBasinSize(heights, i, j))
			}
		}
	}

	sort.Ints(basinSizes)
	result := 1
	for _, size := range basinSizes[len(basinSizes)-3:] {
		result *= size
	}
	return result
}

func Run() {
	common.PrintDay(9)
	input := common.ReadFile("09")
	heights := parseInput(input)
	fmt.Println(p1(heights))
	fmt.Println(p2(heights))
}
