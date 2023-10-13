package day12p1

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"sync"

	"aoc/utils"
)

type Point struct {
	row int
	col int
}

func (p *Point) String() string {
	return fmt.Sprintf("[%v, %v]", p.row, p.col)
}

type Path struct {
	points   []Point
	finished bool
}

func (p *Path) Contains(x *Point) bool {
	return slices.Contains(p.points, *x)
}

func (p *Path) Render(width int, height int) string {
	area := make([][]string, height)
	for i := range area {
		area[i] = make([]string, width)
		for j := range area[i] {
			area[i][j] = "."
		}
	}
	for i, point := range p.points {
		if i == len(p.points)-1 {
			area[point.row][point.col] = "F"
			break
		}
		if p.points[i+1].row == point.row+1 {
			area[point.row][point.col] = "v"
		} else if p.points[i+1].row == point.row-1 {
			area[point.row][point.col] = "^"
		} else if p.points[i+1].col == point.col+1 {
			area[point.row][point.col] = ">"
		} else if p.points[i+1].col == point.col-1 {
			area[point.row][point.col] = "<"
		}
	}
	areaStr := []string{}
	for _, r := range area {
		areaStr = append(areaStr, strings.Join(r, ""))
	}
	return strings.Join(areaStr, "\n")
}

type Heightmap [][]rune

func (h *Heightmap) String() string {
	res := ""
	for _, row := range *h {
		for _, p := range row {
			res += string(p)
		}
		res += "\n"
	}
	return res
}

func (h *Heightmap) Value(p *Point) rune {
	return (*h)[p.row][p.col]
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	heightmap, start, finish := parse(lines)

	paths := walk(&heightmap, Path{points: []Point{*start}, finished: false}, finish)
	if len(paths) == 0 {
		fmt.Fprintf(os.Stderr, "error: walk found no paths\n")
		os.Exit(1)
	}

	finished := []Path{}
	for _, p := range paths {
		if p.finished {
			finished = append(finished, p)
		}
	}
	if len(finished) == 0 {
		fmt.Fprintf(os.Stderr, "error: no paths have finished\n")
		os.Exit(1)
	}

	shortest := slices.MinFunc(finished, func(a, b Path) int {
		return cmp.Compare(len(a.points), len(b.points))
	})
	return len(shortest.points) - 1
}

func walk(heightmap *Heightmap, path Path, finish *Point) []Path {
	from := path.points[len(path.points)-1]

	if from == *finish {
		final := make([]Point, len(path.points))
		copy(final, path.points)
		return []Path{{points: final, finished: true}}
	}

	isValid := func(p *Point) bool {
		inBounds := p.row >= 0 && p.row < len(*heightmap) && p.col >= 0 && p.col < len((*heightmap)[0])
		return inBounds && (!path.Contains(p)) && heightmap.Value(p) <= (heightmap.Value(&from)+1)
	}

	forks := []Point{
		{row: from.row - 1, col: from.col},
		{row: from.row + 1, col: from.col},
		{row: from.row, col: from.col - 1},
		{row: from.row, col: from.col + 1},
	}

	wg := new(sync.WaitGroup)
	paths := []Path{}

	for _, fp := range forks {
		if isValid(&fp) {
			fp := fp
			wg.Add(1)
			go func() {
				defer wg.Done()
				pathCopy := append([]Point(nil), path.points...)
				forkPath := Path{points: append(pathCopy, fp), finished: false}
				subPaths := walk(heightmap, forkPath, finish)
				paths = append(paths, subPaths...)
			}()
		}
	}
	wg.Wait()
	return paths
}

func parse(lines []string) (Heightmap, *Point, *Point) {
	var startRow, startCol, endRow, endCol int
	heightmap := make([][]rune, len(lines))
	for i := range heightmap {
		heightmap[i] = make([]rune, len(lines[0]))
	}
	for row, line := range lines {
		for col, v := range []rune(line) {
			if string(v) == "S" {
				startRow = row
				startCol = col
				v = []rune("a")[0]
			}
			if string(v) == "E" {
				endRow = row
				endCol = col
				v = []rune("z")[0]
			}
			heightmap[row][col] = v
		}
	}
	return heightmap, &Point{row: startRow, col: startCol}, &Point{row: endRow, col: endCol}
}
