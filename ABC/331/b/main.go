package main

import (
	"bufio"
	"fmt"
	"math"
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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s, _ := strconv.Atoi(sc.Text()) // 6
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text()) // 8
	sc.Scan()
	l, _ := strconv.Atoi(sc.Text()) // 12

	min := make([]int, 120)
	for i := range min {
		min[i] = math.MaxInt
	}
	min[0] = 0

	for i := 0; i < 100; i++ {
		if min[i] == math.MaxInt {
			continue
		}
		if min[i]+s < min[i+6] {
			min[i+6] = min[i] + s
		}
		if min[i]+m < min[i+8] {
			min[i+8] = min[i] + m
		}
		if min[i]+l < min[i+12] {
			min[i+12] = min[i] + l
		}
	}

	ans := math.MaxInt
	for i := n; i <= n+12; i++ {
		if min[i] < ans {
			ans = min[i]
		}
	}
	fmt.Println(ans)
}
