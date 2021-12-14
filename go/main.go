package main

import (
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
	"github.com/chenzxcvb/advent-of-code/go/day01"
)

func main() {
	days := []runner.Day{
		day01.New(),
	}

	runner.Run(days)
}
