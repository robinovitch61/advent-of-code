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

func (p Pair) printHelper() {
	if p.left == nil && p.right == nil {
		fmt.Print(p.number)
	} else {
		fmt.Print("[")
		p.left.printHelper()
		fmt.Print(",")
		p.right.printHelper()
		fmt.Print("]")
	}
}

func (p Pair) Print() {
	p.printHelper()
	fmt.Println()
}

func (p Pair) IsNumber() bool {
	return p.left == nil && p.right == nil
}

func (p Pair) Add(o Pair) Pair {
	return Pair{&p, &o, nil, -1}
}

func (p Pair) Reduce() Pair {
	current := p
	visited := make(map[Pair]bool)
	//didAction := false
	for {
		if !visited[*current.left] && !current.left.IsNumber() {
			current = *current.left
		} else if !visited[*current.right] && !current.right.IsNumber() {
			current = *current.right
		} else {
			if current == p && (visited[*current.right] || current.right.IsNumber()) {
				break
			}
			visited[current] = true
			current = *current.parent
		}
	}
	return p
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

func PairFromString(s string) Pair {
	var pair Pair
	left, right := parseLeftRightPairs(s, &pair)
	pair.left = &left
	pair.right = &right
	return pair
}

func parseInput(input []string) []Pair {
	var nums []Pair
	for _, s := range input {
		nums = append(nums, PairFromString(s))
	}
	return nums
}

func p1(nums []Pair) int {
	defer common.Time()()
	for _, num := range nums {
		num.Print()
		num.Reduce()
	}
	//PairFromString("[1,2]").Add(PairFromString("[[3,4],5]")).Print()
	return -1
}

func p2(nums []Pair) int {
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
