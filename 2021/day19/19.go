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

type Matrix [][]int

func (m Matrix) Equal(other Matrix) bool {
	if len(m) != len(other) || len(m[0]) != len(other[0]) {
		return false
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] != other[i][j] {
				return false
			}
		}
	}
	return true
}

func matMult(A, B Matrix) (Matrix, error) {
	if len(A) == 0 || len(B) == 0 || len(A[0]) == 0 || len(B[0]) == 0 || len(A[0]) != len(B) {
		return Matrix{}, fmt.Errorf("cannot multiply")
	}
	var res Matrix
	for i := 0; i < len(A); i++ {
		var row []int
		for j := 0; j < len(B[0]); j++ {
			entry := 0
			for k := 0; k < len(A[0]); k++ {
				entry += A[i][k] * B[k][j]
			}
			row = append(row, entry)
		}
		res = append(res, row)
	}
	return res, nil
}

func testMatMult() {
	var A, B, expected Matrix
	A = Matrix{
		[]int{1, 2},
		[]int{3, 4},
	}
	B = Matrix{
		[]int{1},
		[]int{2},
	}
	expected = Matrix{
		// i k   k j   i k   k j
		// 0,0 * 0,0 + 0,1 * 1,0
		[]int{1*1 + 2*2},
		// i k   k j   i k   k j
		// 1,0 * 0,0 + 1,1 * 1,0
		[]int{3*1 + 4*2},
	}
	if res, err := matMult(A, B); err != nil || !res.Equal(expected) {
		fmt.Println(res)
		fmt.Println(expected)
		panic("")
	}
	B = A
	expected = Matrix{
		[]int{1*1 + 2*3, 1*2 + 2*4},
		[]int{3*1 + 4*3, 3*2 + 4*4},
	}
	if res, err := matMult(A, B); err != nil || !res.Equal(expected) {
		fmt.Println(res)
		fmt.Println(expected)
		panic("")
	}
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
	return -1
}

func p2(scanners []Scanner) int {
	defer common.Time()()
	return -1
}

func test() {
	testMatMult()
}

func Run() {
	common.PrintDay(19)
	test()
	input := common.ReadFile("19")
	scanners := parseInput(input)
	fmt.Println(p1(scanners))
	fmt.Println(p2(scanners))
}
