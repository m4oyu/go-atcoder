package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	s := sc.Text()
	sc.Scan()
	t := sc.Text()

	diffs := math.Abs(float64(int(s[0]) - int(s[1])))
	difft := math.Abs(float64(int(t[0]) - int(t[1])))

	if ((diffs == 1 || diffs == 4) && (difft == 1 || difft == 4)) || ((diffs == 2 || diffs == 3) && (difft == 3 || difft == 2)) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
