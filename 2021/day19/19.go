package day19

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type Beacon struct {
	x, y, z int
}

func beaconFromString(s string) Beacon {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	z, _ := strconv.Atoi(split[2])
	return Beacon{x, y, z}
}

type Scanner struct {
	id      int
	beacons []Beacon
}

func parseInput(input []string) []Scanner {
	text := strings.Join(input, "\n")
	blocks := strings.Split(text, "\n\n")

	var scanners []Scanner
	for _, block := range blocks {
		split := strings.Split(block, "\n")
		header := split[0]
		id, _ := strconv.Atoi(strings.Split(header, " ")[2])
		var beacons []Beacon
		for _, beacon := range split[1:] {
			beacons = append(beacons, beaconFromString(beacon))
		}
		scanners = append(scanners, Scanner{id, beacons})
	}
	return scanners
}

func p1(scanners []Scanner) int {
	defer common.Time()()
	fmt.Println(scanners[0])
	return -1
}

func p2(scanners []Scanner) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(19)
	input := common.ReadFile("19")
	scanners := parseInput(input)
	fmt.Println(p1(scanners))
	fmt.Println(p2(scanners))
}
