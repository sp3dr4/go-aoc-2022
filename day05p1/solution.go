package day05p1

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	stacks := make(map[string][]rune)

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "[") {
			lineRunes := []rune(line)
			for jx, x := range lineRunes {
				if string(x) != "[" {
					continue
				}
				crateIx := jx + 1
				stackNum := strconv.Itoa((jx / 4) + 1)
				_, ok := stacks[stackNum]
				if !ok {
					stacks[stackNum] = []rune{}
				}
				stacks[stackNum] = append(stacks[stackNum], lineRunes[crateIx])
			}
		}

		if strings.HasPrefix(line, "move") {
			instructions := strings.Split(line, " ")
			quantity, err := strconv.Atoi(instructions[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			move(stacks, instructions[3], instructions[5], quantity)
		}
	}

	keys := make([]string, 0, len(stacks))
	for k := range stacks {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	res := ""
	for _, k := range keys {
		res += string(stacks[k][0])
	}
	return res
}

func move(stacks map[string][]rune, from string, to string, quantity int) {
	for i := 0; i < quantity; i++ {
		crate := stacks[from][0]
		stacks[from] = stacks[from][1:]
		stacks[to] = append([]rune{crate}, stacks[to]...)
	}
}
