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

func (p *Pair) ToString() string {
	if p.IsNumber() {
		return strconv.Itoa(p.number)
	} else {
		res := ""
		res += "["
		res += p.left.ToString()
		res += ","
		res += p.right.ToString()
		res += "]"
		return res
	}
}

func (p *Pair) IsNumber() bool {
	return p.left == nil && p.right == nil
}

func shouldSplit(pair *Pair) bool {
	return pair.IsNumber() && pair.number >= 10
}

func splitPair(pair *Pair) {
	num := float64(pair.number)
	*pair = Pair{nil, nil, pair.parent, 0}
	pair.left = &Pair{nil, nil, pair, int(math.Floor(num / 2))}
	pair.right = &Pair{nil, nil, pair, int(math.Ceil(num / 2))}
}

func addToFirstLeft(startPair *Pair, number int) {
	visited := make(map[Pair]bool)
	visited[*startPair] = true
	current := startPair.parent
	for !(current.parent == nil && visited[*current.left]) {
		if visited[*current.left] {
			visited[*current] = true
			current = current.parent
		} else {
			if current.left.IsNumber() {
				current.left.number += number
				return
			} else {
				current = current.left
				for !current.right.IsNumber() {
					current = current.right
				}
				current.right.number += number
				return
			}
		}
	}
}

func addToFirstRight(startPair *Pair, number int) {
	visited := make(map[Pair]bool)
	visited[*startPair] = true
	current := startPair.parent
	for !(current.parent == nil && visited[*current.right]) {
		if visited[*current.right] {
			visited[*current] = true
			current = current.parent
		} else {
			if current.right.IsNumber() {
				current.right.number += number
				return
			} else {
				current = current.right
				for !current.left.IsNumber() {
					current = current.left
				}
				current.left.number += number
				return
			}
		}
		//fmt.Println("\t", current.ToString())
		//fmt.Println("\t", current.parent)
	}
}

func explode(pair *Pair) {
	//fmt.Println("Exploding")
	//fmt.Println(pair.ToString())
	addToFirstLeft(pair, pair.left.number)
	addToFirstRight(pair, pair.right.number)
	*pair = Pair{nil, nil, pair.parent, 0}
}

func traverse(p *Pair) bool {
	current := p
	level := 1
	visited := make(map[Pair]bool)
	for {
		if level > 4 {
			//fmt.Println()
			//fmt.Println(p.ToString())
			explode(current)
			//fmt.Println(p.ToString())
			return true
		}
		if shouldSplit(current.left) {
			splitPair(current.left)
			return true
		}
		if shouldSplit(current.right) {
			splitPair(current.right)
			return true
		}
		if !visited[*current.left] && !current.left.IsNumber() {
			current = current.left
			level++
		} else if !visited[*current.right] && !current.right.IsNumber() {
			current = current.right
			level++
		} else {
			if current.parent == nil && (visited[*current.right] || current.right.IsNumber()) {
				return false
			}
			visited[*current] = true
			current = current.parent
			level--
		}
	}
}

func (p *Pair) Reduce() *Pair {
	for {
		didSplitOrExplode := traverse(p)
		if !didSplitOrExplode {
			break
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

func setParents(toSet *Pair, parent *Pair) {
	toSet.parent = parent
	if toSet.left != nil {
		setParents(toSet.left, toSet)
	}
	if toSet.right != nil {
		setParents(toSet.right, toSet)
	}
}

func addPairs(p1 *Pair, p2 *Pair) *Pair {
	newNode := &Pair{p1, p2, nil, 0}
	setParents(p1, newNode)
	setParents(p2, newNode)
	return newNode
}

func p1(nums []Pair) int {
	defer common.Time()()
	sum := nums[0].Reduce()
	for i := 1; i < len(nums); i++ {
		sum = addPairs(sum, &nums[i])
		//fmt.Println(sum.ToString())
		sum = sum.Reduce()
	}
	fmt.Println(sum.ToString())
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
