package day01p2

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 45000},
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
