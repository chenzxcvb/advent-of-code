package day3

import (
	"fmt"
	"strings"

	"github.com/chenzxcvb/advent-of-code/go/common/inputs"
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
	"github.com/chenzxcvb/advent-of-code/go/types"
)

type day3 struct {
	given []string
}

func New() runner.Day {
	return &day3{}
}
func (d *day3) Open() {
	d.given = inputs.LinesAsString(3)
}

func (d *day3) Close() {
	d.given = nil
}

func (d *day3) Part1() string {
	uniformity := make([]string, 0)
	for _, s := range d.given {
		elf := stringToArray(s)
		size := len(elf)
		left := elf[0 : size/2]
		right := elf[size/2 : size]
		rightSet := types.NewStringSetFromList(right)

		temp := types.NewStringSet()
		for _, el := range left {
			if temp.Contains(el) {
				continue
			}
			if rightSet.Contains(el) {
				temp.Add(el)
			}
		}
		uniformity = append(uniformity, temp.Items()...)
	}

	return fmt.Sprint(d.totalScore(uniformity))
}

func (d *day3) Part2() string {

	badges := make([]string, 0)

	for n, s := range d.given {
		badge := types.NewStringSet()

		if n%3 == 0 {
			elf := stringToArray(s)
			next := types.NewStringSetFromList(stringToArray(d.given[n+1]))
			last := types.NewStringSetFromList(stringToArray(d.given[n+2]))
			for _, el := range elf {
				if badge.Contains(el) {
					continue
				}
				if next.Contains(el) && last.Contains(el) {
					badge.Add(el)
				}
			}
		}

		badges = append(badges, badge.Items()...)
	}

	return fmt.Sprint(d.totalScore(badges))
}

func (d *day3) totalScore(uniformity []string) int {
	score := 0
	for _, re := range uniformity {
		score += d.score(re)
	}
	return score
}

func (d *day3) score(letter string) int {
	letters := stringToArray("abcdefghijklmnopqrstuvwxyz")
	for s, v := range letters {
		if letter == v {
			return s + 1
		}
		if letter == strings.ToUpper(v) {
			return s + 1 + 26
		}
	}
	return 0
}

func stringToArray(str string) []string {
	return strings.Split(str, "")
}
