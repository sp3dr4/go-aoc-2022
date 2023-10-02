package day03p2

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
	batch := 3
	for i := 0; i < len(lines); i += batch {
		rucksack1 := []rune(lines[i+0])
		rucksack2 := []rune(lines[i+1])
		rucksack3 := []rune(lines[i+2])
		found, err := common(rucksack1, rucksack2, rucksack3)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		res = res + (strings.IndexRune(alphabet, found) + 1)
	}

	return res
}

func common(a, b, c []rune) (rune, error) {
	aMap := make(map[rune]bool)
	for _, r := range a {
		aMap[r] = true
	}

	bMap := make(map[rune]bool)
	for _, r := range b {
		bMap[r] = true
	}

	for _, r := range c {
		if aMap[r] && bMap[r] {
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
