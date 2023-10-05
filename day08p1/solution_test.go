package day08p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput1 = `30373
25512
65332
33549
35390`

var testInput2 = `4444
3223
1111
1111
`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput1, 21},
		{testInput2, 14},
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
