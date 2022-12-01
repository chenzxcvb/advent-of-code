package day0

import (
	"fmt"

	"github.com/chenzxcvb/advent-of-code/go/common/inputs"
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
)

func New() runner.Day {
	return &day0{}
}

type day0 struct {
	ms []int
}

func (d *day0) Open() {
	d.ms = inputs.LinesAsInt(0)
}

func (d *day0) Close() {
	d.ms = nil
}

func (d *day0) Part1() string {
	ct := 0
	for i, m := range d.ms[1:] {
		if m > d.ms[i] {
			ct++
		}
	}

	return fmt.Sprint(ct)

}

func (d *day0) Part2() string {
	ct := 0
	for i := 3; i < len(d.ms); i++ {
		if d.sumAt(i) > d.sumAt(i-1) {
			ct++
		}
	}
	return fmt.Sprint(ct)
}

func (d *day0) sumAt(i int) int {
	return d.ms[i-2] + d.ms[i-1] + d.ms[i]
}
