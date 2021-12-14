package main

import (
	"fmt"
	"github.com/chenzxcvb/advend-of-cod/go/day01"
)

func main() {
	days := []runner.Day{
		day01.New()
	}

	runner.Run(days)
}
