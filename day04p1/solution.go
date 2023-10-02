package day04p1

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	res := 0
	for _, v := range lines {
		pair := strings.Split(v, ",")

		elf1Lower, elf1Upper, err := bounds(pair[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		elf2Lower, elf2Upper, err := bounds(pair[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		if elf2Lower >= elf1Lower && elf2Upper <= elf1Upper {
			res++
		} else if elf1Lower >= elf2Lower && elf1Upper <= elf2Upper {
			res++
		}
	}

	return res
}

func bounds(v string) (int, int, error) {
	bounds := strings.Split(v, "-")
	lower, err := strconv.Atoi(bounds[0])
	if err != nil {
		return 0, 0, err
	}
	upper, err := strconv.Atoi(bounds[1])
	if err != nil {
		return 0, 0, err
	}
	return lower, upper, nil
}
