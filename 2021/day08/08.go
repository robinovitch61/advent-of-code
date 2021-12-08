package day08

import (
	"aoc/common"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Entry struct {
	signals [10]string // items are sorted by chars and array in order of increasing length
	outputs [4]string
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func parseInput(input []string) []Entry {
	var entries []Entry
	for _, entry := range input {
		split := strings.Split(entry, " | ")
		var signals [10]string
		var outputs [4]string
		first, second := split[0], split[1]
		signalsSlice := strings.Split(first, " ")
		outputsSlice := strings.Split(second, " ")
		for idx, signal := range signalsSlice {
			signalsSlice[idx] = sortString(signal)
		}
		sort.Slice(signalsSlice, func(i, j int) bool {
			return len(signalsSlice[i]) < len(signalsSlice[j])
		})
		copy(signals[:], signalsSlice)
		copy(outputs[:], outputsSlice)
		newEntry := Entry{signals, outputs}
		entries = append(entries, newEntry)
	}
	return entries
}

func p1(entries []Entry) int {
	uniqueDigits := 0
	uniqueCounts := [4]int{2, 4, 3, 7}
	for _, entry := range entries {
		for _, output := range entry.outputs {
			for _, unique := range uniqueCounts {
				if len(output) == unique {
					uniqueDigits++
				}
			}
		}
	}
	return uniqueDigits
}

func segmentsToNum(segments string) int {
	switch segments {
	case "abcefg":
		return 0
	case "cf":
		return 1
	case "acdeg":
		return 2
	case "acdfg":
		return 3
	case "bcdf":
		return 4
	case "abdfg":
		return 5
	case "abdefg":
		return 6
	case "acf":
		return 7
	case "abcdefg":
		return 8
	case "abcdfg":
		return 9
	}
	fmt.Println(segments)
	panic("wtf")
}

func possibleCombinationsOfReal() [10]string {
	zero := "abcefg"
	one := "cf"
	two := "acdeg"
	three := "acdfg"
	four := "bcdf"
	five := "abdfg"
	six := "abdefg"
	seven := "acf"
	eight := "abcdefg"
	nine := "abcdfg"
	return [10]string{zero, one, two, three, four, five, six, seven, eight, nine}
}

func combinationsOfSameLength(signal string) []string {
	var combos []string
	for _, combo := range possibleCombinationsOfReal() {
		if len(signal) == len(combo) {
			combos = append(combos, combo)
		}
	}
	return combos
}

func segmentOptions() []string {
	return []string{"a", "b", "c", "d", "e", "f", "g"}
}

func possibleCrossedToReal() map[string][]string {
	possible := make(map[string][]string)
	for _, segment := range segmentOptions() {
		possible[segment] = segmentOptions()
	}
	return possible
}

func atLeastOneContains(options []string, substr string) bool {
	for _, option := range options {
		if strings.Contains(option, substr) {
			return true
		}
	}
	return false
}

func commonChars(options []string) []string {
	charCount := make(map[string]int, 26)
	for _, option := range options {
		for i := 0; i < len(option); i++ {
			charCount[string(option[i])]++
		}
	}

	var result []string
	for k, v := range charCount {
		if v == len(options) {
			result = append(result, k)
		}
	}
	return result
}

func updateCrossedToReal(crossedToReal map[string][]string, sortedSignal string, possibleCombos []string) {
	// LiNeAr AlGebRa
	for k, v := range crossedToReal {
		if len(v) == 1 {
			eliminateFromOthers := v[0]
			for _, segmentChar := range segmentOptions() {
				if k != segmentChar {
					// can eliminate from all other options
					var eliminatedRealOptions []string
					for _, oldOption := range crossedToReal[segmentChar] {
						if oldOption != eliminateFromOthers {
							eliminatedRealOptions = append(eliminatedRealOptions, oldOption)
						}
					}
					crossedToReal[segmentChar] = eliminatedRealOptions
				}
			}
		}
	}

	for i := 0; i < len(sortedSignal); i++ {
		signalChar := string(sortedSignal[i])

		// pair down the existing options based on the possible combos
		var newRealOptions []string
		for _, oldOption := range crossedToReal[signalChar] {
			if atLeastOneContains(possibleCombos, oldOption) {
				newRealOptions = append(newRealOptions, oldOption)
			}
		}
		crossedToReal[signalChar] = newRealOptions
	}

	commonChars := commonChars(possibleCombos)
	for _, segmentChar := range segmentOptions() {
		if !strings.Contains(sortedSignal, segmentChar) {
			var eliminatedRealOptions []string
			for _, oldOption := range crossedToReal[segmentChar] {
				if !strings.Contains(strings.Join(commonChars, ""), oldOption) {
					eliminatedRealOptions = append(eliminatedRealOptions, oldOption)
				}
			}
			crossedToReal[segmentChar] = eliminatedRealOptions
		}
	}

}

func solveEntry(entry Entry) map[string]string {
	crossedToReal := possibleCrossedToReal()
	for _, signal := range entry.signals {
		sortedSignal := sortString(signal)
		possibleCombos := combinationsOfSameLength(sortedSignal)
		updateCrossedToReal(crossedToReal, sortedSignal, possibleCombos)
	}
	result := make(map[string]string)
	for k, v := range crossedToReal {
		result[k] = v[0]
	}
	return result
}

func p2(entries []Entry) int {
	answer := 0
	for _, entry := range entries {
		solved := solveEntry(entry)
		fmt.Println(solved)

		for idx, output := range entry.outputs {
			result := ""
			for i := 0; i < len(output); i++ {
				result += solved[string(output[i])]
			}
			num := segmentsToNum(sortString(result)) * int(math.Pow(10, float64(4-idx-1)))
			answer += num
		}
	}
	return answer
}

func Run() {
	common.PrintDay(8)
	input := common.ReadFile("08")
	entries := parseInput(input)
	fmt.Println(p1(entries))
	fmt.Println(p2(entries))
}
