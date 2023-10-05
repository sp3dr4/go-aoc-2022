package day08p1

import (
	"io"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	visible := 0
	for i, row := range lines {
		if row == "" {
			continue
		}

		if i == 0 || i == (len(lines)-1) {
			visible += len(row)
			continue
		}

		visible += 2
		for j := 1; j < len(row)-1; j++ {
			isLower := func(t string) bool {
				return t < string(row[j])
			}
			// west
			if utils.All(strings.Split(row[:j], ""), isLower) {
				visible++
				continue
			}
			// east
			if utils.All(strings.Split(row[j+1:], ""), isLower) {
				visible++
				continue
			}
			// north
			northTrees := []string{}
			for u := i - 1; u >= 0; u-- {
				northTrees = append(northTrees, string(lines[u][j]))
			}
			if utils.All(northTrees, isLower) {
				visible++
				continue
			}
			// south
			southTrees := []string{}
			for u := i + 1; u < len(lines); u++ {
				southTrees = append(southTrees, string(lines[u][j]))
			}
			if utils.All(southTrees, isLower) {
				visible++
			}
		}
	}

	return visible
}
