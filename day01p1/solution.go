package day01p1

import (
	"fmt"
	"io"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	lines = append(lines, "")

	res := 0
	acc := 0

	for _, num := range lines {
		if num == "" {
			res = max(res, acc)
			acc = 0
		} else {
			v, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error:", err)
			}
			acc = acc + v
		}
	}

	return res
}
