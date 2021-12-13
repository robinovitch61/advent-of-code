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
	defer common.Time()()
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

func solveEntry(entry Entry) map[string]string {
	crossedToReal := make(map[string]string)
	signals := entry.signals

	// 1 is the only display with 2 segments, C and F
	cOrF := strings.Split(signals[0], "")

	// 7 is the only display with 3 segments, A, C, and F, so we can determine A
	determineA := strings.Split(signals[1], "")
	for _, char := range determineA {
		if !strings.Contains(strings.Join(cOrF, ""), char) {
			crossedToReal[char] = "a"
		}
	}

	// 4 is the only display with 4 segments, B, C, D, and F
	var bOrD []string
	for _, char := range strings.Split(signals[2], "") {
		if !strings.Contains(strings.Join(cOrF, ""), char) {
			bOrD = append(bOrD, char)
		}
	}

	// there are three displays with 5 segments: 2, 3, and 5
	// the segments they share are A, D, and G
	// A is already determined and bOrD has overlap with another, which is D
	var middle string // used later
	fives := signals[3:6]
	commonInFives := commonChars(fives)
	for _, char := range commonInFives {
		if _, exists := crossedToReal[char]; !exists {
			if strings.Contains(strings.Join(bOrD, ""), char) {
				crossedToReal[char] = "d"
				middle = char
			} else {
				crossedToReal[char] = "g"
			}
		}
	}

	// we determined D above, so we can infer B
	for _, char := range bOrD {
		if _, exists := crossedToReal[char]; !exists {
			crossedToReal[char] = "b"
		}
	}

	// there are three displays with 6 segments: 0, 6, and 9
	// only 0 will have no D (middle) segment
	// 6 and 9 will have middle segments, and both have one of cOrF, which will be F
	sixes := signals[6:9]
	var sixAndNine []string
	for _, sixSegments := range sixes {
		if strings.Contains(sixSegments, middle) {
			sixAndNine = append(sixAndNine, sixSegments)
		}
	}
	for _, char := range cOrF {
		if strings.Contains(sixAndNine[0], char) && strings.Contains(sixAndNine[1], char) {
			crossedToReal[char] = "f"
		} else {
			crossedToReal[char] = "c"
		}
	}

	// we now have 6 of 7, so last one missing is E
	for _, char := range []string{"a", "b", "c", "d", "e", "f", "g"} {
		if _, exists := crossedToReal[char]; !exists {
			crossedToReal[char] = "e"
		}
	}

	return crossedToReal
}

func p2(entries []Entry) int {
	defer common.Time()()
	answer := 0
	for _, entry := range entries {
		solved := solveEntry(entry)

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
