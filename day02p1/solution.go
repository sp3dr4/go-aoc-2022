package day02p1

import (
	"io"
	"strings"

	"aoc/utils"
)

const (
	scissors int = iota
	paper
	rock
)

const (
	draw = 3
	win  = 6
)

var shapeScore = map[int]int{
	scissors: 3,
	paper:    2,
	rock:     1,
}

var shapeRule = map[int]int{
	scissors: paper,
	paper:    rock,
	rock:     scissors,
}

var opponentMap = map[string]int{
	"C": scissors,
	"B": paper,
	"A": rock,
}

var playerMap = map[string]int{
	"Z": scissors,
	"Y": paper,
	"X": rock,
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	res := 0

	for _, v := range lines {
		round := strings.Split(v, " ")
		opponent := opponentMap[round[0]]
		me := playerMap[round[1]]

		if opponent == me {
			res = res + draw
		} else if shapeRule[me] == opponent {
			res = res + win
		}
		res = res + shapeScore[me]
	}

	return res
}
