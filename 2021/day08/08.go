package day08

import (
	"aoc/common"
	"fmt"
	"log"
	"strings"

	"github.com/alex-ant/gomath/gaussian-elimination"
	"github.com/alex-ant/gomath/rational"
)

type Entry struct {
	signals [10]string
	outputs [4]string
}

func parseInput(input []string) []Entry {
	var entries []Entry
	for _, entry := range input {
		split := strings.Split(entry, " | ")
		var signals [10]string
		var outputs [4]string
		first, second := split[0], split[1]
		copy(signals[:], strings.Split(first, " "))
		copy(outputs[:], strings.Split(second, " "))
		newEntry := Entry{signals, outputs}
		entries = append(entries, newEntry)
	}
	return entries
}

func p1(entries []Entry) int {
	uniqueDigits := 0
	uniqueCounts := [4]int{2, 4, 3, 7}
	for _, entry := range entries {
		for _, output := range entry.outputs {
			for _, unique := range uniqueCounts {
				if len(output) == unique {
					uniqueDigits++
				}
			}
		}
	}
	return uniqueDigits
}

func nr(i int64) rational.Rational {
	return rational.New(i, 1)
}

func generateEquation(combo string) []rational.Rational {

}

func solver(entry Entry) map[int][int] {
	equations := make([][]rational.Rational, 10)
	equations[0] = []rational.Rational{nr(1), nr(2), nr(1), nr(1), nr(12)}
	equations[1] = []rational.Rational{nr(3), nr(1), nr(2), nr(2), nr(19)}
	equations[2] = []rational.Rational{nr(2), nr(5), nr(3), nr(1), nr(25)}
	equations[3] = []rational.Rational{nr(1), nr(3), nr(3), nr(2), nr(24)}

	res, gausErr := gaussian.SolveGaussian(equations, false)
	if gausErr != nil {
		log.Fatal(gausErr)
	}

	for _, v := range res {
		log.Println(v)
	}
	return 1
}

func p2(entries []Entry) int {
	nr := func(i int64) rational.Rational {
		return rational.New(i, 1)
	}

	equations := make([][]rational.Rational, 4)
	equations[0] = []rational.Rational{nr(1), nr(2), nr(1), nr(1), nr(12)}
	equations[1] = []rational.Rational{nr(3), nr(1), nr(2), nr(2), nr(19)}
	equations[2] = []rational.Rational{nr(2), nr(5), nr(3), nr(1), nr(25)}
	equations[3] = []rational.Rational{nr(1), nr(3), nr(3), nr(2), nr(24)}

	res, gausErr := gaussian.SolveGaussian(equations, false)
	if gausErr != nil {
		log.Fatal(gausErr)
	}

	for _, v := range res {
		log.Println(v)
	}
	return 1
}

func Run() {
	common.PrintDay(8)
	input := common.ReadFile("08")
	entries := parseInput(input)
	fmt.Println(p1(entries))
	fmt.Println(p2(entries))
}
