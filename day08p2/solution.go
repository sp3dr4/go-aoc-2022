package day08p2

import (
	"io"
	"slices"
	"strings"

	"aoc/utils"
)

func directionScore(trees []string, height string) int {
	i := slices.IndexFunc(trees, func(t string) bool {
		return t >= height
	})
	if i == -1 {
		return len(trees)
	}
	return i + 1
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	score := 0
	for i, row := range lines {
		if row == "" {
			continue
		}

		for j := 0; j < len(row); j++ {
			tree := string(row[j])

			eastTrees := strings.Split(row[j+1:], "")
			eastScore := directionScore(eastTrees, tree)

			westTrees := strings.Split(row[:j], "")
			slices.Reverse(westTrees)
			westScore := directionScore(westTrees, tree)

			northTrees := []string{}
			for u := i - 1; u >= 0; u-- {
				northTrees = append(northTrees, string(lines[u][j]))
			}
			northScore := directionScore(northTrees, tree)

			southTrees := []string{}
			for u := i + 1; u < len(lines); u++ {
				southTrees = append(southTrees, string(lines[u][j]))
			}
			southScore := directionScore(southTrees, tree)

			treeScore := eastScore * westScore * northScore * southScore
			score = max(score, treeScore)
		}
	}

	return score
}
