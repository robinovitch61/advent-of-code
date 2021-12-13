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

func getPaths(pathMap PathMap, current string, smallVisited []string, path []string, canVisitNext func([]string, string) bool) [][]string {
	var paths [][]string
	if current == "end" {
		return [][]string{append(path, current)}
	}
	for _, next := range pathMap[current] {
		if !isLower(next) || next != "start" && canVisitNext(smallVisited, next) {
			var newSmallVisited []string
			if isLower(next) {
				newSmallVisited = append(smallVisited, next)
			} else {
				newSmallVisited = smallVisited
			}
			paths = append(paths, getPaths(pathMap, next, newSmallVisited, append(path[:len(path):len(path)], current), canVisitNext)...)
		}
	}
	return paths
}

func solve(pathMap PathMap, canVisitNext func([]string, string) bool) int {
	defer common.Time()()
	paths := getPaths(pathMap, "start", []string{}, []string{}, canVisitNext)
	return len(paths)
}

func p1(pathMap PathMap) int {
	canVisitNext := func(smallVisited []string, next string) bool {
		return !containsString(smallVisited, next)
	}
	return solve(pathMap, canVisitNext)
}

func p2(pathMap PathMap) int {
	canVisitNext := func(smallVisited []string, next string) bool {
		return !containsString(smallVisited, next) || !hasDuplicate(smallVisited)
	}
	return solve(pathMap, canVisitNext)
}

func Run() {
	common.PrintDay(12)
	input := common.ReadFile("12")
	pathMap := parseInput(input)
	fmt.Println(p1(pathMap))
	fmt.Println(p2(pathMap))
}
