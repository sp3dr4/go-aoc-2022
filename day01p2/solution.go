package day01p2

import (
	"fmt"
	"io"
	"strconv"

	"aoc/utils"

	"golang.org/x/exp/slices"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	lines = append(lines, "")

	acc := 0
	var top [3]int

	for _, num := range lines {
		if num == "" {
			slices.Sort(top[:])
			top[0] = max(top[0], acc)
			acc = 0
		} else {
			v, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error:", err)
			}
			acc = acc + v
		}
	}

	res := 0
	for _, value := range top {
		res += value
	}
	return res
}
