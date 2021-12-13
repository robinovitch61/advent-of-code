package day10

import (
	"aoc/common"
	"fmt"
	"math"
	"sort"
	"strings"
)

func parseInput(input []string) [][]string {
	var lines [][]string
	for _, line := range input {
		lines = append(lines, strings.Split(line, ""))
	}
	return lines
}

func toOppositeChar(char string) string {
	switch char {
	case ")":
		return "("
	case ">":
		return "<"
	case "}":
		return "{"
	case "]":
		return "["
	case "(":
		return ")"
	case "<":
		return ">"
	case "{":
		return "}"
	case "[":
		return "]"
	}
	panic("wtf")
}

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	} else {
		idx := len(*s) - 1
		last := (*s)[idx]
		*s = (*s)[:idx]
		return last, true
	}
}

func (s *Stack) Show() {
	fmt.Println(*s)
}

func validateLine(line []string) (string, bool, Stack) {
	startIdx := 0
	var stack Stack
	for startIdx < len(line) {
		val := line[startIdx]
		if strings.Contains("([<{", val) {
			stack.Push(val)
		} else {
			open, hasVal := stack.Pop()
			if !hasVal || toOppositeChar(val) != open {
				return val, false, stack
			}
		}
		startIdx++
	}
	return "", true, stack
}

var scoreMapCorrupted = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var scoreMapIncomplete = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func p1(lines [][]string) int {
	defer common.Time()()
	score := 0
	for _, line := range lines {
		firstIllegalChar, isValid, _ := validateLine(line)
		if !isValid {
			score += scoreMapCorrupted[firstIllegalChar]
		}
	}
	return score
}

func reverseStack(s Stack) []string {
	var reversed []string
	for {
		item, hasItem := s.Pop()
		if !hasItem {
			break
		}
		reversed = append(reversed, toOppositeChar(item))
	}
	return reversed
}

func getIncompleteScore(rest []string) int {
	score := 0
	for _, closed := range rest {
		score *= 5
		score += scoreMapIncomplete[closed]
	}
	return score
}

func p2(lines [][]string) int {
	defer common.Time()()
	var allScores []int
	for _, line := range lines {
		_, isValid, stack := validateLine(line)
		if isValid {
			rest := reverseStack(stack)
			score := getIncompleteScore(rest)
			allScores = append(allScores, score)
		}
	}
	sort.Ints(allScores)
	return allScores[int(math.Floor(float64((len(allScores)-1)/2)))]
}

func Run() {
	common.PrintDay(10)
	input := common.ReadFile("10")
	lines := parseInput(input)
	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}
