package day19

import (
	"aoc/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Matrix [][]float64

func (m Matrix) Transpose() Matrix {
	t := Matrix{[]float64{0, 0, 0}, []float64{0, 0, 0}, []float64{0, 0, 0}}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			t[j][i] = m[i][j]
		}
	}
	return t
}

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

func (m Matrix) ToBeacon() Beacon {
	return Beacon{int(math.Round(m[0][0])), int(math.Round(m[1][0])), int(math.Round(m[2][0]))}
}

type Beacon struct {
	x, y, z int
}

func (b Beacon) ToMatrix() Matrix {
	return Matrix{
		[]float64{float64(b.x)},
		[]float64{float64(b.y)},
		[]float64{float64(b.z)},
	}
}

func (b Beacon) Add(o Beacon) Beacon {
	return Beacon{b.x + o.x, b.y + o.y, b.z + o.z}
}

func (b Beacon) Sub(o Beacon) Beacon {
	return Beacon{b.x - o.x, b.y - o.y, b.z - o.z}
}

func (b Beacon) Manhattan(o Beacon) int {
	return int(math.Abs(float64(b.x-o.x)) + math.Abs(float64(b.y-o.y)) + math.Abs(float64(b.z-o.z)))
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

func matMult(A, B Matrix) Matrix {
	if len(A) == 0 || len(B) == 0 || len(A[0]) == 0 || len(B[0]) == 0 || len(A[0]) != len(B) {
		fmt.Println(A)
		fmt.Println(B)
		panic("Cannot multiply")
	}
	var res Matrix
	for i := 0; i < len(A); i++ {
		row := make([]float64, len(B[0]))
		for j := 0; j < len(B[0]); j++ {
			var entry float64
			for k := 0; k < len(A[0]); k++ {
				entry += A[i][k] * B[k][j]
			}
			row[j] = entry
		}
		res = append(res, row)
	}
	return res
}

func testMatMult() {
	var A, B, expected Matrix
	A = Matrix{
		[]float64{1, 2},
		[]float64{3, 4},
	}
	B = Matrix{
		[]float64{1},
		[]float64{2},
	}
	expected = Matrix{
		[]float64{1*1 + 2*2},
		[]float64{3*1 + 4*2},
	}
	if res := matMult(A, B); !res.Equal(expected) {
		fmt.Println(res)
		fmt.Println(expected)
		panic("")
	}
	B = A
	expected = Matrix{
		[]float64{1*1 + 2*3, 1*2 + 2*4},
		[]float64{3*1 + 4*3, 3*2 + 4*4},
	}
	if res := matMult(A, B); !res.Equal(expected) {
		fmt.Println(res)
		fmt.Println(expected)
		panic("")
	}
}

func degToRad(deg int) float64 {
	return float64(deg) / 180 * math.Pi
}

func rotate(beacon Beacon, rotationMatrix Matrix) Beacon {
	return matMult(rotationMatrix, beacon.ToMatrix()).ToBeacon()
}

func rotateX(deg int) Matrix {
	rad := degToRad(deg)
	return Matrix{
		[]float64{1, 0, 0},
		[]float64{0, math.Cos(rad), math.Sin(rad)},
		[]float64{0, -math.Sin(rad), math.Cos(rad)},
	}
}

func rotateY(deg int) Matrix {
	rad := degToRad(deg)
	return Matrix{
		[]float64{math.Cos(rad), 0, math.Sin(rad)},
		[]float64{0, 1, 0},
		[]float64{-math.Sin(rad), 0, math.Cos(rad)},
	}
}

func rotateZ(deg int) Matrix {
	rad := degToRad(deg)
	return Matrix{
		[]float64{math.Cos(rad), math.Sin(rad), 0},
		[]float64{-math.Sin(rad), math.Cos(rad), 0},
		[]float64{0, 0, 1},
	}
}

func testRotate() {
	var beacon, rotated, reversed Beacon
	var matrix Matrix
	beacon = Beacon{1, 2, 3}
	for _, f := range []func(int) Matrix{rotateX, rotateY, rotateZ} {
		matrix = f(0)
		rotated = rotate(beacon, matrix)
		if rotated != beacon {
			fmt.Println(rotated)
			fmt.Println(beacon)
			panic("rotation by 0deg does something")
		}

		matrix = f(90)
		rotated = rotate(beacon, matrix)
		matrix = f(-270)
		reversed = rotate(beacon, matrix)
		if rotated != reversed {
			fmt.Println(rotated)
			fmt.Println(reversed)
			panic("rotation incorrect")
		}
	}
}

func allRotationMatrixes() []Matrix {
	var matrices []Matrix
	degrees := []int{0, 90, 180, 270}
	// point x-axis in each of 6 directions, then rotate about it in 4 90deg increments
	directions := []Matrix{rotateZ(0), rotateZ(90), rotateZ(180), rotateZ(270), rotateY(90), rotateY(270)}
	for _, direction := range directions {
		for _, deg := range degrees {
			matrices = append(matrices, matMult(direction, rotateX(deg)))
		}
	}
	return matrices
}

func testAllRotations() {
	beacon := Beacon{1, 2, 3}
	expected := map[Beacon]bool{
		Beacon{-3, -2, -1}: true,
		Beacon{-3, -1, 2}:  true,
		Beacon{-3, 1, -2}:  true,
		Beacon{-3, 2, 1}:   true,
		Beacon{-2, -3, 1}:  true,
		Beacon{-2, -1, -3}: true,
		Beacon{-2, 1, 3}:   true,
		Beacon{-2, 3, -1}:  true,
		Beacon{-1, -3, -2}: true,
		Beacon{-1, -2, 3}:  true,
		Beacon{-1, 2, -3}:  true,
		Beacon{-1, 3, 2}:   true,
		Beacon{1, -3, 2}:   true,
		Beacon{1, -2, -3}:  true,
		Beacon{1, 2, 3}:    true,
		Beacon{1, 3, -2}:   true,
		Beacon{2, -3, -1}:  true,
		Beacon{2, -1, 3}:   true,
		Beacon{2, 1, -3}:   true,
		Beacon{2, 3, 1}:    true,
		Beacon{3, -2, 1}:   true,
		Beacon{3, -1, -2}:  true,
		Beacon{3, 1, 2}:    true,
		Beacon{3, 2, -1}:   true,
	}

	obtained := make(map[Beacon]bool)
	for _, rotation := range allRotationMatrixes() {
		rotated := rotate(beacon, rotation)
		obtained[rotated] = true
	}
	for k := range expected {
		if check, exists := obtained[k]; !check || !exists {
			fmt.Println(k)
			panic("missing expected rotation")
		}
	}
}

func test() {
	testMatMult()
	testRotate()
	testAllRotations()
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

func removeScanner(scanners []Scanner, scanner Scanner) []Scanner {
	var res []Scanner
	for _, thisScanner := range scanners {
		if thisScanner.id != scanner.id {
			res = append(res, thisScanner)
		}
	}
	return res
}

func rotateBeacons(beacons []Beacon, rotation Matrix) []Beacon {
	var res []Beacon
	for _, beacon := range beacons {
		res = append(res, rotate(beacon, rotation))
	}
	return res
}

func matchesNumOffsets(knownBeacons map[Beacon]bool, beacons []Beacon, numMatching int) (Beacon, error) {
	offsets := make(map[Beacon]int)
	for b1 := range knownBeacons {
		for _, b2 := range beacons {
			diff := b1.Sub(b2)
			if count, exists := offsets[diff]; exists {
				if count+1 == numMatching {
					return diff, nil
				}
				offsets[diff]++
			} else {
				offsets[diff] = 1
			}
		}
	}
	return Beacon{}, fmt.Errorf("no match")
}

func matchScanner(knownBeacons map[Beacon]bool, scanner Scanner, rotationMatrices []Matrix) (Beacon, []Beacon) {
	for _, rotation := range rotationMatrices {
		rotatedBeacons := rotateBeacons(scanner.beacons, rotation)

		if scannerLoc, err := matchesNumOffsets(knownBeacons, rotatedBeacons, 12); err == nil {
			matchingBeacons := make([]Beacon, len(rotatedBeacons))
			for i, rotated := range rotatedBeacons {
				matchingBeacons[i] = rotated.Add(scannerLoc)
			}
			return scannerLoc, matchingBeacons
		}
	}
	return Beacon{}, nil
}

func solve(scanners []Scanner) (map[Beacon]bool, []Beacon) {
	rotationMatrices := allRotationMatrixes()

	allMatchedBeacons := make(map[Beacon]bool)
	for _, beacon := range scanners[0].beacons {
		allMatchedBeacons[beacon] = true
	}

	scannerLocs := []Beacon{{}}
	var unmatchedScanners []Scanner
	for _, scanner := range scanners[1:] {
		unmatchedScanners = append(unmatchedScanners, scanner)
	}

	for len(unmatchedScanners) > 0 {
		for _, unmatchedScanner := range unmatchedScanners {
			scannerLoc, matchedBeacons := matchScanner(allMatchedBeacons, unmatchedScanner, rotationMatrices)

			if matchedBeacons != nil {
				for _, beacon := range matchedBeacons {
					allMatchedBeacons[beacon] = true
				}
				scannerLocs = append(scannerLocs, scannerLoc)
				unmatchedScanners = removeScanner(unmatchedScanners, unmatchedScanner)
			}
		}
	}

	return allMatchedBeacons, scannerLocs
}

func both(scanners []Scanner) (int, int) {
	defer common.Time()()
	allMatchedBeacons, scannerLocs := solve(scanners)

	maxDist := 0
	for i := 0; i < len(scannerLocs); i++ {
		for j := 0; j < len(scannerLocs); j++ {
			if i == j {
				continue
			}
			dist := scannerLocs[i].Manhattan(scannerLocs[j])
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return len(allMatchedBeacons), maxDist
}

func Run() {
	common.PrintDay(19)
	test()
	input := common.ReadFile("19")
	scanners := parseInput(input)
	p1, p2 := both(scanners)
	fmt.Println(p1)
	fmt.Println(p2)
}
