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

var (
	she = map[string]string{
		"A": "ROCK",
		"B": "PAPER",
		"C": "SCISSORS",
	}
	me = map[string]string{
		"X": "ROCK",
		"Y": "PAPER",
		"Z": "SCISSORS",
	}
	loseTieWin = map[string][]string{
		"ROCK":     {"SCISSORS", "ROCK", "PAPER"},
		"PAPER":    {"ROCK", "PAPER", "SCISSORS"},
		"SCISSORS": {"PAPER", "SCISSORS", "ROCK"},
	}
	score = map[string]int{
		"ROCK":     1,
		"PAPER":    2,
		"SCISSORS": 3,
	}
)

func (d *day2) Part1() string {
	result := 0
	for _, line := range d.given {
		ro := strings.Split(line, " ")
		l := she[ro[0]]
		r := me[ro[1]]
		result = result + score[r] + match(l, r)

	}
	return fmt.Sprint(result)
}

func (d *day2) Part2() string {
	rule := map[string]int{
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	result := 0
	for _, line := range d.given {
		ro := strings.Split(line, " ")
		her := she[ro[0]]
		outcome := rule[ro[1]] // "X" - 0

		// her: "ROCK", me:?, me ${outcome}
		result = result + outcome*3 + score[loseTieWin[her][outcome]]
	}

	return fmt.Sprint(result)
}

func match(her, me string) int {
	for n, s := range loseTieWin[her] {
		if me == s {
			return n * 3
		}
	}
	return 0
}
