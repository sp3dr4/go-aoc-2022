package day11p1

import (
	"aoc/utils"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var nameRe = regexp.MustCompile(`Monkey (?P<Name>\d+):\s+`)
var itemsRe = regexp.MustCompile(`Starting items: (?P<Items>\d+(,\s\d+)*)\s+`)
var inspectRe = regexp.MustCompile(`Operation: new = old (?P<Op>[*+]) (?P<OpVal>(\d+|old))\s+`)
var testRe = regexp.MustCompile(`Test: divisible by (?P<Divisor>\d+)\s+If true: throw to monkey (?P<TrueMonkey>\d+)\s+If false: throw to monkey (?P<FalseMonkey>\d+)`)

func buildOp(operation string, operand string) (func(int) int, error) {
	switch operation {
	case "*":
		if operand == "old" {
			return func(v int) int { return v * v }, nil
		} else {
			intOperand, err := strconv.Atoi(operand)
			if err != nil {
				return nil, err
			}
			return func(v int) int { return v * intOperand }, nil
		}
	case "+":
		if operand == "old" {
			return func(v int) int { return v + v }, nil
		} else {
			intOperand, err := strconv.Atoi(operand)
			if err != nil {
				return nil, err
			}
			return func(v int) int { return v + intOperand }, nil
		}
	default:
		return nil, fmt.Errorf("invalid operation %v", operation)
	}
}

type Monkey struct {
	name      string
	items     []int
	inspect   func(old int) int
	inspected int
	test      func(v int) string
}

func (m *Monkey) String() string {
	return fmt.Sprintf("<Monkey %v: %v>", m.name, m.items)
}

func Solve(r io.Reader) any {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, r); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	const totRounds = 20

	monkeys, err := parse(buf.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	monkeysOrder := utils.SortedKeys(monkeys)

	for r := 0; r < totRounds; r++ {
		for _, k := range monkeysOrder {
			monkey := monkeys[k]
			for _, item := range monkey.items {
				postRelief := int(math.Floor(float64(monkey.inspect(item)) / 3.0))
				monkey.inspected += 1
				targetMonkeyName := monkey.test(postRelief)
				targetMonkey := monkeys[targetMonkeyName]
				targetMonkey.items = append(targetMonkey.items, postRelief)
				monkeys[targetMonkeyName] = targetMonkey
			}
			monkey.items = []int{}
			monkeys[k] = monkey
		}
	}

	top := make([]int, 2)
	for _, monkey := range monkeys {
		if monkey.inspected > top[0] {
			top[0] = monkey.inspected
			slices.Sort(top)
		}
	}

	return utils.Multiply(top...)
}

func parse(notes string) (map[string]Monkey, error) {
	monkeys := map[string]Monkey{}

	groups := strings.Split(notes, "\n\n")
	for _, g := range groups {
		matches := nameRe.FindStringSubmatch(g)
		name := matches[nameRe.SubexpIndex("Name")]

		matches = itemsRe.FindStringSubmatch(g)
		itemsStr := strings.Split(matches[itemsRe.SubexpIndex("Items")], ", ")
		items := []int{}
		for _, x := range itemsStr {
			v, err := strconv.Atoi(x)
			if err != nil {
				return nil, err
			}
			items = append(items, v)
		}

		matches = inspectRe.FindStringSubmatch(g)
		op := matches[inspectRe.SubexpIndex("Op")]
		inspectFn, err := buildOp(op, matches[inspectRe.SubexpIndex("OpVal")])
		if err != nil {
			return nil, err
		}

		matches = testRe.FindStringSubmatch(g)
		divisor, err := strconv.Atoi(matches[testRe.SubexpIndex("Divisor")])
		if err != nil {
			return nil, err
		}
		trueMonkey := matches[testRe.SubexpIndex("TrueMonkey")]
		falseMonkey := matches[testRe.SubexpIndex("FalseMonkey")]

		monkeys[name] = Monkey{
			name:      name,
			items:     items,
			inspect:   inspectFn,
			inspected: 0,
			test: func(v int) string {
				if v%divisor == 0 {
					return trueMonkey
				}
				return falseMonkey
			},
		}
	}

	return monkeys, nil
}
