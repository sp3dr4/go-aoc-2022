package day12p1

import (
	"fmt"
	"io"

	"aoc/utils"
)

type Point struct {
	row int
	col int
}

func (p *Point) String() string {
	return fmt.Sprintf("[%v, %v]", p.row, p.col)
}

type Heightmap [][]string

func (h *Heightmap) String() string {
	res := ""
	for _, row := range *h {
		for _, p := range row {
			res += p
		}
		res += "\n"
	}
	return res
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	heightmap, start, finish := parse(lines)

	R := len(heightmap)
	C := len(heightmap[0])

	rq, cq := []int{}, []int{}

	moveCount := 0
	nodesLeftInLayer := 1
	nodesInNextLayer := 0

	reachedEnd := false

	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}

	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, 1, -1}

	rq = append(rq, start.row)
	cq = append(cq, start.col)
	visited[start.row][start.col] = true

	for len(rq) > 0 {
		var sr, sc int
		sr, rq = rq[0], rq[1:]
		sc, cq = cq[0], cq[1:]

		if sr == finish.row && sc == finish.col {
			reachedEnd = true
			break
		}

		// findNeighbours
		toSkip := func(vr int, vc int) bool {
			if vr < 0 || vr >= R || vc < 0 || vc >= C {
				return true
			}
			if visited[vr][vc] {
				return true
			}
			if int(heightmap[vr][vc][0]) > int(heightmap[sr][sc][0])+1 {
				return true
			}
			return false
		}
		for i := 0; i < 4; i++ {
			rr := sr + dr[i]
			cc := sc + dc[i]
			if toSkip(rr, cc) {
				continue
			}
			rq = append(rq, rr)
			cq = append(cq, cc)
			visited[rr][cc] = true
			nodesInNextLayer += 1
		}
		// end

		nodesLeftInLayer -= 1
		if nodesLeftInLayer == 0 {
			nodesLeftInLayer = nodesInNextLayer
			nodesInNextLayer = 0
			moveCount += 1
		}
	}

	if reachedEnd {
		return moveCount
	}
	return -1
}

func parse(lines []string) (Heightmap, *Point, *Point) {
	var startRow, startCol, endRow, endCol int
	heightmap := make(Heightmap, len(lines))
	for i := range heightmap {
		heightmap[i] = make([]string, len(lines[0]))
	}
	for row, line := range lines {
		for col, v := range []rune(line) {
			sv := string(v)
			if sv == "S" {
				startRow = row
				startCol = col
				sv = "a"
			}
			if sv == "E" {
				endRow = row
				endCol = col
				sv = "z"
			}
			heightmap[row][col] = sv
		}
	}
	return heightmap, &Point{row: startRow, col: startCol}, &Point{row: endRow, col: endCol}
}
