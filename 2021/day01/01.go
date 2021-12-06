package day01

import (
	"aoc/common"
	"fmt"
	"strconv"
)

func p1(depths []int) int {
	numIncreases := 0
	for i, _ := range depths {
		if i > 0 && depths[i] > depths[i-1] {
			numIncreases += 1
		}
	}
	return numIncreases
}

func sum(input []int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum
}

func p2(depths []int, windowSize int) int {
	numIncreases := 0
	for i, _ := range depths {
		if i >= windowSize && sum(depths[i-windowSize:i]) < sum(depths[i-windowSize+1:i+1]) {
			numIncreases += 1
		}
	}
	return numIncreases
}

func Run() {
	common.PrintDay(1)
	depthStrings := common.ReadFile("01")
	var depths = make([]int, len(depthStrings), len(depthStrings))
	for i, depth := range depthStrings {
		intDepth, _ := strconv.Atoi(depth)
		depths[i] = intDepth
	}
	fmt.Println(p1(depths))
	fmt.Println(p2(depths, 3))
}
