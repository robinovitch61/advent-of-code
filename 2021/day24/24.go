package day24

import (
	"aoc/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	kind, target string
	valIsInt     bool
	stringVal    string
	intVal       int
}

type PuzzleInput []Instruction

type Registers map[string]int

func setVal(val string, instruction *Instruction) {
	number, err := strconv.Atoi(val)
	if err == nil {
		instruction.intVal = number
		instruction.valIsInt = true
	} else {
		instruction.stringVal = val
	}
}

func parseInput(input []string) PuzzleInput {
	var instructions []Instruction
	for _, instrString := range input {
		var instruction Instruction
		split := strings.Split(instrString, " ")
		instruction.kind = split[0]
		instruction.target = split[1]
		if len(split) == 3 {
			setVal(split[2], &instruction)
		}
		instructions = append(instructions, instruction)
	}
	return instructions
}

func getNext(stream *strings.Reader) int {
	char, _ := stream.ReadByte()
	num, _ := strconv.Atoi(string(char))
	return num
}

func applyInstruction(stream *strings.Reader, reg Registers, instr Instruction) {
	target := instr.target
	var val int
	if instr.kind != "inp" {
		if instr.valIsInt {
			val = instr.intVal
		} else {
			val = reg[instr.stringVal]
		}
	}
	switch instr.kind {
	case "inp":
		reg[target] = getNext(stream)
	case "add":
		reg[target] = reg[target] + val
	case "mul":
		reg[target] = reg[target] * val
	case "div":
		reg[target] = int(math.Floor(float64(reg[target] / val)))
	case "mod":
		reg[target] = reg[target] % val
	case "eql":
		if reg[target] == val {
			reg[target] = 1
		} else {
			reg[target] = 0
		}
	}
}

func containsZero(num int) bool {
	str := strconv.Itoa(num)
	if strings.Contains(str, "0") {
		return true
	}
	return false
}

func runProgram(in int, reg Registers, instructions []Instruction) Registers {
	reg["w"], reg["x"], reg["y"], reg["z"] = 0, 0, 0, 0
	stream := strings.NewReader(strconv.Itoa(in))
	for _, instr := range instructions {
		applyInstruction(stream, reg, instr)
	}
	return reg
}

func p1(puzzleInput PuzzleInput) int {
	defer common.Time()()
	reg := make(Registers)
	for in := int(1e15 - 1); in >= 1e14; in-- {
		if containsZero(in) {
			continue
		}
		finalReg := runProgram(in, reg, puzzleInput)
		if finalReg["z"] == 0 {
			return in
		}
	}
	return -1
}

func p2(puzzleInput PuzzleInput) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(24)
	input := common.ReadFile("24")
	puzzleInput := parseInput(input)
	fmt.Println("99999999999999 =", runProgram(99999999999999, make(Registers), puzzleInput)["z"])
	fmt.Println("11111111111111 =", runProgram(11111111111111, make(Registers), puzzleInput)["z"])
	//fmt.Println(p1(puzzleInput))
	//fmt.Println(p2(puzzleInput))
}
