package day2

import (
	"fmt"
	"strings"

	"github.com/chenzxcvb/advent-of-code/go/common/inputs"
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
)

type day2 struct {
	given []string
}

func New() runner.Day {
	return &day2{}
}
func (d *day2) Open() {
	d.given = inputs.LinesAsString(2)
}

func (d *day2) Close() {
	d.given = nil
}

func (d *day2) Part1() string {
	result := 0
	for _, l := range d.given {
		players := strings.Fields(l)
		score := match(players[0], players[1])
		result += score
	}
	return fmt.Sprint(result)
}

func (d *day2) Part2() string {

	return ""
}

func match(op, me string) int {
	result := 0

	points1 := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	points2 := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	s1 := points1[op]
	s2 := points2[me]
	result += s2

	if s1 == s2 {
		result += 3
	} else if s1 < s2 {
		result += 6
	}

	return result
}
