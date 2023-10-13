package day12p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput1 = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput1, 31},
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
