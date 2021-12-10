package day10

import (
	"aoc/common"
	"fmt"
	"strings"
)

func parseInput(input []string) [][]string {
	var lines [][]string
	for _, line := range input {
		lines = append(lines, strings.Split(line, ""))
	}
	return lines
}

func toOpenChar(char string) string {
	switch char {
	case ")":
		return "("
	case ">":
		return "<"
	case "}":
		return "{"
	case "]":
		return "["
	}
	panic("wtf")
}

func getClose(line []string, startIdx int) (int, string) {
	for fromStartIdx, nextChar := range line[startIdx+1:] {
		if strings.Contains("(<{[", nextChar) {
			return getClose(line, startIdx+fromStartIdx+1)
		} else {
			return startIdx + fromStartIdx + 1, nextChar
		}
	}
	return 0, ""
}

func validateLine(line []string) (string, bool) {
	startIdx := 0
	for startIdx < len(line)-1 {
		openChar := line[startIdx]
		closeIdx, closeChar := getClose(line, startIdx)
		if toOpenChar(closeChar) != openChar {
			return closeChar, false
		}
		startIdx = closeIdx + 1
	}
	return "", true
}

func p1(lines [][]string) int {
	score := 0
	scoreMap := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	for _, line := range lines {
		firstIllegalChar, isValid := validateLine(line)
		if !isValid {
			score += scoreMap[firstIllegalChar]
		}
	}
	return score
}

func p2(lines [][]string) int {
	return -1
}

func Run() {
	common.PrintDay(9)
	//input := common.ReadFile("10")
	//lines := parseInput(input)
	lines := [][]string{strings.Split("{([(<{}[<>[]}>{[]{[(<()>", "")}
	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}
