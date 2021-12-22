package day22

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	low, high int
}

type Instruction struct {
	isOn                   bool
	xRange, yRange, zRange Range
}

type PuzzleInput []Instruction

type Cube struct {
	x, y, z int
}

func parseRange(str string) Range {
	str = str[2:]
	split := strings.Split(str, "..")
	low, _ := strconv.Atoi(split[0])
	high, _ := strconv.Atoi(split[1])
	return Range{low, high}
}

func parseInstruction(instruction string) Instruction {
	isOn := false
	split := strings.Split(instruction, " ")
	if split[0] == "on" {
		isOn = true
	}
	ranges := strings.Split(split[1], ",")
	xRange := parseRange(ranges[0])
	yRange := parseRange(ranges[1])
	zRange := parseRange(ranges[2])
	return Instruction{isOn, xRange, yRange, zRange}
}

func parseInput(input []string) PuzzleInput {
	var instructions []Instruction
	for _, row := range input {
		instructions = append(instructions, parseInstruction(row))
	}
	return instructions
}

func getRange(r Range) Range {
	var low, high int
	if r.low < -50 {
		low = -50
	} else {
		low = r.low
	}
	if r.high > 50 {
		high = 50
	} else {
		high = r.high
	}
	return Range{low, high}
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()

	cubes := make(map[Cube]bool)
	for _, instr := range puzzleInput {
		xRange := getRange(instr.xRange)
		for x := xRange.low; x <= xRange.high; x++ {
			yRange := getRange(instr.yRange)
			for y := yRange.low; y <= yRange.high; y++ {
				zRange := getRange(instr.zRange)
				for z := zRange.low; z <= zRange.high; z++ {
					cubes[Cube{x, y, z}] = instr.isOn
				}
			}
		}
	}

	count := 0
	for _, isOn := range cubes {
		if isOn {
			count++
		}
	}
	return count
}

func p2(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(22)
	input := common.ReadFile("22")
	puzzleInput := parseInput(input)
	fmt.Println(p1(puzzleInput))
	//fmt.Println(p2(puzzleInput))
}
