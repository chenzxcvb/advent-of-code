package inputs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type CloseFn = func()

func Scanner(day uint) (*bufio.Scanner, CloseFn) {
	basepath, _ := os.Getwd()
	fmt.Sprintf("root path %s", basepath)

	df := filepath.Join(basepath, "../../inputs", fmt.Sprintf("%d.txt", day))
	f, err := os.Open(df)
	if err != nil {
		panic(fmt.Sprintf("Error reading lines from input file %s: %v", df, err))
	}
	return bufio.NewScanner(f), func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}
}

func LinesAsString(day uint) []string {
	lines := make([]string, 0)
	s, close := Scanner(day)
	defer close()
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func LinesAsInt(day uint) []int {
	lines := LinesAsString(day)
	intLs := make([]int, len(lines))
	for i, s := range lines {
		if val, err := strconv.ParseInt(s, 10, 64); err != nil {
			panic(fmt.Sprintf("Error parsing input string %d '%s' to int: %v", i, s, err))
		} else {
			intLs[i] = int(val)
		}
	}
	return intLs
}
