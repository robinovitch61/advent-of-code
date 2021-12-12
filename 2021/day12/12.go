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

func containsStringNTimes(strs []string, str string, times int) bool {
	count := 0
	for _, item := range strs {
		if str == item {
			count++
			if count >= times {
				return true
			}
		}
	}
	return false
}

func isLower(str string) bool {
	return strings.ToUpper(str) != str
}

func hasNonUniqueValue(smallVisited []string) bool {
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

func findPaths(pathMap PathMap, current string, smallVisited []string, path []string, canVisitNext func([]string, string) bool) [][]string {
	var pathsTraversed [][]string
	if current == "end" {
		return [][]string{append(path, current)}
	}
	for _, next := range pathMap[current] {
		if next != "start" && canVisitNext(smallVisited, next) {
			var newSmallVisited []string
			if isLower(next) {
				newSmallVisited = append(smallVisited, next)
			} else {
				newSmallVisited = smallVisited
			}
			pathsTraversed = append(pathsTraversed, findPaths(pathMap, next, newSmallVisited, append(path, current), canVisitNext)...)
		}
	}
	return pathsTraversed
}

func p1(pathMap PathMap) int {
	canVisitNext := func(smallVisited []string, next string) bool {
		return !containsStringNTimes(smallVisited, next, 1)
	}
	paths := findPaths(pathMap, "start", []string{}, []string{}, canVisitNext)
	return len(paths)
}

func p2(pathMap PathMap) int {
	canVisitNext := func(smallVisited []string, next string) bool {
		if !isLower(next) {
			return true
		}
		if next == "end" {
			return true
		}
		return !containsStringNTimes(smallVisited, next, 1) || !hasNonUniqueValue(smallVisited)
	}

	paths := findPaths(pathMap, "start", []string{}, []string{}, canVisitNext)
	return len(paths)
}

func Run() {
	common.PrintDay(12)
	input := common.ReadFile("12")
	pathMap := parseInput(input)
	fmt.Println(p1(pathMap))
	//fmt.Println(p2(pathMap)) // answer is right here but paths are wrong...go slices arg
}
