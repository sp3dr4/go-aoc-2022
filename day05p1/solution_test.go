package day05p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer string
	}{
		{testInput, "CMZ"},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(string)

		if result != test.answer {
			t.Errorf("Expected %s, got %s", test.answer, result)
		}
	}
}
