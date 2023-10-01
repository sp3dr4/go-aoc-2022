package day02p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `A Y
B X
C Z`

var fullInput = `A X
A Y
A Z
B X
B Y
B Z
C X
C Y
C Z`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 12},
		{fullInput, 45},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(int)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
