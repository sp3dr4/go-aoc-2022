package day03p1

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"aoc/utils"
)

var alphabet string = buildAlphabet()

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	res := 0
	for _, v := range lines {
		rucksack := []rune(v)
		rucksackSize := len(rucksack)
		if rucksackSize%2 != 0 {
			err := "rucksack compartments do not have same size"
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		found, err := common(rucksack[:rucksackSize/2], rucksack[rucksackSize/2:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		res = res + (strings.IndexRune(alphabet, found) + 1)
	}

	return res
}

func common(a, b []rune) (rune, error) {
	aMap := make(map[rune]bool)
	for _, r := range a {
		aMap[r] = true
	}

	for _, r := range b {
		if aMap[r] {
			return r, nil
		}
	}
	return 0, errors.New("no common element found")
}

func buildAlphabet() string {
	const lowercaseStart = rune('a')
	const uppercaseStart = rune('A')

	alphabet := ""
	for ch := lowercaseStart; ch < lowercaseStart+26; ch++ {
		alphabet = alphabet + string(ch)
	}
	for ch := uppercaseStart; ch < uppercaseStart+26; ch++ {
		alphabet = alphabet + string(ch)
	}
	return alphabet
}
