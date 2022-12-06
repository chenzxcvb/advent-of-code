package day6

import (
	"fmt"

	"github.com/chenzxcvb/advent-of-code/go/common/inputs"
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
	"github.com/chenzxcvb/advent-of-code/go/types"
)

func New() runner.Day {
	return &day6{}
}

type day6 struct {
	given string
}

func (d *day6) Open() {
	d.given = inputs.LinesAsString(6)[0]
}

func (d *day6) Close() {
	d.given = ""
}

func (d *day6) Part1() string {
	marker := make([]string, 0)

	result := 0
	for n, s := range d.given {
		str := string(s)
		marker = append(marker, str)
		if len(marker) == 4 {
			if unique(marker) {
				return fmt.Sprint(n + 1)
			}
			_, marker = shift(marker)
		}
	}

	return fmt.Sprint(result)
}

func (d *day6) Part2() string {
	marker := make([]string, 0)

	result := 0
	for n, s := range d.given {
		str := string(s)
		marker = append(marker, str)
		if len(marker) == 14 {
			if unique(marker) {
				return fmt.Sprint(n + 1)
			}
			_, marker = shift(marker)
		}
	}

	return fmt.Sprint(result)
}

func unique(arr []string) bool {
	res := types.NewStringSet()
	for _, str := range arr {
		if res.Contains(str) {
			return false
		}
		res.Add(str)
	}
	return true
}

func shift(s []string) (string, []string) {
	l := len(s)
	if l == 0 {
		return "", s
	}
	return s[0], s[1:]
}
