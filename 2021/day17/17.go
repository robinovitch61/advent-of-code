package day17

import (
	"aoc/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TargetArea struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func parseInput(input string) TargetArea {
	split := strings.Split(input, "target area: ")
	split2 := strings.Split(split[1], ", ")
	x, y := split2[0], split2[1]
	x = strings.TrimPrefix(x, "x=")
	y = strings.TrimPrefix(y, "y=")
	xSplit := strings.Split(x, "..")
	ySplit := strings.Split(y, "..")
	xMin, _ := strconv.Atoi(xSplit[0])
	xMax, _ := strconv.Atoi(xSplit[1])
	yMin, _ := strconv.Atoi(ySplit[0])
	yMax, _ := strconv.Atoi(ySplit[1])
	return TargetArea{xMin, xMax, yMin, yMax}
}

func inTargetArea(x int, y int, targetArea TargetArea) bool {
	return x >= targetArea.xMin && x <= targetArea.xMax && y >= targetArea.yMin && y <= targetArea.yMax
}

func canStillReachTargetArea(x int, y int, targetArea TargetArea) bool {
	return x <= targetArea.xMax && y >= targetArea.yMin
}

func getMaxYPos(xVel int, yVel int, targetArea TargetArea) int {
	xPos, yPos := 0, 0
	maxYPos := -math.MaxInt64
	for canStillReachTargetArea(xPos, yPos, targetArea) {
		if inTargetArea(xPos, yPos, targetArea) {
			return maxYPos
		}
		xPos += xVel
		yPos += yVel
		if yPos > maxYPos {
			maxYPos = yPos
		}
		if xVel > 0 {
			xVel--
		}
		if xVel < 0 {
			xVel++
		}
		yVel--
	}
	return -math.MaxInt64
}

func p1(targetArea TargetArea) int {
	defer common.Time()()
	maxYPos := -math.MaxInt64
	maxVel := 550
	for xVel := 1; xVel <= maxVel; xVel++ {
		for yVel := 1; yVel <= maxVel; yVel++ {
			maxYAchieved := getMaxYPos(xVel, yVel, targetArea)
			if maxYAchieved > maxYPos {
				maxYPos = maxYAchieved
			}
		}
	}
	return maxYPos
}

func p2(targetArea TargetArea) int {
	defer common.Time()()
	numAchieved := 0
	maxVel := 250
	for xVel := -maxVel; xVel <= maxVel; xVel++ {
		for yVel := -maxVel; yVel <= maxVel; yVel++ {
			maxYAchieved := getMaxYPos(xVel, yVel, targetArea)
			if maxYAchieved != -math.MaxInt64 {
				numAchieved++
			}
		}
	}
	return numAchieved
}

func Run() {
	common.PrintDay(17)
	input := common.ReadFile("17")[0]
	targetArea := parseInput(input)
	fmt.Println(p1(targetArea))
	fmt.Println(p2(targetArea))
}
