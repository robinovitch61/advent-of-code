package day21

import (
	"aoc/common"
	"fmt"
)

func p1(p1Start int, p2Start int) int {
	defer common.Time()()
	const target = 1000
	p1Score, p2Score := 0, 0
	p1Pos, p2Pos := p1Start, p2Start
	move := 6
	rolls := 0
	for {
		p1Pos = newPosition(p1Pos, move)
		p1Score += p1Pos
		rolls += 3
		if p1Score >= target {
			break
		}
		move = updateMove(move)

		p2Pos = newPosition(p2Pos, move)
		p2Score += p2Pos
		rolls += 3
		if p2Score >= target {
			break
		}
		move = updateMove(move)
	}
	if p1Score >= target {
		return p2Score * rolls
	} else {
		return p1Score * rolls
	}
}

type Player struct {
	score, position int
}

type GameState struct {
	p1, p2 Player
}

type GameResult struct {
	p1Won, p2Won int
}

type Memo map[GameState]GameResult

func cacheResult(state GameState, result GameResult, memo Memo) GameResult {
	memo[state] = result
	return result
}

func getNumWon(state GameState, memo Memo, splitRollCounts map[int]int) GameResult {
	if v, exists := memo[state]; exists {
		return v
	}
	if state.p2.score >= 21 {
		return cacheResult(state, GameResult{0, 1}, memo)
	}
	p1Won, p2Won := 0, 0
	for move, count := range splitRollCounts {
		newPos := newPosition(state.p1.position, move)
		newScore := state.p1.score + newPos
		result := getNumWon(GameState{state.p2, Player{newScore, newPos}}, memo, splitRollCounts)
		p1Won += result.p2Won * count
		p2Won += result.p1Won * count
	}
	return cacheResult(state, GameResult{p1Won, p2Won}, memo)
}

func p2(p1Start int, p2Start int) int {
	defer common.Time()()
	memo := make(map[GameState]GameResult)
	result := getNumWon(GameState{Player{0, p1Start}, Player{0, p2Start}}, memo, getSplitRollCounts())
	fmt.Println(result)
	if result.p1Won > result.p2Won {
		return result.p1Won
	} else {
		return result.p2Won
	}
}

func Run() {
	common.PrintDay(21)
	//p1Start, p2Start := 4, 8
	p1Start, p2Start := 3, 4
	fmt.Println(p1(p1Start, p2Start))
	fmt.Println(p2(p1Start, p2Start))
}

func getSplitRollCounts() map[int]int {
	counts := make(map[int]int)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				sum := i + j + k
				_, exists := counts[sum]
				if exists {
					counts[sum]++
				} else {
					counts[sum] = 1
				}

			}
		}
	}
	return counts
}

func updateMove(move int) int {
	if move == 0 {
		return 9
	} else {
		return move - 1
	}
}

func newPosition(position int, move int) int {
	newPos := (position + move) % 10
	if newPos == 0 {
		return 10
	}
	return newPos
}
