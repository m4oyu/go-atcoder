package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	const buf = 200100
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, buf), buf)
}

func debug() {
	input, err := os.Open("./input")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no input file")
		os.Exit(1)
	}
	sc = bufio.NewScanner(input)
}

func main() {
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	g, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	glass := 0
	mug := 0

	for i := 0; i < k; i++ {
		if glass == g {
			glass = 0
		} else if mug == 0 {
			mug = m
		} else {
			if g-glass > mug {
				glass += mug
				mug = 0
			} else {
				mug = mug - (g - glass)
				glass = g
			}
		}
	}
	fmt.Printf("%d %d\n", glass, mug)
}
