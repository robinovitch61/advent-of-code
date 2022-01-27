package day23

import (
	"aoc/common"
	"fmt"
	"math"
	"strings"
)

type EnergyMap map[string]int

type MinEnergyMemo map[State]int

type Pos struct {
	inHallway                      bool
	hallwayPos, roomNum, roomDepth int
}

func (p Pos) Print() {
	fmt.Println("inHallway:", p.inHallway, "hallwayPos:", p.hallwayPos, "roomNum:", p.roomNum, "roomDepth:", p.roomDepth)
}

type Move struct {
	from, to Pos
}

func (m Move) Print() {
	fmt.Print("From ")
	m.from.Print()
	fmt.Print("To   ")
	m.to.Print()
}

type State struct {
	hallway [7]string
	rooms   [4][2]string
}

func (s State) Print() {
	fmt.Println("#############")
	fmt.Println("#" + strings.Join(s.hallway[:2], "") + "." + s.hallway[2] + "." + s.hallway[3] + "." + s.hallway[4] + "." + strings.Join(s.hallway[5:], "") + "#")
	fmt.Println("###" + s.rooms[0][0] + "#" + s.rooms[1][0] + "#" + s.rooms[2][0] + "#" + s.rooms[3][0] + "###")
	fmt.Println("  #" + s.rooms[0][1] + "#" + s.rooms[1][1] + "#" + s.rooms[2][1] + "#" + s.rooms[3][1] + "#  ")
	fmt.Println("  #########  ")
}

func (s State) IsOrganized() bool {
	for _, room := range s.rooms {
		first := room[0]
		if first == "." {
			return false
		}
		for _, other := range room[1:] {
			if first != other {
				return false
			}
		}
	}
	return true
}

func parseInput(input []string) State {
	hallwayString := input[1]
	var hallway [7]string
	for i, pos := range []int{1, 2, 4, 6, 8, 10, 11} {
		hallway[i] = string(hallwayString[pos])
	}
	var rooms [4][2]string
	for i, roomString := range input[2:4] {
		rooms[0][i] = string(roomString[3])
		rooms[1][i] = string(roomString[5])
		rooms[2][i] = string(roomString[7])
		rooms[3][i] = string(roomString[9])
	}
	return State{hallway, rooms}
}

func getEnergyMap() EnergyMap {
	return EnergyMap{
		"A": 1,
		"B": 10,
		"C": 100,
		"D": 1000,
	}
}

func getStepsFromHallwayToRoomEntrance(hallwayPos, roomNum int) int {
	return map[int]map[int]int{
		0: {
			0: 2,
			1: 1,
			2: 1,
			3: 3,
			4: 5,
			5: 7,
			6: 8,
		},
		1: {
			0: 4,
			1: 3,
			2: 1,
			3: 1,
			4: 3,
			5: 5,
			6: 6,
		},
		2: {
			0: 6,
			1: 5,
			2: 3,
			3: 1,
			4: 1,
			5: 3,
			6: 4,
		},
		3: {
			0: 8,
			1: 7,
			2: 5,
			3: 3,
			4: 1,
			5: 1,
			6: 2,
		},
	}[roomNum][hallwayPos]
}

func moveState(state State, move Move) State {
	newState := State{state.hallway, state.rooms}
	var movedAmphipod string

	if move.from.inHallway {
		movedAmphipod = state.hallway[move.from.hallwayPos]
		newState.hallway[move.from.hallwayPos] = "."
	} else {
		movedAmphipod = state.rooms[move.from.roomNum][move.from.roomDepth]
		newState.rooms[move.from.roomNum][move.from.roomDepth] = "."
	}

	if move.to.inHallway {
		newState.hallway[move.to.hallwayPos] = movedAmphipod
	} else {
		newState.rooms[move.to.roomNum][move.to.roomDepth] = movedAmphipod
	}

	return newState
}

func numSteps(state State, move Move) int {
	if move.from.inHallway && move.to.inHallway {
		return int(math.Abs(float64(move.from.hallwayPos - move.to.hallwayPos)))
	} else if !move.from.inHallway && !move.to.inHallway {
		if move.from.roomNum == move.to.roomNum {
			return int(math.Abs(float64(move.from.roomDepth - move.to.roomDepth)))
		} else {
			return move.from.roomDepth + 1 + (2 * int(math.Abs(float64(move.from.roomNum-move.to.roomNum)))) + move.to.roomDepth + 1
		}
	} else if move.from.inHallway && !move.to.inHallway {
		return getStepsFromHallwayToRoomEntrance(move.from.hallwayPos, move.to.roomNum) + move.to.roomDepth + 1
	} else {
		return getStepsFromHallwayToRoomEntrance(move.to.hallwayPos, move.from.roomNum) + move.from.roomDepth + 1
	}
}

func energyForMove(state State, move Move, energyMap EnergyMap) int {
	var movedAmphipod string
	if move.from.inHallway {
		movedAmphipod = state.hallway[move.from.hallwayPos]
	} else {
		movedAmphipod = state.rooms[move.from.roomNum][move.from.roomDepth]
	}
	energyPerStep := energyMap[movedAmphipod]
	return energyPerStep * numSteps(state, move)
}

func getPossibleMoves(state State) []Move {
	var moves []Move

	// hallway to hallway (not possible)
	//for start := 0; start < 7; start++ {
	//	if string(state.hallway[start]) == "." {
	//		continue
	//	}
	//	for end := 0; end < 7; end++ {
	//		if start == end || string(state.hallway[end]) != "." {
	//			continue
	//		}
	//		minBetween := int(math.Min(float64(start), float64(end))) + 1
	//		maxBetween := int(math.Max(float64(start), float64(end))) - 1
	//		blocked := false
	//		for between := minBetween; between <= maxBetween; between++ {
	//			if string(state.hallway[between]) != "." {
	//				blocked = true
	//				break
	//			}
	//		}
	//		if !blocked {
	//			from := Pos{true, start, 0, 0}
	//			to := Pos{true, end, 0, 0}
	//			moves = append(moves, Move{from, to})
	//		}
	//	}
	//}

	// hallway to room
	for start := 0; start < 7; start++ {
		amphipod := state.hallway[start]
		if amphipod == "." {
			continue
		}

		destinationRoom := map[string]int{
			"A": 0,
			"B": 1,
			"C": 2,
			"D": 3,
		}[amphipod]

		// check the room can be entered
		canEnterRoom := true
		for i, roomSpot := range state.rooms[destinationRoom] {
			if i == 0 && roomSpot != "." {
				canEnterRoom = false
				break
			} else if !(roomSpot == "." || roomSpot == amphipod) {
				canEnterRoom = false
				break
			}
		}
		if !canEnterRoom {
			continue
		}

		// check the way is clear to the room
		// #############
		// #01.2.3.4.56#
		//    0 1 2 3
		// ###A#C#B#B###
		//   #D#D#A#C#
		//   #########
		canGetToRoom := true
		if start >= destinationRoom+2 {
			// start is to the right of destinationRoom
			for between := start - 1; between > destinationRoom+1; between-- {
				if state.hallway[between] != "." {
					canGetToRoom = false
					break
				}
			}
		} else {
			// start is to the left of destinationRoom
			for between := start + 1; between <= destinationRoom+1; between++ {
				if state.hallway[between] != "." {
					canGetToRoom = false
					break
				}
			}
		}
		if !canGetToRoom {
			continue
		}

		// add move into bottom of room
		for depth := len(state.rooms[destinationRoom]) - 1; depth >= 0; depth-- {
			roomSpot := state.rooms[destinationRoom][depth]
			if roomSpot == amphipod {
				continue
			} else {
				var from, to Pos
				from.inHallway = true
				from.hallwayPos = start
				to.inHallway = false
				to.roomNum = destinationRoom
				to.roomDepth = depth
				moves = append(moves, Move{from, to})
				break
			}
		}
	}

	// room to hallway
	for roomNum := 0; roomNum < 4; roomNum++ {
		for depth := 0; depth < len(state.rooms[roomNum]); depth++ {
			if amphipod := state.rooms[roomNum][depth]; amphipod != "." {
				destinationRoom := map[string]int{
					"A": 0,
					"B": 1,
					"C": 2,
					"D": 3,
				}[amphipod]

				// don't move out of the room if you and everything behind you is where it's supposed to be
				if roomNum == destinationRoom {
					behindSorted := true
					for d := depth + 1; d < len(state.rooms[roomNum]); d++ {
						if state.rooms[roomNum][d] != amphipod {
							behindSorted = false
							break
						}
					}
					if behindSorted {
						break
					}
				}

				// move left through hallway
				for hallwayPos := roomNum + 1; hallwayPos >= 0; hallwayPos-- {
					if state.hallway[hallwayPos] != "." {
						break
					} else {
						var from, to Pos
						from.inHallway = false
						from.roomNum = roomNum
						from.roomDepth = depth
						to.inHallway = true
						to.hallwayPos = hallwayPos
						moves = append(moves, Move{from, to})
					}
				}

				// move right through hallway
				for hallwayPos := roomNum + 2; hallwayPos < 7; hallwayPos++ {
					if state.hallway[hallwayPos] != "." {
						break
					} else {
						var from, to Pos
						from.inHallway = false
						from.roomNum = roomNum
						from.roomDepth = depth
						to.inHallway = true
						to.hallwayPos = hallwayPos
						moves = append(moves, Move{from, to})
					}
				}
				break
			}
		}
	}

	// room to room
	for startRoomNum := 0; startRoomNum < 4; startRoomNum++ {
		for depth := 0; depth < len(state.rooms[startRoomNum]); depth++ {
			if amphipod := state.rooms[startRoomNum][depth]; amphipod != "." {
				destinationRoom := map[string]int{
					"A": 0,
					"B": 1,
					"C": 2,
					"D": 3,
				}[amphipod]

				if startRoomNum == destinationRoom {
					break
				}

				// check the room can be entered
				canEnterRoom := true
				for i, roomSpot := range state.rooms[destinationRoom] {
					if i == 0 && roomSpot != "." {
						canEnterRoom = false
						break
					} else if !(roomSpot == "." || roomSpot == amphipod) {
						canEnterRoom = false
						break
					}
				}
				if !canEnterRoom {
					continue
				}

				// check the way is clear to the room
				// #############
				// #01.2.3.4.56#
				//    0 1 2 3
				// ###A#C#B#B###
				//   #D#D#A#C#
				//   #########
				canGetToRoom := true
				if startRoomNum > destinationRoom {
					// startRoomNum is to the right of destinationRoom
					for between := startRoomNum + 1; between >= destinationRoom+2; between-- {
						if state.hallway[between] != "." {
							canGetToRoom = false
							break
						}
					}
				} else {
					// startRoomNum is to the left of destinationRoom
					for between := startRoomNum + 2; between <= destinationRoom+1; between++ {
						if state.hallway[between] != "." {
							canGetToRoom = false
							break
						}
					}
				}
				if !canGetToRoom {
					continue
				}

				// add move into bottom of room
				for endDepth := len(state.rooms[destinationRoom]) - 1; endDepth >= 0; endDepth-- {
					roomSpot := state.rooms[destinationRoom][endDepth]
					if roomSpot == amphipod {
						continue
					} else {
						var from, to Pos
						from.inHallway = false
						from.roomNum = startRoomNum
						from.roomDepth = depth
						to.inHallway = false
						to.roomNum = destinationRoom
						to.roomDepth = endDepth
						moves = append(moves, Move{from, to})
						break
					}
				}
			}
		}
	}

	return moves
}

func getMinEnergy(state State, energyMap EnergyMap, memo MinEnergyMemo, prevStates []State, prevMoves []Move, prevPossible [][]Move) int {
	if v, exists := memo[state]; exists {
		return v
	}

	if state.IsOrganized() {
		for i := 0; i < len(prevStates); i++ {
			prevMoves[i].Print()
			prevStates[i].Print()
			for _, p := range prevPossible[i] {
				p.Print()
				//fmt.Println(numSteps(prevStates[i]))
				fmt.Println()
			}
			fmt.Println()
		}
		state.Print()
		memo[state] = 0
		return 0
	}

	possibleMoves := getPossibleMoves(state)
	minEnergy := int(1e9)

	for _, move := range possibleMoves {
		newState := moveState(state, move)
		energyUsedInMove := energyForMove(state, move, energyMap)
		minEnergyFromNewState := getMinEnergy(newState, energyMap, memo, append(prevStates, newState), append(prevMoves, move), append(prevPossible, possibleMoves))
		if minEnergyFromNewState+energyUsedInMove < minEnergy {
			minEnergy = minEnergyFromNewState + energyUsedInMove
		}
	}

	memo[state] = minEnergy
	return minEnergy
}

func p1(initialState State) int {
	defer common.Time()()
	energyMap := getEnergyMap()
	minEnergyMemo := make(MinEnergyMemo)
	prevStates := []State{initialState}
	prevMoves := []Move{{}}
	prevPossible := [][]Move{{}}
	return getMinEnergy(initialState, energyMap, minEnergyMemo, prevStates, prevMoves, prevPossible)
}

func p2(initialState State) int {
	defer common.Time()()
	return -1
}

func Run() {
	common.PrintDay(23)
	input := common.ReadFile("23")
	initialState := parseInput(input)
	initialState.Print()
	fmt.Println(initialState.IsOrganized())
	fmt.Println(p1(initialState))
	//fmt.Println(p2(initialState))
}
