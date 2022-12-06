package day5

import (
	"fmt"
	"sort"

	"github.com/chenzxcvb/advent-of-code/go/common/inputs"
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
)

func New() runner.Day {
	return &day5{}
}

type move struct {
	count int
	from  int
	to    int
}

type day5 struct {
	stacks [][]string
	moves  []move
}

func (d *day5) Open() {
	input := inputs.LinesAsString(5)

	// find blank index
	blankIndex := 0
	for i, line := range input {
		if line == "" {
			blankIndex = i
			break
		}
	}

	// set up initial state
	stacks := make([][]string, 0)
	stackLabels := input[blankIndex-1]
	for i := 0; i < len(stackLabels); i++ {
		if string(stackLabels[i]) == " " {
			continue
		}

		currentStack := make([]string, 0)
		for lineIndex := blankIndex - 2; lineIndex >= 0; lineIndex-- {
			line := input[lineIndex]
			crate := string(line[i])
			if crate == " " {
				break
			}
			currentStack = append(currentStack, crate)
		}

		stacks = append(stacks, currentStack)
	}

	d.stacks = stacks

	// process moves
	moves := make([]move, 0)
	for _, m := range input[blankIndex+1:] {
		count := 0
		from := 0
		to := 0
		_, _ = fmt.Sscanf(m, "move %d from %d to %d", &count, &from, &to)
		moves = append(moves, move{count: count, from: from - 1, to: to - 1})
	}
	d.moves = moves
}

func (d *day5) Close() {
	d.stacks = nil
	d.moves = nil
}

func (d *day5) Part1() string {

	for _, m := range d.moves {
		count := m.count
		for count > 0 {
			crate, stack := pop(d.stacks[m.from])
			d.stacks[m.from] = stack
			d.stacks[m.to] = append(d.stacks[m.to], crate)
			count--
		}
	}

	output := make([]string, 0)
	for _, st := range d.stacks {
		output = append(output, st[len(st)-1])
	}

	// fmt.Println(d.stacks)
	return fmt.Sprint(output)
}

func (d *day5) Part2() string {

	return fmt.Sprint(0)
}

func shift(s []string) (string, []string) {
	l := len(s)
	if l == 0 {
		return "", s
	}
	return s[0], s[1:]
}

func pop(s []string) (string, []string) {
	l := len(s)
	if l == 0 {
		return "", s
	}
	return s[l-1], s[:l-1]
}

func reverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
