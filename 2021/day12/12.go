package day12

import (
	"aoc/common"
	"fmt"
	"strings"
)

type PathMap map[string][]string

func appendToMap(first string, second string, pathMap PathMap) {
	if v, exists := pathMap[first]; exists {
		pathMap[first] = append(v, second)
	} else {
		pathMap[first] = []string{second}
	}
}

func parseInput(input []string) PathMap {
	pathMap := make(PathMap)
	for _, row := range input {
		split := strings.Split(row, "-")
		first, second := split[0], split[1]
		appendToMap(first, second, pathMap)
		appendToMap(second, first, pathMap)
	}
	return pathMap
}

func containsString(strs []string, str string) bool {
	for _, item := range strs {
		if str == item {
			return true
		}
	}
	return false
}

func isLower(str string) bool {
	return strings.ToUpper(str) != str
}

func hasDuplicate(smallVisited []string) bool {
	countMap := make(map[string]int)
	for _, v := range smallVisited {
		_, exists := countMap[v]
		if exists {
			return true
		} else {
			countMap[v] = 1
		}
	}
	return false
}

func countPaths(pathMap PathMap, current string, smallVisited []string, canVisitNext func([]string, string) bool) int {
	var pathCount int
	if current == "end" {
		return 1
	}
	for _, next := range pathMap[current] {
		if !isLower(next) || next != "start" && canVisitNext(smallVisited, next) {
			var newSmallVisited []string
			if isLower(next) {
				newSmallVisited = append(smallVisited, next)
			} else {
				newSmallVisited = smallVisited
			}
			pathCount += countPaths(pathMap, next, newSmallVisited, canVisitNext)
		}
	}
	return pathCount
}

func p1(pathMap PathMap) int {
	defer common.Time()()
	canVisitNext := func(smallVisited []string, next string) bool {
		return !containsString(smallVisited, next)
	}
	pathCount := countPaths(pathMap, "start", []string{}, canVisitNext)
	return pathCount
}

func p2(pathMap PathMap) int {
	defer common.Time()()
	canVisitNext := func(smallVisited []string, next string) bool {
		return !containsString(smallVisited, next) || !hasDuplicate(smallVisited)
	}
	pathCount := countPaths(pathMap, "start", []string{}, canVisitNext)
	return pathCount
}

func Run() {
	common.PrintDay(12)
	input := common.ReadFile("12")
	pathMap := parseInput(input)
	fmt.Println(p1(pathMap))
	fmt.Println(p2(pathMap))
}
