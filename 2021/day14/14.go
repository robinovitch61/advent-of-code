package day14

import (
	"aoc/common"
	"fmt"
	"math"
	"strings"
)

type Rule struct {
	left    string
	right   string
	between string
}

type Rules map[string]string

type PuzzleInput struct {
	initial []string
	rules   Rules
}

func parseInput(input []string) PuzzleInput {
	rules := make(map[string]string)
	for _, ruleRow := range input[2:] {
		split := strings.Split(ruleRow, " -> ")
		between := split[1]
		rules[split[0]] = between
	}
	return PuzzleInput{strings.Split(input[0], ""), rules}
}

func step(polymer []string, rules map[string]string) []string {
	i := 0
	for {
		if i == len(polymer)-1 {
			break
		}
		chunk := strings.Join(polymer[i:i+2], "")
		if between, exists := rules[chunk]; exists {
			betweenList := strings.Split(between, "")
			polymer = append(polymer[:i+1], append(betweenList, polymer[i+1:]...)...)
			i += len(betweenList)
		}
		i++
	}
	return polymer
}

func score(polymer []string) int {
	counts := make(map[string]int)
	for _, char := range polymer {
		_, exists := counts[char]
		if exists {
			counts[char]++
		} else {
			counts[char] = 1
		}
	}
	minCount := math.MaxInt64
	maxCount := 0
	for _, v := range counts {
		if v < minCount {
			minCount = v
		}
		if v > maxCount {
			maxCount = v
		}
	}
	return maxCount - minCount
}

func stepNTimes(initial []string, rules Rules, n int) []string {
	for num := 0; num < n; num++ {
		initial = step(initial, rules)
	}
	return initial
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()
	var polymer []string
	polymer = puzzleInput.initial
	polymer = stepNTimes(polymer, puzzleInput.rules, 10)
	return score(polymer)
}

func memoizeRules(rules Rules, size int) Rules {
	memo := make(map[string]string)
	for rule, _ := range rules {
		expanded := stepNTimes(strings.Split(rule, ""), rules, size)
		memo[rule] = strings.Join(expanded, "")
	}
	return memo
}

func p2(puzzleInput PuzzleInput) int {
	defer common.Time()()
	memoStepSize := 10
	memo := memoizeRules(puzzleInput.rules, memoStepSize)
	var polymer []string
	polymer = puzzleInput.initial
	polymer = stepNTimes(polymer, memo, 40/memoStepSize)
	return score(polymer)
}

func Run() {
	common.PrintDay(14)
	input := common.ReadFile("14")
	puzzleInput := parseInput(input)
	fmt.Println(p1(puzzleInput))
	fmt.Println(p2(puzzleInput))
}
