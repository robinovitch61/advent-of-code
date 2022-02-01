package day24

import (
	"aoc/common"
	"fmt"
)

func getAllNumsDecreasing() []int {
	var nums []int
	for a := 9; a > 0; a-- {
		for b := 9; b > 0; b-- {
			for c := 9; c > 0; c-- {
				for d := 9; d > 0; d-- {
					for e := 9; e > 0; e-- {
						for f := 9; f > 0; f-- {
							for g := 9; g > 0; g-- {
								combo := []int{a, b, c, d, e, f, g}
								num := 0
								tens := 1
								for _, c := range combo {
									num += c * tens
									tens *= 10
								}
								nums = append(nums, num)
							}
						}
					}
				}
			}
		}
	}
	return nums
}

func getAllNumsIncreasing() []int {
	var nums []int
	for a := 1; a < 10; a++ {
		for b := 1; b < 10; b++ {
			for c := 1; c < 10; c++ {
				for d := 1; d < 10; d++ {
					for e := 1; e < 10; e++ {
						for f := 1; f < 10; f++ {
							for g := 1; g < 10; g++ {
								combo := []int{a, b, c, d, e, f, g}
								num := 0
								tens := 1
								for _, c := range combo {
									num += c * tens
									tens *= 10
								}
								nums = append(nums, num)
							}
						}
					}
				}
			}
		}
	}
	return nums
}

func split(num int) []int {
	var nums []int
	for num > 0 {
		next := num % 10
		nums = append([]int{next}, nums...)
		num = num / 10
	}
	return nums
}

func join(nums []int) int {
	num := 0
	tens := 1
	for i := 0; i < len(nums); i++ {
		num += nums[i] * tens
		tens *= 10
	}
	return num
}

func checkNum(num int, nums []Num) (int, bool) {
	splitNum := split(num)
	splitIdx := 0
	z := 0
	var aluInput []int

	for _, n := range nums {
		if n.isTwentySix {
			next := (z % 26) + n.comp
			if !(next > 0 && next < 10) {
				return 0, false
			}
			aluInput = append([]int{next}, aluInput...)
			z /= 26
		} else {
			next := splitNum[splitIdx]
			aluInput = append([]int{next}, aluInput...)
			z = z*26 + next + n.add
			splitIdx++
		}
	}

	return join(aluInput), true
}

type Num struct {
	isTwentySix bool
	comp, add   int
}

func Run() {
	common.PrintDay(24)
	defer common.Time()()
	nums := []Num{
		{false, 11, 8},
		{false, 12, 8},
		{false, 10, 12},
		{true, -8, 10},
		{false, 15, 2},
		{false, 15, 8},
		{true, -11, 4},
		{false, 10, 9},
		{true, -3, 10},
		{false, 15, 3},
		{true, -3, 7},
		{true, -1, 7},
		{true, -10, 2},
		{true, -16, 2},
	}
	// p1
	allNums := getAllNumsDecreasing()
	for _, num := range allNums {
		if aluInput, isValid := checkNum(num, nums); isValid {
			fmt.Println(aluInput)
			break
		}
	}
	// p2
	allNums = getAllNumsIncreasing()
	for _, num := range allNums {
		if aluInput, isValid := checkNum(num, nums); isValid {
			fmt.Println(aluInput)
			break
		}
	}
}
