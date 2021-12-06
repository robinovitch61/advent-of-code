package day06

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input []string) []int {
	split := strings.Split(input[0], ",")
	var nums []int
	for _, str := range split {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	return nums
}

func solveForDays(start []int, days int) int {
	numFishWithAge := make(map[int]int)
	for _, left := range start {
		if numFishWithAge[left] == 0 {
			numFishWithAge[left] = 1
		} else {
			numFishWithAge[left]++
		}
	}

	for day := 0; day < days; day++ {
		numRegen := numFishWithAge[0]
		for left := 1; left < 9; left++ {
			numFishWithAge[left-1] = numFishWithAge[left]
		}
		numFishWithAge[6] += numRegen
		numFishWithAge[8] = numRegen
	}

	numFish := 0
	for _, count := range numFishWithAge {
		numFish += count
	}
	return numFish
}

func p1(start []int) int {
	return solveForDays(start, 80)
}

func p2(start []int) int {
	return solveForDays(start, 256)
}

func Run() {
	common.PrintDay(6)
	input := common.ReadFile("06")
	start := parseInput(input)
	fmt.Println(p1(start))
	fmt.Println(p2(start))
}
