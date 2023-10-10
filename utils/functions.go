package utils

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"sort"
)

func All[T any](items []T, fn func(T) bool) bool {
	for _, v := range items {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}

func Multiply(args ...int) int {
	acc := 1
	for _, n := range args {
		acc *= n
	}

	return acc
}

func AbsInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// Greatest Common Denominator
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Least Common Multiple
func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

// Kernighan's Bit Counting Algorithm
func CountBits(n uint64) int64 {
	var count int64 = 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}

	return count
}

// Check if error is not nil and panic with message if it is.
func Check(e error, format string, a ...any) {
	if e != nil {
		message := fmt.Sprintf(format, a...)
		panic(fmt.Errorf("%s: %s", message, e))
	}
}

// Read all lines from reader. Panic if there is an issue
func ReadLines(r io.Reader) []string {
	result := []string{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	err := scanner.Err()
	Check(err, "error reading lines")

	return result
}
