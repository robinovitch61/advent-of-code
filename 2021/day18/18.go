package day18

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	left   *Pair
	right  *Pair
	number int
}

type SnailFishNumber struct {
	root Pair
}

func (p Pair) Print() {
	if p.left == nil && p.right == nil {
		fmt.Print(p.number)
	} else {
		fmt.Print("[")
		p.left.Print()
		fmt.Print(",")
		p.right.Print()
		fmt.Print("]")
	}
}

func (s SnailFishNumber) Print() {
	left, right := s.root.left, s.root.right
	fmt.Print("[")
	left.Print()
	fmt.Print(",")
	right.Print()
	fmt.Println("]")
}

func getLeftRightStr(s string) (string, string) {
	var leftStr string
	var rightStr string
	s = s[1 : len(s)-1]
	if s[0] == '[' {
		var leftEndIdx int
		countOpen := 0
		for leftEndIdx = 0; !(s[leftEndIdx] == ']' && countOpen == 1); leftEndIdx++ {
			if s[leftEndIdx] == '[' {
				countOpen++
			} else if s[leftEndIdx] == ']' {
				countOpen--
			}
		}
		leftStr = s[0 : leftEndIdx+1]
		rightStr = s[leftEndIdx+2:]
	} else {
		split := strings.SplitN(s, ",", 2)
		leftStr = split[0]
		rightStr = split[1]
	}
	return leftStr, rightStr
}

func parseToPair(s string) Pair {
	if !strings.ContainsAny(s, "[],") {
		num, _ := strconv.Atoi(s)
		return Pair{nil, nil, num}
	} else {
		subLeft, subRight := getLeftRight(s)
		return Pair{&subLeft, &subRight, -1}
	}
}

func getLeftRight(s string) (Pair, Pair) {
	leftStr, rightStr := getLeftRightStr(s)
	left := parseToPair(leftStr)
	right := parseToPair(rightStr)
	return left, right
}

func fromString(s string) SnailFishNumber {
	root := Pair{}
	left, right := getLeftRight(s)
	root.left = &left
	root.right = &right
	snail := SnailFishNumber{root}
	return snail
}

func parseInput(input []string) []SnailFishNumber {
	var nums []SnailFishNumber
	for _, s := range input {
		nums = append(nums, fromString(s))
	}
	return nums
}

func p1(nums []SnailFishNumber) int {
	defer common.Time()()
	for _, num := range nums {
		num.Print()
	}
	return -1
}

func p2(nums []SnailFishNumber) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(18)
	input := common.ReadFile("18")
	nums := parseInput(input)
	fmt.Println(p1(nums))
	//fmt.Println(p2(nums))
}
