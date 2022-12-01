package day1

import (
	"fmt"

	"github.com/chenzxcvb/advent-of-code/go/common/inputs"
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
)

type day1 struct {
	given [][]int
}

func New() runner.Day {
	return &day1{}
}
func (d *day1) Open() {
	d.given = inputs.LinesAsIntChunk(1)
}

func (d *day1) Close() {
	d.given = nil
}

func (d *day1) Part1() string {
	result := 0
	for _, r := range d.given {
		total := sum(r)
		result = max(result, total)
	}
	return fmt.Sprint(result)
}

func (d *day1) Part2() string {
	first := 0
	second := 0
	third := 0

	for _, r := range d.given {
		total := sum(r)
		if total > first {
			third = second
			second = first
			first = total
		} else if total > second {
			third = second
			second = total
		} else if total > third {
			third = total
		}
	}

	return fmt.Sprint(sum([]int{first, second, third}))
}

func sum(arr []int) int {
	total := 0
	for _, c := range arr {
		total += c
	}
	return total
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
