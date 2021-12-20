package day18

import (
	"aoc/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Pair struct {
	parent *Pair
	number int
	left   *Pair
	right  *Pair
}

func (p Pair) ToString() string {
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

func (p Pair) IsNumber() bool {
	return p.left == nil && p.right == nil
}

func (p Pair) ShouldSplit() bool {
	return p.IsNumber() && p.number >= 10
}

func splitPair(pair *Pair) {
	num := float64(pair.number)
	fmt.Println(num)
	*pair = Pair{parent: pair.parent, number: 0}
	pair.left = &Pair{parent: pair, number: int(math.Floor(num / 2))}
	pair.right = &Pair{parent: pair, number: int(math.Ceil(num / 2))}
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
	addToFirstLeft(pair, pair.left.number)
	fmt.Println("After left")
	fmt.Println(pair.ToString())
	addToFirstRight(pair, pair.right.number)
	fmt.Println("After right")
	fmt.Println(pair.ToString())
	*pair = Pair{parent: pair.parent, number: 0}
}

func possiblyExplode(p *Pair) bool {
	current := p
	level := 1
	visited := make(map[Pair]bool)
	for {
		if level > 4 {
			fmt.Println()
			fmt.Println("Exploding")
			fmt.Println(p.ToString())
			fmt.Println(current.ToString())
			addToFirstLeft(current, current.left.number)
			fmt.Println("After left")
			fmt.Println(p.ToString())
			addToFirstRight(current, current.right.number)
			fmt.Println("After right")
			fmt.Println(p.ToString())
			*current = Pair{parent: current.parent, number: 0}
			fmt.Println("After reset")
			fmt.Println(p.ToString())
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

func possiblySplit(p *Pair) bool {
	current := p
	level := 1
	visited := make(map[Pair]bool)
	for {
		if current.left.ShouldSplit() {
			fmt.Println()
			fmt.Println("Splitting")
			fmt.Println(p.ToString())
			splitPair(current.left)
			fmt.Println("After split")
			fmt.Println(p.ToString())
			return true
		} else if current.right.ShouldSplit() {
			fmt.Println()
			fmt.Println("Splitting")
			fmt.Println(p.ToString())
			splitPair(current.right)
			fmt.Println("After split")
			fmt.Println(p.ToString())
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
	for didSplit, didExplode := true, true; didSplit || didExplode; {
		for didExplode {
			didExplode = possiblyExplode(p)
		}
		didSplit = possiblySplit(p)
		didExplode = possiblyExplode(p)
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
	right := parseToPair(rightStr, parent)
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
	newNode := &Pair{left: p1, right: p2}
	setParents(p1, newNode)
	setParents(p2, newNode)
	return newNode
}

func test() {
	var pairs []Pair
	var sum Pair

	testSum := func(input []string, expected string) {
		pairs = parseInput(input)
		sum = addAllPairs(pairs)
		if sum.ToString() != expected {
			fmt.Println("Got " + sum.ToString())
			fmt.Println("Not " + expected)
			panic("")
		}
	}

	//testExplode := func(toExplode string, expected string) {
	//	pair := parseToPair(toExplode, nil)
	//	possiblyExplode(&pair)
	//	if pair.ToString() != expected {
	//		fmt.Println("Got " + pair.ToString())
	//		fmt.Println("Not " + expected)
	//		panic("")
	//	}
	//}
	//
	//testExplode("[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]")
	//testExplode("[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]")
	//testExplode("[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]")
	//testExplode("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	//testExplode("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]")
	//
	//testSum([]string{"[1,2]", "[[3,4],5]"}, "[[1,2],[[3,4],5]]")
	//testSum([]string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"}, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	//testSum([]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}, "[[[[1,1],[2,2]],[3,3]],[4,4]]")
	//testSum([]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}, "[[[[3,0],[5,3]],[4,4]],[5,5]]")
	//testSum([]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"}, "[[[[5,0],[7,4]],[5,5]],[6,6]]")
	//testSum([]string{"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]", "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"}, "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]")
	//testSum([]string{"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]", "[2,9]"}, "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]")

	testSum([]string{"[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]", "[[[[4,2],2],6],[8,7]]"}, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
	//testSum([]string{"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]", "[7,[5,[[3,8],[1,4]]]]"}, "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]")

	//testSum([]string{
	//	"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
	//	"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
	//	"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
	//	"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
	//	"[7,[5,[[3,8],[1,4]]]]",
	//	"[[2,[2,2]],[8,[8,1]]]",
	//	"[2,9]",
	//	"[1,[[[9,3],9],[[9,0],[0,7]]]]",
	//	"[[[5,[7,4]],7],1]",
	//	"[[[[4,2],2],6],[8,7]]",
	//}, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")

	fmt.Println("Tests passed")
}

func addAllPairs(pairs []Pair) Pair {
	sum := pairs[0].Reduce()
	for i := 1; i < len(pairs); i++ {
		sum = addPairs(sum, &pairs[i])
		sum = sum.Reduce()
	}
	return *sum
}

func p1(pairs []Pair) int {
	defer common.Time()()
	sum := addAllPairs(pairs)
	fmt.Println(sum.ToString())
	return -1
}

func p2(nums []Pair) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(18)
	test()
	input := common.ReadFile("18")
	nums := parseInput(input)
	fmt.Println(p1(nums))
	//fmt.Println(p2(nums))
}
