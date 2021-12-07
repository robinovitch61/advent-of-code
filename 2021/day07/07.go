package day07

import (
	"aoc/common"
	"fmt"
	"math"
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

func minAndMax(array []int) (int, int) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func calcFuel(start []int, moveTo int) int {
	fuel := 0
	for _, pos := range start {
		fuel += int(math.Abs(float64(moveTo - pos)))
	}
	return fuel
}

func getGeoMap() map[int]int {
	size := 1500
	var geoMap = make(map[int]int, size)
	count := 0
	geoMap[0] = 0
	for i := 1; i < size; i++ {
		geoMap[i] = geoMap[i-1] + count
		count++
	}
	return geoMap
}

func calcFuelGeo(start []int, moveTo int) int {
	fuel := 0
	geoMap := getGeoMap()
	for _, pos := range start {
		fuel += geoMap[int(math.Abs(float64(moveTo-pos)))]
	}
	return fuel
}

func solve(start []int, calcFuel func([]int, int) int) int {
	min, max := minAndMax(start)
	minFuel := int(1e10)
	for i := min; i <= max; i++ {
		fuel := calcFuel(start, i)
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func p1(start []int) int {
	return solve(start, calcFuel)
}

func p2(start []int) int {
	return solve(start, calcFuelGeo)
}

func Run() {
	common.PrintDay(7)
	input := common.ReadFile("07")
	start := parseInput(input)
	fmt.Println(p1(start))
	fmt.Println(p2(start))
}
