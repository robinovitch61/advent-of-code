package day20

import (
	"aoc/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Pair struct {
	x, y int
}

type Algo map[int]bool

type Image map[Pair]bool

func (i Image) Print() {
	const margin = 3
	minMax := getMinMax(i)
	var row string
	for x := minMax.minX - margin; x <= minMax.maxX+margin; x++ {
		row = ""
		for y := minMax.minY - margin; y <= minMax.maxY+margin; y++ {
			if i[Pair{x, y}] {
				row += "#"
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
}

type PuzzleInput struct {
	algo  Algo
	image Image
}

type MinMax struct {
	minX, maxX, minY, maxY int
}

const margin = 3

func solve(puzzleInput PuzzleInput, numEnhance int) int {
	image := puzzleInput.image
	for enhanced := 0; enhanced < numEnhance; enhanced++ {
		image = enhance(puzzleInput.algo, image)
		if enhanced%2 == 1 {
			image = trim(image, 2*margin+1)
		}
	}
	count := 0
	for _, isLight := range image {
		if isLight {
			count++
		}
	}
	return count
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return solve(puzzleInput, 2)
}

func p2(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return solve(puzzleInput, 50)
}

func Run() {
	common.PrintDay(20)
	input := common.ReadFile("20")
	puzzle := parseInput(input)
	fmt.Println(p1(puzzle))
	fmt.Println(p2(puzzle))
}

func trim(image Image, num int) Image {
	minMax := getMinMax(image)
	newImage := make(Image)
	for pair, light := range image {
		if shouldKeep(pair, minMax, num) {
			newImage[pair] = light
		}
	}
	return newImage
}

func shouldKeep(pair Pair, minMax MinMax, num int) bool {
	return pair.x > minMax.minX+num/2 && pair.x < minMax.maxX-num/2 && pair.y > minMax.minY+num/2 && pair.y < minMax.maxY-num/2
}

func enhance(algo Algo, image Image) Image {
	newImage := make(Image)
	for _, pair := range getCandidates(image) {
		newImage[pair] = algo[getDec(pair, image)]
	}
	return newImage
}

func getCandidates(image Image) []Pair {
	var pairs []Pair
	minMax := getMinMax(image)
	for i := minMax.minX - margin; i <= minMax.maxX+margin; i++ {
		for j := minMax.minY - margin; j <= minMax.maxY+margin; j++ {
			pairs = append(pairs, Pair{i, j})
		}
	}
	return pairs
}

func getMinMax(image Image) MinMax {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := -math.MaxInt64, -math.MaxInt64
	for pair := range image {
		if pair.x < minX {
			minX = pair.x
		}
		if pair.y < minY {
			minY = pair.y
		}
		if pair.x > maxX {
			maxX = pair.x
		}
		if pair.y > maxY {
			maxY = pair.y
		}
	}
	return MinMax{minX, maxX, minY, maxY}
}

func getDec(pair Pair, image Image) int {
	binary := ""
	for _, surr := range getSurrounding(pair) {
		if image[surr] {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	dec, _ := strconv.ParseInt(binary, 2, 64)
	return int(dec)
}

func getSurrounding(pair Pair) []Pair {
	var pairs []Pair
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			pairs = append(pairs, Pair{pair.x + i, pair.y + j})
		}
	}
	return pairs
}

func getIsLight(char string) bool {
	if char == "#" {
		return true
	} else {
		return false
	}
}

func parseInput(input []string) PuzzleInput {
	i := 0
	algo := make(Algo)
	image := make(Image)
	x, y := 0, 0
	isImage := false
	for _, row := range input {
		if !isImage {
			for _, char := range strings.Split(row, "") {
				algo[i] = getIsLight(char)
				i++
			}
		} else {
			y = 0
			for _, char := range strings.Split(row, "") {
				image[Pair{x, y}] = getIsLight(char)
				y++
			}
			x++
		}
		if strings.TrimSpace(row) == "" {
			isImage = true
		}
	}
	return PuzzleInput{algo, image}
}
