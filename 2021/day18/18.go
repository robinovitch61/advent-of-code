package day18

import (
	"aoc/common"
	"fmt"
	"math"
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

func shouldSplit(pair *Pair) bool {
	return pair.IsNumber() && pair.number >= 10
}

func splitPair(pair *Pair, parent *Pair) {
	num := float64(pair.number)
	pair.left = &Pair{nil, nil, parent, int(math.Floor(num / 2))}
	pair.right = &Pair{nil, nil, parent, int(math.Ceil(num / 2))}
}

func addToFirstLeft(startPair Pair, number int) {
	visited := make(map[Pair]bool)
	visited[startPair] = true
	current := *startPair.parent
	for !(current.parent == nil && visited[*current.left]) {
		if visited[*current.left] {
			visited[current] = true
			current = *current.parent
		} else {
			if current.left.IsNumber() {
				current.left.number += number
				return
			} else {
				current = *current.left
				for !current.right.IsNumber() {
					current = *current.right
				}
				current.right.number += number
				break
			}
		}
	}
}

func addToFirstRight(startPair Pair, number int) {
	visited := make(map[Pair]bool)
	visited[startPair] = true
	current := *startPair.parent
	for !(current.parent == nil && visited[*current.right]) {
		if visited[*current.right] {
			visited[current] = true
			current = *current.parent
		} else {
			if current.right.IsNumber() {
				current.right.number += number
				return
			} else {
				current = *current.right
				for !current.left.IsNumber() {
					current = *current.left
				}
				current.left.number += number
				break
			}
		}
	}
}

func (p Pair) Reduce() Pair {
	splitOrExplode := true
	for splitOrExplode == true {
		splitOrExplode = false
		current := p
		level := 1
		visited := make(map[Pair]bool)
		for {
			if level > 4 {
				fmt.Println("EXPLODING")
				current.Print()
				addToFirstLeft(current, current.left.number)
				addToFirstRight(current, current.right.number)
				newPair := Pair{nil, nil, current.parent, 0}
				if *current.parent.right == current {
					*current.parent.right = newPair
				} else if *current.parent.left == current {
					*current.parent.left = newPair
				} else {
					panic("")
				}
				p.Print()
				splitOrExplode = true
				break
			}
			if shouldSplit(current.left) {
				fmt.Println("Splitting")
				current.left.Print()
				splitPair(current.left, &current)
				splitOrExplode = true
				break
			}
			if shouldSplit(current.right) {
				fmt.Println("Splitting")
				current.right.Print()
				splitPair(current.right, &current)
				splitOrExplode = true
				break
			}
			if !visited[*current.left] && !current.left.IsNumber() {
				current = *current.left
				level++
			} else if !visited[*current.right] && !current.right.IsNumber() {
				current = *current.right
				level++
			} else {
				if current.parent == nil && (visited[*current.right] || current.right.IsNumber()) {
					break
				}
				visited[current] = true
				current = *current.parent
				level--
			}
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
	sum := nums[0].Reduce()
	for i := 1; i < len(nums); i++ {
		fmt.Println("Sum")
		sum.Print()
		fmt.Println("Next num")
		nums[i].Print()
		fmt.Println("Reduced")
		reduced := nums[i].Reduce()
		sum = sum.Add(reduced)
		fmt.Println("New sum")
		sum.Print()
		sum = sum.Reduce()
		fmt.Println("Reduced")
		sum.Print()
	}
	sum.Print()
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
