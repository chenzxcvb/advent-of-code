package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chenzxcvb/advent-of-code/go/common/inputs"
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
)

type section struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

type day4 struct {
	given []section
}

func New() runner.Day {
	return &day4{}
}
func (d *day4) Open() {
	in := make([]section, 0)
	lines := inputs.LinesAsString(4)
	for _, l := range lines {
		pair := strings.Split(l, ",")     // ["2-4", "6-8"]
		r1 := strings.Split(pair[0], "-") // ["2", "4"]
		r2 := strings.Split(pair[1], "-") // ["6", "8"]
		x1, _ := strconv.Atoi(r1[0])      // 2
		y1, _ := strconv.Atoi(r1[1])      // 4
		x2, _ := strconv.Atoi(r2[0])      // 6
		y2, _ := strconv.Atoi(r2[1])      // 8
		section := section{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
		in = append(in, section)
	}
	d.given = in
}

func (d *day4) Close() {
	d.given = nil
}

func (d *day4) Part1() string {
	result := 0

	for _, s := range d.given {
		if isInclusive(s) {
			result++
		}
	}

	return fmt.Sprint(result)
}

func (d *day4) Part2() string {
	result := 0

	for _, s := range d.given {
		if isOverlap(s) || isInclusive(s) {
			result++
		}
	}

	return fmt.Sprint(result)
}

func isInclusive(s section) bool {
	return (s.x1 <= s.x2 && s.y1 >= s.y2) ||
		(s.x1 >= s.x2 && s.y1 <= s.y2)
}

func isOverlap(s section) bool {
	return (s.x1 <= s.x2 && s.x2 <= s.y1 && s.y1 <= s.y2) ||
		(s.x2 <= s.x1 && s.x1 <= s.y2 && s.y2 <= s.y1)
}
