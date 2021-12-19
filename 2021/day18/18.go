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
	parent *Pair
	number int
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

func (p Pair) IsNumber() bool {
	return p.left == nil && p.right == nil
}

type SnailFishNumber struct {
	root Pair
}

func (s SnailFishNumber) Print() {
	left, right := s.root.left, s.root.right
	fmt.Print("[")
	left.Print()
	fmt.Print(",")
	right.Print()
	fmt.Println("]")
}

func (s SnailFishNumber) Add(o SnailFishNumber) SnailFishNumber {
	return SnailFishNumber{Pair{&s.root, &o.root, nil, -1}}
}

func (s SnailFishNumber) Reduce() SnailFishNumber {
	current := s.root
	visited := make(map[Pair]bool)
	//didAction := false
	for {
		if !(current.left == nil) && !visited[*current.left] && !current.left.IsNumber() {
			current = *current.left
		} else if !(current.right == nil) && !visited[*current.right] && !current.right.IsNumber() {
			current = *current.right
		} else {
			if current == s.root && (visited[*current.right] || current.right.IsNumber()) {
				break
			}
			visited[current] = true
			current = *current.parent
		}
	}
	return s
}

func parseToPair(s string, parent *Pair) Pair {
	var pair Pair
	pair.parent = parent
	if !strings.ContainsAny(s, "[],") {
		num, _ := strconv.Atoi(s)
		pair.number = num
		return pair
	} else {
		subLeft, subRight := parseLeftRightPairs(s, &pair)
		pair.left = &subLeft
		pair.right = &subRight
		return pair
	}
}

func splitLeftRightStrings(s string) (string, string) {
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

func parseLeftRightPairs(s string, parent *Pair) (Pair, Pair) {
	leftStr, rightStr := splitLeftRightStrings(s)
	left := parseToPair(leftStr, parent)
	left.parent = parent
	right := parseToPair(rightStr, parent)
	right.parent = parent
	return left, right
}

func snailFishNumberFromString(s string) SnailFishNumber {
	var root Pair
	left, right := parseLeftRightPairs(s, &root)
	root.left = &left
	root.right = &right
	snail := SnailFishNumber{root}
	return snail
}

func parseInput(input []string) []SnailFishNumber {
	var nums []SnailFishNumber
	for _, s := range input {
		nums = append(nums, snailFishNumberFromString(s))
	}
	return nums
}

func p1(nums []SnailFishNumber) int {
	defer common.Time()()
	for _, num := range nums {
		num.Print()
		num.Reduce()
	}
	//snailFishNumberFromString("[1,2]").Add(snailFishNumberFromString("[[3,4],5]")).Print()
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
