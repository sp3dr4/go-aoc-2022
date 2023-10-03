package day06p1

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	stream := make(chan rune)
	var wg sync.WaitGroup

	go dataStreamer(lines[0], stream)

	wg.Add(1)
	var characters []rune
	var err error
	go func() {
		defer wg.Done()
		characters, err = markerLocker(4, stream)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}()
	wg.Wait()

	return len(characters)
}

func dataStreamer(buffer string, stream chan<- rune) {
	defer close(stream)
	for _, r := range buffer {
		stream <- r
	}
}

func markerLocker(length int, stream <-chan rune) ([]rune, error) {
	received := []rune{}
	for r := range stream {
		received = append(received, r)
		receivedLen := len(received)
		if receivedLen >= length && len(dedupe[rune](received[receivedLen-length:])) == length {
			return received, nil
		}
	}
	return nil, errors.New("no marker found")
}

func dedupe[T comparable](slice []T) []T {
	dedupeMap := make(map[T]struct{})
	list := []T{}

	for _, slice := range slice {
		if _, exists := dedupeMap[slice]; !exists {
			dedupeMap[slice] = struct{}{}
			list = append(list, slice)
		}
	}

	return list
}
