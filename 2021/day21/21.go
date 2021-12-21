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
		if isGameOver(p1Score, target) {
			break
		}
		move = updateMove(move)

		p2Pos = newPosition(p2Pos, move)
		p2Score += p2Pos
		rolls += 3
		if isGameOver(p2Score, target) {
			break
		}
		move = updateMove(move)
	}
	if isGameOver(p1Score, target) {
		return p2Score * rolls
	} else {
		return p1Score * rolls
	}
}

func getMoves() []int {
	var moves []int
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				moves = append(moves, i+j+k)
			}
		}
	}
	return moves
}

func getNumWon(p1Pos, p2Pos, p1Score, p2Score, target int, p1Won, p2Won int64, p1Turn bool) (int64, int64) {
	if isGameOver(p1Score, target) {
		//fmt.Println("p1 won with", p1Score)
		return p1Won + 1, p2Won
	} else if isGameOver(p2Score, target) {
		//fmt.Println("p2 won with", p2Score)
		return p1Won, p2Won + 1
	} else if p1Turn {
		//fmt.Println("p1 turn with pos", p1Pos, "and score", p1Score)
		for _, move := range getMoves() {
			splitP1Pos := newPosition(p1Pos, move)
			//fmt.Println("p1 rolled", move, "now at pos", splitP1Pos)
			splitP1Score := p1Score + splitP1Pos
			newP1Won, newP2Won := getNumWon(splitP1Pos, p2Pos, splitP1Score, p2Score, target, p1Won, p2Won, false)
			p1Won += newP1Won - p1Won
			p2Won += newP2Won - p2Won
		}
	} else {
		//fmt.Println("p2 turn with pos", p2Pos, "and score", p2Score)
		for _, move := range getMoves() {
			splitP2Pos := newPosition(p2Pos, move)
			splitP2Score := p2Score + splitP2Pos
			newP1Won, newP2Won := getNumWon(p1Pos, splitP2Pos, p1Score, splitP2Score, target, p1Won, p2Won, true)
			p1Won += newP1Won - p1Won
			p2Won += newP2Won - p2Won
		}
	}
	return p1Won, p2Won
}

func p2(p1Start int, p2Start int) int64 {
	defer common.Time()()
	const target = 21
	p1Won, p2Won := getNumWon(p1Start, p2Start, 0, 0, target, 0, 0, true)
	fmt.Println(p1Won, p2Won)
	if p1Won > p2Won {
		return p1Won
	} else {
		return p2Won
	}
}

func Run() {
	common.PrintDay(21)
	p1Start, p2Start := 4, 8
	//p1Start, p2Start := 3, 4
	fmt.Println(p1(p1Start, p2Start))
	fmt.Println(p2(p1Start, p2Start))
}

func isGameOver(score int, target int) bool {
	return score >= target
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
