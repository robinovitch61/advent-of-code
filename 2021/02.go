package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2P1(directions []string) int {
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

func day2P2(directions []string) int {
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

func day2() {
	printDay(2)
	directions := readFile("./02_input.txt")
	fmt.Println(day2P1(directions))
	fmt.Println(day2P2(directions))
}
