package day11

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input []string) [][]int {
	var energyLevels [][]int
	for _, line := range input {
		energyRow := strings.Split(line, "")
		ints := make([]int, len(energyRow))
		for i, energy := range energyRow {
			val, _ := strconv.Atoi(energy)
			ints[i] = val
		}
		energyLevels = append(energyLevels, ints)
	}
	return energyLevels
}

func incrementAll(energyLevels [][]int, by int) {
	for i, row := range energyLevels {
		for j, _ := range row {
			energyLevels[i][j]++
		}
	}
}

func inFlashed(flashed [][2]int, i int, j int) bool {
	for _, flashLoc := range flashed {
		if flashLoc[0] == i && flashLoc[1] == j {
			return true
		}
	}
	return false
}

func incrementAdjacentToFlashedThisRound(energyLevels [][]int, flashedThisRound [][2]int) {
	maxWidth := len(energyLevels[0]) - 1
	maxHeight := len(energyLevels) - 1
	for _, flash := range flashedThisRound {
		i, j := flash[0], flash[1]
		// up
		if i-1 >= 0 {
			energyLevels[i-1][j]++
		}
		// left
		if j-1 >= 0 {
			energyLevels[i][j-1]++
		}
		// down
		if i+1 <= maxHeight {
			energyLevels[i+1][j]++
		}
		// right
		if j+1 <= maxWidth {
			energyLevels[i][j+1]++
		}
		// up left
		if i-1 >= 0 && j-1 >= 0 {
			energyLevels[i-1][j-1]++
		}
		// up right
		if i-1 >= 0 && j+1 <= maxWidth {
			energyLevels[i-1][j+1]++
		}
		// down left
		if i+1 <= maxHeight && j-1 >= 0 {
			energyLevels[i+1][j-1]++
		}
		// down right
		if i+1 <= maxHeight && j+1 <= maxWidth {
			energyLevels[i+1][j+1]++
		}
	}
}

func step(energyLevels [][]int) int {
	flashes := 0
	incrementAll(energyLevels, 1)
	var flashedThisStep [][2]int

	for {
		flashedThisRound := [][2]int{}
		for i, row := range energyLevels {
			for j, _ := range row {
				val := energyLevels[i][j]
				if val > 9 && !inFlashed(flashedThisStep, i, j) {
					flashes++
					flashedThisRound = append(flashedThisRound, [2]int{i, j})
				}
			}
		}
		if len(flashedThisRound) == 0 {
			break
		}
		flashedThisStep = append(flashedThisStep, flashedThisRound...)
		incrementAdjacentToFlashedThisRound(energyLevels, flashedThisRound)
	}
	// finally, set all flashed this step to zero
	for _, flash := range flashedThisStep {
		energyLevels[flash[0]][flash[1]] = 0
	}
	return flashes
}

func p1(input []string) int {
	energyLevels := parseInput(input)
	totalFlashes := 0
	numSteps := 100
	for s := 0; s < numSteps; s++ {
		totalFlashes += step(energyLevels)
	}
	return totalFlashes
}

func synced(energyLevels [][]int) bool {
	for i, row := range energyLevels {
		for j, _ := range row {
			if energyLevels[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func p2(input []string) int {
	energyLevels := parseInput(input)
	s := 0
	for {
		step(energyLevels)
		s++
		if synced(energyLevels) {
			return s
		}
	}
}

func Run() {
	common.PrintDay(11)
	input := common.ReadFile("11")
	fmt.Println(p1(input))
	fmt.Println(p2(input))
}
