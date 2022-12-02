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
	lines := inputs.LinesAsString(2)
	d.given = lines
	fmt.Println(d.given)
}

func (d *day2) Close() {
	d.given = nil
}

func (d *day2) Part1() string {
	result := []int{}
	for _, l := range d.given {
		players := strings.Fields(l)
		fmt.Println(players)
		_, p2 := match(players[0], players[1])
		result = append(result, p2)
	}
	return fmt.Sprint(sum(result))
}

func (d *day2) Part2() string {

	return ""
}

func match(op, me string) (int, int) {
	r1 := 0
	r2 := 0

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
	r1 += s1

	s2 := points2[me]
	r2 += s2

	if s1 == s2 {
		r1 += 3
		r2 += 3
	} else if s1 < s2 {
		r2 += 6
	} else if s1 > s2 {
		r1 += 6
	}

	return r1, r2
}

func sum(arr []int) int {
	total := 0
	for _, c := range arr {
		total += c
	}
	return total
}
