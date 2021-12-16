package day16

import (
	"aoc/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point [2]int

type PuzzleInput map[Point]int

func hexCharToBits(char string) string {
	return map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}[char]
}

func hexToBits(hex string) string {
	bits := ""
	for _, char := range strings.Split(hex, "") {
		bits += hexCharToBits(char)
	}
	return bits
}

func bitsToDec(bits string) int {
	dec := 0
	for idx := len(bits) - 1; idx >= 0; idx-- {
		if string(bits[idx]) == "1" {
			dec += int(math.Pow(2, float64(len(bits)-idx-1)))
		}
	}
	return dec
}

func parseLiteralValue(bits string, idx int) (int, int) {
	nextIdx := idx
	lastOne := false
	numberBits := ""
	for !lastOne {
		fmt.Println("\tLiteral chunk", bits[nextIdx:nextIdx+5])
		if string(bits[nextIdx]) == "0" {
			lastOne = true
		}
		numberBits += bits[nextIdx+1 : nextIdx+5]
		nextIdx += 5
	}
	number := bitsToDec(numberBits)
	return nextIdx, number
}

func tooBig(idx int, bits string) bool {
	return idx > len(bits)-1
}

func parseOperator(bits string, idx int, versionSum int) (int, int) {
	lengthTypeId, _ := strconv.Atoi(string(bits[idx]))
	fmt.Println("Found operator I", lengthTypeId)
	var length int
	if lengthTypeId == 0 {
		length = 15
		idx++
		bitsForPacketLength := bits[idx : idx+length]
		subPacketsLength := bitsToDec(bitsForPacketLength)
		fmt.Println("Found subPacketsLength from", bitsForPacketLength, "as", subPacketsLength)
		idx += length
		newIdx := idx
		for newIdx-idx < subPacketsLength { // concerning
			//fmt.Println("Parsing packet, started at idx", idx, "and now at idx", newIdx, ",", newIdx-idx, "of", subPacketsLength)
			newIdx, versionSum = parsePacket(bits, newIdx, versionSum)
			//fmt.Println("Parsed packet, started at idx", idx, "and now at idx", newIdx, ",", newIdx-idx, "of", subPacketsLength)
		}
		idx = newIdx + 1 // concerning
	} else {
		length = 11
		if idx+length > len(bits)-1 {
			fmt.Println(versionSum)
		}
		idx++
		bitsForNumPackets := bits[idx : idx+length]
		numPackets := bitsToDec(bitsForNumPackets)
		fmt.Println("Found numPackets from", bitsForNumPackets, "as", numPackets, "starting at idx", idx)
		idx += length
		startIdx := idx
		numParsed := 0
		for numParsed < numPackets {
			fmt.Println("Parsing subpackets", numParsed+1, "of", numPackets, "at idx", startIdx)
			startIdx, versionSum = parsePacket(bits, startIdx, versionSum)
			fmt.Println("Parsed subpackets", numParsed+1, "of", numPackets, "now at idx", startIdx)
			numParsed++
		}
		idx = startIdx + 1 // concerning
	}
	return idx, versionSum
}

func parsePacket(bits string, idx int, versionSum int) (int, int) {
	bitsForVersion := bits[idx : idx+3]
	version := bitsToDec(bitsForVersion)
	versionSum += version
	bitsForType := bits[idx+3 : idx+6]
	typeId := bitsToDec(bitsForType)
	idx += 6
	var literalNumber int
	if typeId == 4 {
		fmt.Println("Parsing literal version", version, "from", bitsForVersion, "and typeId", typeId, "from", bitsForType)
		idx, literalNumber = parseLiteralValue(bits, idx)
		fmt.Println("Parsed literalNumber", literalNumber)
	} else {
		fmt.Println("Parsing operator version", version, "from", bitsForVersion, "and typeId", typeId, "from", bitsForType)
		idx, versionSum = parseOperator(bits, idx, versionSum)
		fmt.Println("Parsed operator")
	}
	return idx, versionSum
}

func p1(input string) int {
	defer common.Time()()
	bits := hexToBits(input)
	fmt.Println(bits)
	fmt.Println("Max index is", len(bits)-1)
	idx := 0
	versionSum := 0
	for {
		fmt.Println("Parsing packet starting at idx", idx, "with versionSum", versionSum)
		idx, versionSum = parsePacket(bits, idx, versionSum)
		fmt.Println("Parsed packet and now at idx", idx, "of total max idx", len(bits)-1)
		if len(bits)-idx <= 6 { // concerning
			break
		}
	}
	return versionSum
}

func p2(input string) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(16)
	input := common.ReadFile("16")[0]
	fmt.Println(p1(input))
	fmt.Println(p2(input))
}
