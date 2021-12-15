package day14

import (
	"aoc/common"
	"fmt"
	"math"
	"strings"
)

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

func getPairGenerates(rules Rules) map[string][]string {
	generatedPairs := make(map[string][]string)
	for k, v := range rules {
		first := string(k[0]) + v
		second := v + string(k[1])
		generatedPairs[k] = []string{first, second}
	}
	return generatedPairs
}

func incrementMap(m map[string]int, k string, n int) {
	_, exists := m[k]
	if exists {
		m[k] += n
	} else {
		m[k] = n
	}
}

func solve(puzzleInput PuzzleInput, steps int) int {
	var polymer []string
	polymer = puzzleInput.initial
	pairGenerates := getPairGenerates(puzzleInput.rules)

	pairCount := make(map[string]int)
	letterCount := make(map[string]int)
	for i := 0; i < len(polymer); i++ {
		if i < len(polymer)-1 {
			pair := strings.Join(polymer[i:i+2], "")
			incrementMap(pairCount, pair, 1)
		}
		incrementMap(letterCount, polymer[i], 1)
	}

	var newPairCount map[string]int
	for step := 0; step < steps; step++ {
		newPairCount = make(map[string]int)
		for pair, count := range pairCount {
			newLetter := puzzleInput.rules[pair]
			incrementMap(letterCount, newLetter, count)
			newPairs := pairGenerates[pair]
			for _, newPair := range newPairs {
				incrementMap(newPairCount, newPair, count)
			}
		}
		pairCount = make(map[string]int)
		for k, v := range newPairCount {
			pairCount[k] = v
		}
	}

	min := math.MaxInt64
	max := 0
	for _, count := range letterCount {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	return max - min
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return solve(puzzleInput, 10)
}

func p2(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return solve(puzzleInput, 40)
}

func Run() {
	common.PrintDay(14)
	input := common.ReadFile("14")
	puzzleInput := parseInput(input)
	fmt.Println(p1(puzzleInput))
	fmt.Println(p2(puzzleInput))
}
