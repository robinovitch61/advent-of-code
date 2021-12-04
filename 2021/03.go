package main

import (
	"fmt"
	"strconv"
)

type Counts struct {
	zeroCount int
	oneCount  int
}

func getCounts(diagnostics []string, pos int) Counts {
	var zeroCount, oneCount int = 0, 0
	for _, diagnostic := range diagnostics {
		char := string(diagnostic[pos])
		if char == "0" {
			zeroCount++
		} else {
			oneCount++
		}
	}
	return Counts{zeroCount, oneCount}
}

func day3P1(diagnostics []string) int {
	var numBits = len(diagnostics[0])
	var gamma, epsilon string = "", ""
	for pos := 0; pos < numBits; pos++ {
		counts := getCounts(diagnostics, pos)
		if counts.zeroCount > counts.oneCount {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gammaDec, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gammaDec * epsilonDec)
}

func eliminate(diagnostics []string, pos int, mostCommon bool) string {
	var char string
	var remaining []string
	counts := getCounts(diagnostics, pos)
	if mostCommon {
		if counts.zeroCount > counts.oneCount {
			char = "0"
		} else {
			char = "1"
		}
	} else {
		if counts.zeroCount > counts.oneCount {
			char = "1"
		} else {
			char = "0"
		}
	}

	for _, diagnostic := range diagnostics {
		if char == string(diagnostic[pos]) {
			remaining = append(remaining, diagnostic)
		}
	}
	if len(remaining) == 1 {
		return remaining[0]
	} else {
		return eliminate(remaining, pos+1, mostCommon)
	}
}

func day3P2(diagnostics []string) int {
	oxygenGenerator := eliminate(diagnostics, 0, true)
	co2Scrubber := eliminate(diagnostics, 0, false)
	oxygenGeneratorDec, _ := strconv.ParseInt(oxygenGenerator, 2, 64)
	co2ScrubberDec, _ := strconv.ParseInt(co2Scrubber, 2, 64)
	return int(oxygenGeneratorDec * co2ScrubberDec)
}

func day3() {
	printDay(3)
	diagnostics := readFile("./03_input.txt")
	fmt.Println(day3P1(diagnostics))
	fmt.Println(day3P2(diagnostics))
}
