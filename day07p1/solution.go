package day07p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	fs := parse(lines)

	res := 0
	for _, size := range fs {
		if size <= 100000 {
			res += size
		}
	}
	return res
}

func parse(lines []string) map[string]int {
	fs := map[string]int{}
	cwd := []string{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		args := strings.Split(line, " ")
		switch strings.Join(args[:2], "") {
		case "$cd":
			if args[2] == ".." {
				cwd = cwd[:len(cwd)-1]
			} else {
				cwd = append(cwd, args[2])
				cwdStr := strings.Join(cwd, " ")
				if _, ok := fs[cwdStr]; !ok {
					fs[cwdStr] = 0
				}
			}
		default:
			if size, err := strconv.Atoi(args[0]); err == nil {
				for dx := 0; dx < len(cwd); dx++ {
					cwdStr := strings.Join(cwd[:dx+1], " ")
					fs[cwdStr] += size
				}
			}
		}
	}
	return fs
}
