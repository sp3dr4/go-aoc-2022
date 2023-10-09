package day10p2

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"aoc/utils"
)

func render(cycle int, x int) string {
	sprite := []int{x - 1, x, x + 1}
	if slices.Contains(sprite, cycle%40) {
		return "#"
	} else {
		return "."
	}
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return v
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	width := 40

	res := ""
	x, c := 1, 0

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		switch instruction[0] {
		case "noop":
			res += render(c, x)
			c++
		case "addx":
			v := atoi(instruction[1])
			res += render(c, x)
			c++
			res += render(c, x)
			c++
			x += v
		default:
			fmt.Fprintf(os.Stderr, "error: unknown instruction %v\n", line)
			os.Exit(1)
		}
	}

	crt := strings.Join(utils.Chunks(res, width), "\n")
	return crt
}
