package day09p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 13},
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
