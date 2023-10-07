package day09p1

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"aoc/utils"
)

type direction uint

const (
	unknown direction = iota
	up
	down
	left
	right
	upRight
	upLeft
	downRight
	downLeft
)

type position struct {
	x int
	y int
}

func (p *position) String() string {
	return fmt.Sprintf("[x:%v y:%v]", p.x, p.y)
}

type knot struct {
	position
	visited map[position]int
}

func (k *knot) String() string {
	return fmt.Sprint(k.position.String())
}

func newKnot() *knot {
	return &knot{
		position: position{x: 0, y: 0},
		visited: map[position]int{
			{x: 0, y: 0}: 1,
		},
	}
}

func (k *knot) move(d direction) error {
	switch d {
	case up:
		k.y += 1
	case down:
		k.y -= 1
	case left:
		k.x -= 1
	case right:
		k.x += 1
	case upLeft:
		k.y += 1
		k.x -= 1
	case upRight:
		k.y += 1
		k.x += 1
	case downLeft:
		k.y -= 1
		k.x -= 1
	case downRight:
		k.y -= 1
		k.x += 1
	default:
		return errors.New("invalid direction")
	}
	if _, ok := k.visited[k.position]; !ok {
		k.visited[k.position] = 0
	}
	k.visited[k.position] += 1

	return nil
}

func (k *knot) distance(o *knot) (int, int) {
	xDistance := utils.AbsInt(k.x - o.x)
	yDistance := utils.AbsInt(k.y - o.y)
	return xDistance, yDistance
}

type rope struct {
	head knot
	tail knot
}

func (r *rope) String() string {
	return fmt.Sprintf("head:%v - tail:%v", r.head.String(), r.tail.String())
}

func newRope() *rope {
	return &rope{
		head: *newKnot(),
		tail: *newKnot(),
	}
}

func (r *rope) moveHead(d direction) error {
	if err := r.head.move(d); err != nil {
		return err
	}
	r.pullTail()

	return nil
}

func (r *rope) pullTail() error {
	xDistance, yDistance := r.tail.distance(&r.head)

	if xDistance < 2 && yDistance < 2 {
		return nil
	}

	var move direction
	if yDistance == 0 {
		if r.head.x > r.tail.x {
			move = right
		} else {
			move = left
		}
	}
	if xDistance == 0 {
		if r.head.y > r.tail.y {
			move = up
		} else {
			move = down
		}
	}
	if yDistance == 1 {
		if r.head.x > r.tail.x {
			if r.head.y > r.tail.y {
				move = upRight
			} else {
				move = downRight
			}
		} else {
			if r.head.y > r.tail.y {
				move = upLeft
			} else {
				move = downLeft
			}
		}
	}
	if xDistance == 1 {
		if r.head.x > r.tail.x {
			if r.head.y > r.tail.y {
				move = upRight
			} else {
				move = downRight
			}
		} else {
			if r.head.y > r.tail.y {
				move = upLeft
			} else {
				move = downLeft
			}
		}
	}

	if err := r.tail.move(move); err != nil {
		return err
	}
	return nil
}

var motionMap = map[string]direction{
	"R": right,
	"L": left,
	"U": up,
	"D": down,
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	rope := newRope()

	for _, v := range lines {
		motion := strings.Split(v, " ")
		dir, ok := motionMap[motion[0]]
		if !ok {
			fmt.Fprintf(os.Stderr, "error: unknown direction %v\n", motion[0])
			os.Exit(1)
		}
		n, err := strconv.Atoi(motion[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		for i := 0; i < n; i++ {
			rope.moveHead(dir)
		}
	}

	return len(rope.tail.visited)
}
