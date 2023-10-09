package day10p1

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	res := 0
	x, c := 1, 1
	interesting := []int{20, 60, 100, 140, 180, 220}
	checkSignal := func(cycle int) int {
		if slices.Contains(interesting, cycle) {
			return cycle * x
		}
		return 0
	}

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		switch instruction[0] {
		case "noop":
			c++
		case "addx":
			v, err := strconv.Atoi(instruction[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			c++
			res += checkSignal(c)
			c++
			x += v
		default:
			fmt.Fprintf(os.Stderr, "error: unknown instruction %v\n", line)
			os.Exit(1)
		}
		res += checkSignal(c)
	}

	return res
}
