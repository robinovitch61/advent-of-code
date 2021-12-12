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
		fmt.Println(first, second)
		appendToMap(first, second, pathMap)
		appendToMap(second, first, pathMap)
	}
	return pathMap
}

func constainsString(strs []string, str string) bool {
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

func findPaths(pathMap PathMap, current string, smallVisited []string, path []string) [][]string {
	pathsTraversed := [][]string{}
	if current == "end" {
		return [][]string{append(path, current)}
	}
	for _, next := range pathMap[current] {
		if next != "start" && !constainsString(smallVisited, next) {
			var newSmallVisitied []string
			if isLower(next) {
				newSmallVisitied = append(smallVisited, next)
			} else {
				newSmallVisitied = smallVisited
			}
			pathsTraversed = append(pathsTraversed, findPaths(pathMap, next, newSmallVisitied, append(path, current))...)
		}
	}
	return pathsTraversed
}

func p1(pathMap PathMap) int {
	fmt.Println(pathMap)
	paths := findPaths(pathMap, "start", []string{}, []string{})
	fmt.Println(paths)
	return len(paths)
}

func p2(pathMap PathMap) int {
	return -1
}

func Run() {
	common.PrintDay(12)
	input := common.ReadFile("12")
	pathMap := parseInput(input)
	fmt.Println(p1(pathMap))
	fmt.Println(p2(pathMap))
}
