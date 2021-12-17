package day16

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type Packet struct {
	version      int
	typeId       int
	literalValue int
	subPackets   []Packet
}

func (p Packet) Equal(x Packet) bool {
	fieldsEq := p.version == x.version && p.typeId == x.typeId && p.literalValue == x.literalValue
	subPacketsEqual := true

	var maxLen int
	if len(p.subPackets) > len(x.subPackets) {
		maxLen = len(p.subPackets)
	} else {
		maxLen = len(x.subPackets)
	}

	for i := 0; i < maxLen; i++ {
		if i >= len(p.subPackets) || i >= len(x.subPackets) || !p.subPackets[i].Equal(x.subPackets[i]) {
			subPacketsEqual = false
		}
	}
	return fieldsEq && subPacketsEqual
}

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
	dec, _ := strconv.ParseInt(bits, 2, 64)
	return int(dec)
}

func read(stream *strings.Reader, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		b, _ := stream.ReadByte()
		out += string(b)
	}
	return out
}

func parseLiteral(stream *strings.Reader) int {
	bitString := ""
	for {
		next := read(stream, 5)
		if string(next[0]) == "0" {
			bitString += next[1:]
			return bitsToDec(bitString)
		}
		bitString += next[1:]
	}
}

func getVersionAndTypeId(stream *strings.Reader) (int, int) {
	version := bitsToDec(read(stream, 3))
	typeId := bitsToDec(read(stream, 3))
	return version, typeId
}

func currentPos(stream *strings.Reader) int {
	return int(stream.Size()) - stream.Len()
}

func getOperatorSubpackets(stream *strings.Reader) []Packet {
	var subPackets []Packet
	I := read(stream, 1)
	switch I {
	case "0":
		length := bitsToDec(read(stream, 15))
		start := currentPos(stream)
		for currentPos(stream)-start < length {
			subPackets = append(subPackets, parsePacket(stream))
		}
	case "1":
		numPackets := bitsToDec(read(stream, 11))
		for parsed := 0; parsed < numPackets; parsed++ {
			subPackets = append(subPackets, parsePacket(stream))
		}
	}
	return subPackets
}

func parsePacket(stream *strings.Reader) Packet {
	version, typeId := getVersionAndTypeId(stream)
	var subPackets []Packet
	if typeId == 4 {
		literalValue := parseLiteral(stream)
		return Packet{version, typeId, literalValue, subPackets}
	} else {
		subPackets = getOperatorSubpackets(stream)
	}
	return Packet{version, typeId, -1, subPackets}
}

func sumVersions(packet Packet) int {
	sum := packet.version
	for _, p := range packet.subPackets {
		sum += sumVersions(p)
	}
	return sum
}

func operateOnSubPackets(subPackets []Packet, operator func(int, int) int) int {
	result := evaluate(subPackets[0])
	for i := 1; i < len(subPackets); i++ {
		result = operator(result, evaluate(subPackets[i]))
	}
	return result
}

func evaluate(packet Packet) int {
	switch packet.typeId {
	case 0:
		return operateOnSubPackets(packet.subPackets, func(x int, y int) int { return x + y })
	case 1:
		return operateOnSubPackets(packet.subPackets, func(x int, y int) int { return x * y })
	case 2:
		return operateOnSubPackets(packet.subPackets, func(x int, y int) int {
			if x < y {
				return x
			} else {
				return y
			}
		})
	case 3:
		return operateOnSubPackets(packet.subPackets, func(x int, y int) int {
			if x < y {
				return y
			} else {
				return x
			}
		})
	case 4:
		return packet.literalValue
	case 5:
		return operateOnSubPackets(packet.subPackets, func(x int, y int) int {
			if x > y {
				return 1
			} else {
				return 0
			}
		})
	case 6:
		return operateOnSubPackets(packet.subPackets, func(x int, y int) int {
			if x < y {
				return 1
			} else {
				return 0
			}
		})
	case 7:
		return operateOnSubPackets(packet.subPackets, func(x int, y int) int {
			if x == y {
				return 1
			} else {
				return 0
			}
		})
	}
	return -1
}

func p1(input string) int {
	defer common.Time()()
	bits := hexToBits(input)
	stream := strings.NewReader(bits)
	return sumVersions(parsePacket(stream))
}

func p2(input string) int {
	defer common.Time()()
	bits := hexToBits(input)
	stream := strings.NewReader(bits)
	return evaluate(parsePacket(stream))
}

func test() {
	var hex string
	var expected Packet
	var got Packet
	parsePacketFromHex := func(hex string) Packet {
		return parsePacket(strings.NewReader(hexToBits(hex)))
	}
	tester := func(hex string, expected Packet, got Packet) {
		if !got.Equal(expected) {
			fmt.Println("Got")
			fmt.Println(got)
			fmt.Println("Expected")
			fmt.Println(expected)
			panic(hex)
		}
	}

	hex = "D2FE28"
	//fmt.Println("Testing", hex, hexToBits(hex))
	expected = Packet{6, 4, 2021, []Packet{}}
	got = parsePacketFromHex(hex)
	tester(hex, expected, got)

	hex = "38006F45291200"
	//fmt.Println("Testing", hex, hexToBits(hex))
	expected = Packet{1, 6, -1, []Packet{Packet{6, 4, 10, []Packet{}}, Packet{2, 4, 20, []Packet{}}}}
	got = parsePacketFromHex(hex)
	tester(hex, expected, got)
}

func Run() {
	common.PrintDay(16)
	test()
	input := common.ReadFile("16")[0]
	fmt.Println(p1(input))
	fmt.Println(p2(input))
}
