package main

import (
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
	"github.com/chenzxcvb/advent-of-code/go/day1"
)

func main() {
	days := []runner.Day{
		// day0.New(),
		day1.New(),
	}

	runner.Run(days)
}
