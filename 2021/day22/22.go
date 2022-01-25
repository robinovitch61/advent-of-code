package day22

import (
	"aoc/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	low, high int
}

type Instruction struct {
	isOn bool
	cube Cuboid
}

type PuzzleInput []Instruction

type Cuboid struct {
	xMin, xMax, yMin, yMax, zMin, zMax int
}

func (c Cuboid) Contains(o Cuboid) bool {
	return c.xMin <= o.xMin && c.xMax >= o.xMax &&
		c.yMin <= o.yMin && c.yMax >= o.yMax &&
		c.zMin <= o.zMin && c.zMax >= o.zMax
}

func (c Cuboid) Intersects(o Cuboid) bool {
	return c.xMin < o.xMax && c.xMax > o.xMin &&
		c.yMin < o.yMax && c.yMax > o.yMin &&
		c.zMin < o.zMax && c.zMax > o.zMin
}

func (c Cuboid) Volume() int {
	return (c.xMax - c.xMin) * (c.yMax - c.yMin) * (c.zMax - c.zMin)
}

func (c Cuboid) Subtract(o Cuboid) []Cuboid {
	if !c.Intersects(o) {
		return []Cuboid{c}
	} else if o.Contains(c) {
		return []Cuboid{}
	}

	xVals := []int{c.xMin, c.xMax, o.xMin, o.xMax}
	sort.Ints(xVals)
	yVals := []int{c.yMin, c.yMax, o.yMin, o.yMax}
	sort.Ints(yVals)
	zVals := []int{c.zMin, c.zMax, o.zMin, o.zMax}
	sort.Ints(zVals)

	var xMin, xMax, yMin, yMax, zMin, zMax int
	var cuboids []Cuboid
	for i := 0; i < len(xVals)-1; i++ {
		xMin = xVals[i]
		xMax = xVals[i+1]
		for j := 0; j < len(yVals)-1; j++ {
			yMin = yVals[j]
			yMax = yVals[j+1]
			for k := 0; k < len(zVals)-1; k++ {
				zMin = zVals[k]
				zMax = zVals[k+1]

				subCuboid := Cuboid{xMin, xMax, yMin, yMax, zMin, zMax}
				if c.Contains(subCuboid) && !subCuboid.Intersects(o) {
					cuboids = append(cuboids, subCuboid)
				}
			}
		}
	}
	return cuboids
}

func parseRange(str string) Range {
	str = str[2:]
	split := strings.Split(str, "..")
	low, _ := strconv.Atoi(split[0])
	high, _ := strconv.Atoi(split[1])
	return Range{low, high + 1}
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
	return Instruction{isOn, Cuboid{xRange.low, xRange.high, yRange.low, yRange.high, zRange.low, zRange.high}}
}

func parseInput(input []string) PuzzleInput {
	var instructions []Instruction
	for _, row := range input {
		instructions = append(instructions, parseInstruction(row))
	}
	return instructions
}

func clampRange(low, high int) Range {
	if low < -50 {
		low = -50
	}
	if high > 50 {
		high = 50
	}
	return Range{low, high}
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()

	cubes := make(map[[3]int]bool)
	for _, instr := range puzzleInput {
		xRange := clampRange(instr.cube.xMin, instr.cube.xMax)
		for x := xRange.low; x < xRange.high; x++ {
			yRange := clampRange(instr.cube.yMin, instr.cube.yMax)
			for y := yRange.low; y < yRange.high; y++ {
				zRange := clampRange(instr.cube.zMin, instr.cube.zMax)
				for z := zRange.low; z < zRange.high; z++ {
					cubes[[3]int{x, y, z}] = instr.isOn
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
	var cubes []Cuboid
	for _, instruction := range puzzleInput {
		var newCubes []Cuboid

		for _, cube := range cubes {
			for _, subCube := range cube.Subtract(instruction.cube) {
				newCubes = append(newCubes, subCube)
			}
		}

		if instruction.isOn {
			newCubes = append(newCubes, instruction.cube)
		}
		cubes = newCubes
	}
	numOn := 0
	for _, cube := range cubes {
		numOn += cube.Volume()
	}
	return numOn
}

func Run() {
	common.PrintDay(22)
	input := common.ReadFile("22")
	puzzleInput := parseInput(input)
	fmt.Println(p1(puzzleInput))
	fmt.Println(p2(puzzleInput))
}
