package day02

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

func p1(directions []string) int {
	defer common.Time()()
	var horizontal, depth int = 0, 0
	for _, direction := range directions {
		split := strings.Split(direction, " ")
		order := split[0]
		value, _ := strconv.Atoi(split[1])

		switch order {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	return horizontal * depth
}

func p2(directions []string) int {
	defer common.Time()()
	var horizontal, depth, aim int = 0, 0, 0
	for _, direction := range directions {
		split := strings.Split(direction, " ")
		order := split[0]
		value, _ := strconv.Atoi(split[1])

		switch order {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	return horizontal * depth
}

func Run() {
	common.PrintDay(2)
	directions := common.ReadFile("02")
	fmt.Println(p1(directions))
	fmt.Println(p2(directions))
}
