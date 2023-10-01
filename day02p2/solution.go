package day02p2

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

var outcomeMap = map[string]map[int]int{
	"Z": {
		rock:     paper,
		paper:    scissors,
		scissors: rock,
	},
	"Y": {
		rock:     rock,
		paper:    paper,
		scissors: scissors,
	},
	"X": {
		rock:     scissors,
		paper:    rock,
		scissors: paper,
	},
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	res := 0

	for _, v := range lines {
		round := strings.Split(v, " ")
		opponent := opponentMap[round[0]]
		me := outcomeMap[round[1]][opponent]

		if opponent == me {
			res = res + draw
		} else if shapeRule[me] == opponent {
			res = res + win
		}
		res = res + shapeScore[me]
	}

	return res
}
