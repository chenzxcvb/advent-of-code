package main

import (
	"github.com/chenzxcvb/advent-of-code/go/common/day1"
	"github.com/chenzxcvb/advent-of-code/go/common/runner"
)

func main() {
	days := []runner.Day{
		// day0.New(),
		day1.New(),
	}

	runner.Run(days)
}
