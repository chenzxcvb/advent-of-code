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
		arr := strings.Split(s, "")
		size := len(arr)
		l := arr[0 : size/2]
		r := arr[size/2 : size]

		temp := types.NewStringSet()
		for _, el := range l {
			if temp.Contains(el) {
				continue
			}
			found := d.found(el, r)
			if found {
				temp.Add(el)
			}
		}
		uniformity = append(uniformity, temp.Items()...)
	}

	score := 0
	for _, re := range uniformity {
		score += d.score(re)
	}

	return fmt.Sprint(score)
}

func (d *day3) Part2() string {

	return fmt.Sprint(0)
}

func (d *day3) found(el string, r []string) bool {
	for _, er := range r {
		if el == er {
			return true
		}
	}
	return false
}

func (d *day3) score(letter string) int {
	letters := "a b c d e f g h i j k l m n o p q r s t u v w x y z"
	c := strings.Split(letters, " ")
	for s, v := range c {
		if letter == v {
			return s + 1
		}
		if letter == strings.ToUpper(v) {
			return s + 1 + 26
		}
	}
	return 0
}
