package day09p2

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
	name string
	position
	visited map[position]int
}

func (k *knot) String() string {
	return fmt.Sprint(k.position.String())
}

func newKnot(name string) *knot {
	return &knot{
		name:     name,
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
	knots []knot
}

func (r *rope) String() string {
	knotsStr := []string{}
	for _, k := range r.knots {
		knotsStr = append(knotsStr, fmt.Sprintf("%v:%v", k.name, k.String()))
	}
	return strings.Join(knotsStr, ", ")
}

func newRope(n int) *rope {
	knots := []knot{*newKnot("H")}
	for i := 1; i < n; i++ {
		knots = append(knots, *newKnot(fmt.Sprint(i)))
	}
	return &rope{
		knots: knots,
	}
}

func (r *rope) moveHead(d direction, n int) error {
	for i := 0; i < n; i++ {
		if err := r.knots[0].move(d); err != nil {
			return err
		}
		r.pullTail()
	}
	return nil
}

func (r *rope) pullTail() error {
	for ix := 1; ix < len(r.knots); ix++ {
		if err := r.pullTailKnot(ix); err != nil {
			return err
		}
	}
	return nil
}

func (r *rope) pullTailKnot(ix int) error {
	this, front := &r.knots[ix], &r.knots[ix-1]
	xDistance, yDistance := this.distance(front)
	if xDistance < 2 && yDistance < 2 {
		return nil
	}

	var move direction

	if yDistance == 0 {
		if front.x > this.x {
			move = right
		} else {
			move = left
		}
	}
	if xDistance == 0 {
		if front.y > this.y {
			move = up
		} else {
			move = down
		}
	}
	if yDistance == 1 {
		if front.x > this.x {
			if front.y > this.y {
				move = upRight
			} else {
				move = downRight
			}
		} else {
			if front.y > this.y {
				move = upLeft
			} else {
				move = downLeft
			}
		}
	}
	if xDistance == 1 {
		if front.x > this.x {
			if front.y > this.y {
				move = upRight
			} else {
				move = downRight
			}
		} else {
			if front.y > this.y {
				move = upLeft
			} else {
				move = downLeft
			}
		}
	}
	if xDistance == 2 && yDistance == 2 {
		if front.x > this.x {
			if front.y > this.y {
				move = upRight
			} else {
				move = downRight
			}
		} else {
			if front.y > this.y {
				move = upLeft
			} else {
				move = downLeft
			}
		}
	}
	if err := this.move(move); err != nil {
		return err
	}
	return nil
}

func (r *rope) render(topleftX int, topleftY int, bottomrightX int, bottomrightY int) {
	for y := topleftY; y >= bottomrightY; y-- {
		for x := topleftX; x <= bottomrightX; x++ {
			c := "."
			for _, k := range r.knots {
				if k.x == x && k.y == y {
					c = k.name
					break
				}
			}
			fmt.Printf("%v", c)
		}
		fmt.Println()
	}
	fmt.Println()
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	debug := false

	rope := newRope(10)
	if debug {
		fmt.Printf("== Initial State ==\n\n")
		rope.render(-11, 15, 14, -5)
	}

	for _, v := range lines {
		motion := strings.Split(v, " ")
		dir, ok := map[string]direction{
			"R": right,
			"L": left,
			"U": up,
			"D": down,
		}[motion[0]]
		if !ok {
			fmt.Fprintf(os.Stderr, "error: unknown direction %v\n", motion[0])
			os.Exit(1)
		}
		n, err := strconv.Atoi(motion[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		rope.moveHead(dir, n)
		if debug {
			fmt.Printf("== %v %v ==\n\n", motion[0], motion[1])
			rope.render(-11, 15, 14, -5)
		}
	}

	return len(rope.knots[len(rope.knots)-1].visited)
}
